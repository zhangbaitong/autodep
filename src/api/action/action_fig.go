package action

import (
	"api/common"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"strings"
	"time"
)

type FigParams struct {
	Project_name   string
	Fig_project_id string
	Servers        []Server
}

type Server struct {
	Server_name string
	Image       string
	Ports       []string
	Links       []string
	Volumes     []string
	Command     string
}

type FigProject struct {
	ProjectID   int
	ProjectName string
	Machine_ip  string
	Directory   string
	Param       string
	Content     string
	CreateTime  int
}

type Template struct {
	Template_name    string
	Template_type    string
	Template_content string
	Create_time      int
	Remark           string
}

type QueryTemplate struct {
	Template_name string
	Template_type string
}

type UpdateFigStr struct {
	Fig_project_id string
}

func project_count(strServerIP string,strProjectName string) (nCount int) {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		logger.Println(err)
		return 0
	}
	defer db.Close()
	strSql := fmt.Sprintf("select count(fig_project_id) from fig_project where machine_ip= '%s' and project_name = '%s' ", strServerIP,strProjectName)
	fmt.Println("strSql=",strSql)
	rows, err := db.Query(strSql)
	if err != nil {
		logger.Println(err)
		return 0
	}
	defer rows.Close()

	nCount=0
	for rows.Next() {
		rows.Scan(&nCount)
	}
	fmt.Println("nCount=",nCount)

	return nCount
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
	fmt.Println(strRemoteDir)

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
	fmt.Println("commands=", params["commands"].([]map[string]string))
	commands := params["commands"].([]map[string]string)
	if len(commands) > 0 {
		//strRemoteDir = FIG_PATH + strProjectName + "/startup"
		strRemoteDir = strFigDirectory + "/startup"
		//创建远程目录
		ret1, _ = common.ExecRemoteCMD(strServerIP, "mkdir", strRemoteDir)
		if ret1 > 0 {
			return false, "Create fig Remote Path faild!!!!"
		}

		for i := 0; i < len(commands); i++ {
			for k, v := range commands[i] {
				//保存启动文件
				strStartDir := strRemoteDir + "/" + k
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

func FigCreate(request common.RequestData) (code int, result string) {
	ret,params := dealParams(request.ServerIP, request.Params)
	if ret==1{
		return 1, "project was existed"
	}
	ok, err := fig_transfer(request.ServerIP, params)
	code = 1
	result = err
	if ok {
		code = 0
		result = "ok"
	}
	//common.DisplayJson(params)
	return code, result
}

func GetFigDirectory(params string) (ret string, ok bool) {
	fmt.Println("GetFigDirectory params", params)

	var req interface{}
	err := json.Unmarshal([]byte(params), &req)
	if err != nil {
		return "", false
	}
	data, _ := req.(map[string]interface{})
	strFigDirectory, ok := data["fig_directory"].(string)
	if !ok {
		return "", false
	}
	return strFigDirectory, true
}

func GetProjectName(params string) (ret string, ok bool) {
	var req interface{}
	err := json.Unmarshal([]byte(params), &req)
	if err != nil {
		return "", false
	}
	data, _ := req.(map[string]interface{})
	strProjectName, ok := data["project_name"].(string)
	if !ok {
		return "", false
	}
	return strProjectName, true
}

func GetProjectInfo(request common.RequestData) (code int, result string) {
	strProjectName, ok := GetProjectName(request.Params)
	if !ok {
		return 1, "faild"
	}

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		logger.Println(err)
		return 1, "faild"
	}
	defer db.Close()
	strSql := fmt.Sprintf("select fig_project_id, project_name,machine_ip, fig_directory, fig_param, fig_content, create_time from fig_project where project_name like '%%%s%%' ", strProjectName)
	rows, err := db.Query(strSql)
	if err != nil {
		logger.Println(err)
		return 1, "faild"
	}
	defer rows.Close()

	var infoList []FigProject = make([]FigProject, 0)
	for rows.Next() {
		var m FigProject
		rows.Scan(&m.ProjectID, &m.ProjectName, &m.Machine_ip, &m.Directory, &m.Param, &m.Content, &m.CreateTime)
		infoList = append(infoList, m)
	}

	strInfo, err := json.Marshal(infoList)
	if err != nil {
		logger.Println(err)
		return 1, "faild"
	}

	return 0, string(strInfo)
}

func GetInfoById(request common.RequestData) (code int, result string) {

	var params UpdateFigStr

	err := json.Unmarshal([]byte(request.Params), &params)
	if err != nil {
		logger.Println("json data decode faild :", err)
	}

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		logger.Println(err)
		return 1, "faild"
	}
	defer db.Close()
	strSql := fmt.Sprintf("select fig_param from fig_project where fig_project_id = '%s' ", params.Fig_project_id)

	fmt.Println(strSql)
	rows, err := db.Query(strSql)
	if err != nil {
		logger.Println(err)
		return 1, "faild"
	}
	defer rows.Close()

	var infoList []FigProject = make([]FigProject, 0)
	for rows.Next() {
		var m FigProject
		rows.Scan(&m.Param)
		infoList = append(infoList, m)
	}

	strInfo, err := json.Marshal(infoList)
	if err != nil {
		logger.Println(err)
		return 1, "faild"
	}

	return 0, string(strInfo)
}

func FigPS(request common.RequestData) (code int, result string) {

	//获取项目名称
	strFigDirectory, ok := GetFigDirectory(request.Params)
	if !ok {
		return 1, "fig directory empty!!!!"
	}

	ret, out := common.ExecRemoteShell(request.ServerIP, " cd "+strFigDirectory+" && "+" fig ps")
	if ret > 0 {
		fmt.Println("exec fig up is error!")
		code = 1
	} else {
		code = 0
	}
	return code, out
}

func FigRm(request common.RequestData) (code int, result string) {
	//获取项目名称
	strFigDirectory, ok := GetFigDirectory(request.Params)
	if !ok {
		return 1, "fig directory empty!!!!"
	}

	ret, out := common.ExecRemoteShell(request.ServerIP, " cd "+strFigDirectory+" && "+" fig rm  --force")
	if ret > 0 {
		fmt.Println("exec fig up is error!")
		code = 1
	} else {
		code = 0
	}
	return code, out
}

func FigStop(request common.RequestData) (code int, result string) {
	//获取项目名称
	strFigDirectory, ok := GetFigDirectory(request.Params)
	if !ok {
		return 1, "fig directory empty!!!!"
	}

	ret, out := common.ExecRemoteShell(request.ServerIP, " cd "+strFigDirectory+" && "+" fig stop")
	if ret > 0 {
		fmt.Println("exec fig up is error!")
		code = 1
	} else {
		code = 0
	}
	return code, out
}

func FigRestart(request common.RequestData) (code int, result string) {
	//获取项目名称
	strFigDirectory, ok := GetFigDirectory(request.Params)
	if !ok {
		return 1, "fig directory empty!!!!"
	}

	ret, out := common.ExecRemoteShell(request.ServerIP, " cd "+strFigDirectory+" && "+" fig restart")
	if ret > 0 {
		fmt.Println("exec fig up is error!")
		code = 1
	} else {
		code = 0
	}
	return code, out
}

func FigStart(request common.RequestData) (code int, result string) {
	//获取项目名称
	strFigDirectory, ok := GetFigDirectory(request.Params)
	if !ok {
		return 1, "fig directory empty!!!!"
	}

	ret, out := common.ExecRemoteShell(request.ServerIP, " cd "+strFigDirectory+" && "+" fig up -d")
	if ret > 0 {
		fmt.Println("exec fig up is error!")
		code = 1
	} else {
		code = 0
	}
	return code, out
}

func FigRecreate(request common.RequestData) (code int, result string) {
	//获取项目名称
	strFigDirectory, ok := GetFigDirectory(request.Params)
	if !ok {
		return 1, "fig directory empty!!!!"
	}

	ret, out := common.ExecRemoteShell(request.ServerIP, " cd "+strFigDirectory+" && "+" fig stop"+"&&"+"fig rm --force"+"&&"+"fig up -d")
	if ret > 0 {
		fmt.Println("exec fig up is error!")
		code = 1
	} else {
		code = 0
	}
	return code, out
}

//处理从前台传过来的函数
func dealParams(strServerIp string, strParam string) (code int,map[string]interface{}) {

	fmt.Println("传来的参数：", strParam)

	ret := map[string]interface{}{}
	figData := ""
	commands := []map[string]string{}
	temp, _ := common.Config().String("fig", "figDirectory")

	var params FigParams
	err := json.Unmarshal([]byte(strParam), &params)
	if err != nil {
		logger.Println("json data decode faild :", err)
	}

	nCount:=project_count(strServerIP,params.Project_name)
	if nCount>0 {
		return 1, nil
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

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		logger.Println(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		logger.Println(err)
	}

	if params.Fig_project_id != "" {
		stmt, err := tx.Prepare("update fig_project  set fig_param=?,fig_content=?,create_time=? where fig_project_id=?")
		if err != nil {
			logger.Println(err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(strParam, figData, time.Now().Unix(), params.Fig_project_id)

		if err != nil {
			logger.Println("参数1", err)
		}
	} else {
		stmt, err := tx.Prepare("insert into fig_project(project_name,machine_ip,fig_directory,fig_param,fig_content,create_time) values(?, ?, ?, ?, ?, ?)")
		if err != nil {
			logger.Println(err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(params.Project_name, strServerIp, figDirectory, strParam, figData, time.Now().Unix())

		if err != nil {
			logger.Println("参数1", err)
		}
	}
	tx.Commit()

	return 0,ret
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

func GetFigTemplate(request common.RequestData) (code int, result string) {

	var params QueryTemplate
	err := json.Unmarshal([]byte(request.Params), &params)
	if err != nil {
		logger.Println("json data decode faild :", err)
	}

	templateName := "fig_" + params.Template_name

	if templateName == "" {
		return 1, "tempalteName can't be empty"
	}

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		logger.Println(err)
		return 1, "open db error"
	}
	defer db.Close()
	strSql := fmt.Sprintf("select template_name, template_type,template_content, create_time, remark from template where template_name = '%s' and template_type='fig'", templateName)
	rows, err := db.Query(strSql)
	if err != nil {
		logger.Println(err)
		return 1, "perform sql error"
	}
	defer rows.Close()

	var templates []Template = make([]Template, 0)
	for rows.Next() {
		var t Template
		rows.Scan(&t.Template_name, &t.Template_type, &t.Template_content, &t.Create_time, &t.Remark)
		templates = append(templates, t)
	}

	strInfo, err := json.Marshal(templates)
	if err != nil {
		logger.Println(err)
		return 1, "query result turn json error"
	}

	return 0, string(strInfo)
}
