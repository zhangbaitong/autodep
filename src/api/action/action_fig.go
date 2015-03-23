package action
import (
    "github.com/codeskyblue/go-sh"

    "fmt"
    "strings"
    "api/common"
)

const FIG_PATH="/home/tomzhao/fig/"

func fig_transfer(strServerIP string,params map[string]interface{})(ret bool,err string){
    var(
        strRemoteDir string
        ok bool
    )
    //获取项目名称
    strFigDirectory,ok:=params["fig_directory"].(string)
    if !ok {
       return false,"fig directory empty!!!!"
    }
    str:=strings.Split(strFigDirectory,"/")

    strProjectName:=str[len(str)-1]

     strFigData,ok:=params["fig_data"].(string)
     if !ok {
        return false,"fig_data empty!!!!"
    }

    //生成项目fig文件
    strFileName:=strProjectName+".yml"
    ok=common.SaveFile(strFileName,strFigData)
    if !ok{
       return false,"save fig file empty!!!!"
    }

    //创建远程目录
    strRemoteDir=FIG_PATH+strProjectName
    ret1,_:=common.ExecRemoteCMD(strServerIP,"mkdir",strRemoteDir)
    if(ret1>0){
       return false,"Create fig Remote Path faild!!!!"
    }

    //传输文件到远程目录
    strRemoteFile:=strServerIP+":"+strRemoteDir+"/"+strFileName
    ret1,_=common.TransferFileSSH(strFileName,strRemoteFile)
    if(ret1>0){
       return false,"Transfer File faild!!!!"
    }

    //创建启动文件
    mapCommands,ok:=params["commands"].(map[string]interface{})
    if ok {
        //创建远程目录
        strRemoteDir=FIG_PATH+strProjectName+"/startup"
        ret1,_:=common.ExecRemoteCMD(strServerIP,"mkdir",strRemoteDir)
        if(ret1>0){
           return false,"Create fig Remote Path faild!!!!"
        }

        common.DisplayJson(mapCommands);
        for k, v := range mapCommands  {
            switch v2 := v.(type) {
            case string:

                //保存启动文件
                strStartFile:="start.sh"
                ok=common.SaveFile(strStartFile,v2)
                if !ok{
                   return false,"save start file empty!!!!"
                }

                //传输文件到远程目录
                strRemoteFile:=strServerIP+":"+strRemoteDir+"/"+strStartFile
                ret1,_=common.TransferFileSSH(strStartFile,strRemoteFile)
                if(ret1>0){
                   return false,"Transfer File faild!!!!"
                }

                //远程脚本设置执行权限
                strRemoteFile=strRemoteDir+"/"+strStartFile
                fmt.Println("strRemoteFile=",strRemoteFile)
                ret1,_=common.ExecRemoteChmod(strServerIP,"+x",strRemoteFile)
                if(ret1>0){
                   return false,"Exec Remote Shell faild!!!!"
                }

                //执行远程脚本
                ret1,_=common.ExecRemoteShell(strServerIP,strRemoteFile)
                if(ret1>0){
                   return false,"Exec Remote Shell faild!!!!"
                }

                fmt.Println(k, "is string", v2)
            default:
                fmt.Println(k, "is another type not handle yet")
            }
        }            
    }
    fmt.Println("strFigDirectory=",strFigDirectory)
    fmt.Println("strFile=",strFileName)
    return true,"ok"
}

func FigCreate(request map[string]interface{}) string {
    session := sh.NewSession()
    session.ShowCMD = true    
    //strVersion,_:= request["Version"].(string)
    strServerIP,_:= request["ServerIP"].(string)
    //nPort,_:= request["Port"].(int)
    //strMethod,_:= request["Method"].(string)

     params,_:=request["Params"].(map[string]interface{})
    ok,_:=fig_transfer(strServerIP,params)
    if ok {
        return "ok"
    }
    //common.DisplayJson(params)
    return "faild"
}
