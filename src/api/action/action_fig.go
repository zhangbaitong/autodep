package action

import (
	"api/common"
	"encoding/json"
	"fmt"
	"strings"
)

const FIG_PATH = "/home/tomzhao/fig/"

func getProPath(params map[string]interface{}) (ret bool, err string) {
	//获取项目名称
	strFigDirectory, ok := params["fig_directory"].(string)
	if !ok {
		return false, "fig directory empty!!!!"
	}
	str := strings.Split(strFigDirectory, "/")

	strProjectName := str[len(str)-1]

	strRemoteDir := FIG_PATH + strProjectName

	return true, strRemoteDir

}

func fig_transfer(strServerIP string, params map[string]interface{}) (ret bool, err string) {
	var (
		strRemoteDir string
		ok           bool
	)
	//获取项目名称
	strFigDirectory, ok := params["fig_directory"].(string)
	if !ok {
		return false, "fig directory empty!!!!"
	}
	//str := strings.Split(strFigDirectory, "/")

	//strProjectName := str[len(str)-1]

	strFigData, ok := params["fig_data"].(string)
	if !ok {
		return false, "fig_data empty!!!!"
	}

	//生成项目fig文件（这个必须使用fig作为文件名）
	strFileName := "fig.yml"
	ok = common.SaveFile(strFileName, strFigData)
	if !ok {
		return false, "save fig file empty!!!!"
	}

	//创建远程目录
	//strRemoteDir = FIG_PATH + strProjectName
	strRemoteDir = strFigDirectory
		fmt.Println(strRemoteDir);

	//删除远程目录
	_, _ = common.ExecRemoteRM(strServerIP, strRemoteDir)
	//支持递归生成不存在目录
	//TODO:需要改进
	ret1, _ := common.ExecRemoteCMD(strServerIP, "mkdir -p", strRemoteDir)
	if ret1 > 0 {
		return false, "Create fig Remote Path faild!!!!"
	}

	//传输文件到远程目录
	strRemoteFile := strServerIP + ":" + strRemoteDir + "/" + strFileName
	ret1, _ = common.TransferFileSSH(strFileName, strRemoteFile)
	if ret1 > 0 {
		return false, "Transfer File faild!!!!"
	}

	//创建启动文件
	//mapCommands, ok := params["commands"].(map[string]interface{})
	fmt.Println("commands=",params["commands"].([]map[string]string))
	commands:=params["commands"].([]map[string]string);
	if(len(commands)>0){
		//strRemoteDir = FIG_PATH + strProjectName + "/startup"
		strRemoteDir = strFigDirectory + "/startup"
		//创建远程目录
		ret1, _ = common.ExecRemoteCMD(strServerIP, "mkdir", strRemoteDir)
		if ret1 > 0 {
			return false, "Create fig Remote Path faild!!!!"
		}

		for i := 0; i < len(commands); i++ {
			for k,v := range commands[i] {
				//保存启动文件
				strStartDir := strRemoteDir+ "/"+k
				ret1, _ = common.ExecRemoteCMD(strServerIP, "mkdir -p", strStartDir)
				if ret1 > 0 {
					return false, "Create fig Remote Path faild!!!!"
				}

				strStartFile := "start.sh"
				ok = common.SaveFile(strStartFile, v)
				if !ok {
					return false, "save start file empty!!!!"
				}

				//传输文件到远程目录
				strRemoteFile := strServerIP + ":" + strStartDir + "/" + strStartFile
				ret1, _ = common.TransferFileSSH(strStartFile, strRemoteFile)
				if ret1 > 0 {
					return false, "Transfer File faild!!!!"
				}

				//远程脚本设置执行权限
				strRemoteFile = strStartDir + "/" + strStartFile
				fmt.Println("strRemoteFile=", strRemoteFile)
				ret1, _ = common.ExecRemoteChmod(strServerIP, "+x", strRemoteFile)
				if ret1 > 0 {
					return false, "Exec Remote Shell faild!!!!"
				}
		}		
		}
	}

	fmt.Println("strFigDirectory=", strFigDirectory)
	fmt.Println("strFile=", strFileName)
	return true, "ok"
}

func FigCreate(request common.RequestData) string {
	strServerIP := request.ServerIP

	params := dealParams(request.Params)
	fmt.Println("params=",params)
	ok, _ := fig_transfer(strServerIP, params)
	if ok {
		//执行fig命令
		//TODO:exec multi cmd
		retFlag, strRemoteDir := getProPath(params)
		if !retFlag {
			fmt.Println("Get project path is error!")
		}
		ret, _ := common.ExecRemoteShell(strServerIP, " cd "+strRemoteDir+" && "+" fig up")
		if ret > 0 {
			fmt.Println("exec fig up is error!")
		} else {
			return "ok"
		}
	}
	//common.DisplayJson(params)
	return "faild"
}

//处理从前台传过来的函数
func dealParams(strParam string) map[string]interface{} {
	ret := map[string]interface{}{}
	figData := ""
	commands := []map[string]string{}
	temp, _ := common.Config().String("fig", "figDirectory")

	var params FigParams
	err := json.Unmarshal([]byte(strParam), &params)
	if err != nil {
		logger.Println("json data decode faild :", err)
	}

	figDirectory := temp + "/" + params.Project_name
	//servers := params["servers"].([]map[string]interface{})
	servers := params.Servers
	//servers := params["servers"].([][]interface{})
	for _, server := range servers {
		figData += dealServer(server.Server_name)
		figData += dealOneValue("image", server.Image)
		figData += dealMoreValues("ports", server.Ports)
		figData += dealMoreValues("links", server.Links)
		figData += dealMoreValues("volumes", server.Volumes)
		figData += dealCommand(server.Server_name, server.Command, figDirectory)
		command := dealCommandContent(server.Server_name, server.Command)
		commands = append(commands, command)
	}
	ret["fig_data"] = figData
	ret["commands"] = commands
	ret["fig_directory"] = figDirectory
	return ret
}

//处理fig.yml中的服务名格式
func dealServer(server string) (ret string) {
	return server + ":" + "\n"
}

//处理fig.yml文件中只有单个值的数据，比如image,command
func dealOneValue(name, value string) string {
	return "  " + name + ": " + value + "\n"
}

//处理fig.yml文件中有多个值的数据，比如ports,links
func dealMoreValues(name string, values []string) string {
	ret := ""
	for _, value := range values {
		if "" != strings.TrimSpace(value) {
			if "port" == name {
				ret += "    - \"" + value + "\"" + "\n"
			} else {
				ret += "    - " + value + "\n"
			}
		}
	}

	if "" != ret {
		ret = "  " + name + ":" + "\n" + ret
	}
	return ret
}

//处理fig.yml文件中的command命令
func dealCommand(serverName, command, figDirectory string) string {
	if "" != strings.TrimSpace(command) {
		return "  command: " + figDirectory + "/startup/" + serverName + "/start.sh" + "\n"
	}
	return ""
}

//处理fig.yml文件中command的内容
func dealCommandContent(serverName, command string) map[string]string {
	ret := map[string]string{}
	if "" != strings.TrimSpace(command) {
		tmpcommand := "#!/bin/bash" + "\n" + command + "\n" + "tail -f /root/docker/daemonize/daemonize" + "\n"
		ret[serverName] = tmpcommand
	}
	return ret
}

type FigParams struct {
	Project_name string
	Servers      []Server
}

type Server struct {
	Server_name string
	Image       string
	Ports       []string
	Links       []string
	Volumes     []string
	Command     string
}
