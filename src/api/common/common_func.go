package common

import (
	"fmt"
	"github.com/codeskyblue/go-sh"
	"github.com/robfig/config"
	"io/ioutil"
	"log"
	"os"
)

func DisplayJson(obj_json map[string]interface{}) {
	Log().Println(obj_json)
	Log().Println("----------------------parse start------------------------")
	for k, v := range obj_json {
		switch v2 := v.(type) {
		case string:
			fmt.Println(k, "is string", v2)
		case int:
			fmt.Println(k, "is int ", v2)
		case bool:
			fmt.Println(k, "is bool", v2)
		case []interface{}:
			fmt.Println(k, "is array", v2)
			for i, iv := range v2 {
				fmt.Println(i, iv)
			}
		case map[string]interface{}:
			fmt.Println(k, "is map")
			DisplayJson(v2)
		default:
			fmt.Println(k, "is another type not handle yet")
		}
	}
	Log().Println("----------------------parse end------------------------")
}

const (
	SUCCESS     int    = 0
	FAILT       int    = 1
	SUCCESS_MSG string = "ok"
)

func execsh(msg string, cmd string, cmds ...interface{}) (ret int, errmsg string) {
	session := sh.NewSession()
	session.ShowCMD = true
	out, err := session.Command(cmd, cmds...).Output()
	fmt.Println("out=", string(out), "err=", err)
	if err != nil {
		Log().Println(msg, ":", err)
		return FAILT, string(out)
	}

	return SUCCESS, string(out)
}

func Execsh(msg string, cmd string, cmds ...interface{}) (ret int, errmsg string) {
	return execsh(msg, cmd, cmds)
}

func TransferFileSSH(strSrcFile string, strDestFile string) (ret int, err string) {
	return execsh("transfer remote file faild error", "scp", strSrcFile, strDestFile)
}

func ExecRemoteDocker(strServerIP string, strCMD string) (ret int, err string) {
	return execsh("exec remote shell faild error", "ssh", strServerIP, strCMD)
}

func ExecRemoteCMD(strServerIP string, strCMD string, strPath string) (ret int, err string) {
	return execsh("exec remote shell faild error", "ssh", strServerIP, strCMD, strPath)
}

func ExecRemoteChmod(strServerIP string, strPrivilege string, strFile string) (ret int, err string) {
	return execsh("exec remote shell faild error", "ssh", strServerIP, "chmod", strPrivilege, strFile)
}

func ExecRemoteRM(strServerIP string, strFile string) (ret int, err string) {
	return execsh("exec remote shell faild error", "ssh", strServerIP, "rm", "-rf", strFile)
}

func ExecRemoteShell(strServerIP string, strShell string) (ret int, err string) {
	return execsh("exec remote shell faild error", "ssh", strServerIP, strShell)
}

func SaveFile(strFileName string, strData string) (ok bool) {
	f, err := os.Create(strFileName)
	if err != nil {
		fmt.Println("create file faild error:", err)
		return false
	}
	_, err_w := f.Write([]byte(strData))
	if err_w != nil {
		fmt.Println("Server start faild error:", err_w)
		return false
	}
	return true
}

func Config() *config.Config {
	file, _ := os.Getwd()
	c, _ := config.ReadDefault(file + "/common/config.cfg")
	fmt.Println("Config init success ...")
	return c
}

func Log() *log.Logger {
	logger := log.New(os.Stdout, "autodep log : ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Print("logger init success ...")
	return logger
}

//截取固定位置以前的字符串
func SubstrBefore(s string, l int) string {
	if len(s) <= l {
		return s
	}
	ret, rs := "", []rune(s)

	for i, r := range rs {
		if i >= l {
			break
		}

		ret += string(r)
	}
	return ret
}

//截取固定位置以后的字符串
func SubstrAfter(s string, l int) string {
	if len(s) <= l {
		return s
	}
	ret, rs := "", []rune(s)

	for i, r := range rs {
		if i <= l {
			continue
		}

		ret += string(r)
	}
	return ret
}

//判断文件夹是否存在
func IsDirExists(path string) bool {
	fi, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}
}

//判断文件是否存在
func IsFileExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	} else {
		return !fi.IsDir()
	}
}

//从文件中读取内容
func ReadFile(filePth string) string {
	bytes, err := ioutil.ReadFile(filePth)
	if err != nil {
		fmt.Println("读取文件失败: ", err)
		return ""
	}

	return string(bytes)
}
