package action

import (
	"github.com/samalba/dockerclient"
	"log"
	"fmt"
	"api/common"
	"encoding/json"
)

func DisplayContainers(list []dockerclient.Container) {
	for i := 0; i < len(list); i++ {
		info:=fmt.Sprintf("ID=%s;Names=%s;Image=%s",list[i].Id,list[i].Names,list[i].Image)
		fmt.Println(info);
	}
	
}

func GetContainerID(params string) (ret string, ok bool){
	var req interface{}
	err := json.Unmarshal([]byte(params), &req)
	if err != nil {
		return "",false
	}
	data, _ := req.(map[string]interface{})
	strID, ok := data["id"].(string)
	if !ok {
		return "",false
	}
	return strID,true
}

func CreateContainer(request common.RequestData)(code int,result string) {
	strDockerServer:= fmt.Sprintf("%s:%d",request.ServerIP,request.Port)
	fmt.Println("strDockerServer=", strDockerServer)	
	client, _ := dockerclient.NewDockerClient(strDockerServer, nil)

	var data map[string]string
	err := json.Unmarshal([]byte(request.Params), &data)
	if err != nil {
		fmt.Println("json data decode faild :", err)
		code=1;result="faild"
		return code,result
	}

	fmt.Println("Info=", data)
	//fmt.Println("imageName:%s;containerName=%s", Info.imageName,Info.containerName)
	cmd:=[]string{"/bin/bash"}
	containerConfig := &dockerclient.ContainerConfig{
	    Image:          data["imageName"],
	    Cmd:			 cmd,
	    Tty:			 true,
	}

	fmt.Println("containerName=", data["containerName"])
	containerID, err := client.CreateContainer(containerConfig, data["containerName"])
     if err != nil {
		log.Fatal("cannot create container: %s", err)
		code=1;result="faild"
		return code,result
	}

    // Start the container
	createContHostConfig := &dockerclient.HostConfig{
	    Binds:           []string{"/var/run:/var/run", "/sys:/sys", "/var/lib/docker:/var/lib/docker"},
	    PublishAllPorts: true,
	    Privileged:      false,
	}
    err = client.StartContainer(containerID,createContHostConfig)
    if err != nil {
	fmt.Println(err)
	code=1;result="faild"
    }
	return code,result
}

func ListContainers(request common.RequestData) (code int,result string)  {
	strDockerServer:= fmt.Sprintf("%s:%d",request.ServerIP,request.Port)
	fmt.Println("strDockerServer=", strDockerServer)	
	client, _ := dockerclient.NewDockerClient(strDockerServer, nil)

	containers, err := client.ListContainers(true, false,"")
	if err != nil {
		log.Fatal("cannot get containers: %s", err)
		code=1;result="faild"
		return code,result
	}

	DisplayContainers(containers)
	rets, _ := json.Marshal(containers)
	return 0,string(rets)
}

func InspectContainer(request common.RequestData)(code int,result string){
	strDockerServer:= fmt.Sprintf("%s:%d",request.ServerIP,request.Port)
	fmt.Println("strDockerServer=", strDockerServer)	
	client, _ := dockerclient.NewDockerClient(strDockerServer, nil)

	strID, ok := GetContainerID(request.Params)
	if !ok {
		strID=""
	}
	//fmt.Println("strID=", strID)

	containerInfo, err := client.InspectContainer(strID)
	if err != nil {
		log.Fatal("cannot get containers: %s", err)
		code=1;result="faild"
		return code,result
	}

	rets, _ := json.Marshal(containerInfo)
	return 0,string(rets)
}

func ContainerChanges(request common.RequestData)(code int,result string)  {
	strDockerServer:= fmt.Sprintf("%s:%d",request.ServerIP,request.Port)
	fmt.Println("strDockerServer=", strDockerServer)	
	client, _ := dockerclient.NewDockerClient(strDockerServer, nil)

	strID, ok := GetContainerID(request.Params)
	if !ok {
		strID=""
	}
	fmt.Println("strID=", strID)

	changes, err := client.ContainerChanges(strID)
	if err != nil {
		log.Fatal("cannot get containers: %s", err)
		code=1;result="faild"
		return code,result
	}

	rets, _ := json.Marshal(changes)
	return 0,string(rets)
}

func StopContainer(request common.RequestData) (code int,result string)  {
	strDockerServer:= fmt.Sprintf("%s:%d",request.ServerIP,request.Port)
	fmt.Println("strDockerServer=", strDockerServer)	
	client, _ := dockerclient.NewDockerClient(strDockerServer, nil)

	strID, ok := GetContainerID(request.Params)
	if !ok {
		strID=""
	}
	//fmt.Println("strID=", strID)
	nTime:=30
	err := client.StopContainer(strID,nTime)
	if err != nil {
		log.Fatal("cannot get containers: %s", err)
		code=1;result="faild"
		return code,result
	}

	return 0,"ok"
}

func RestartContainer(request common.RequestData) (code int,result string)  {
	strDockerServer:= fmt.Sprintf("%s:%d",request.ServerIP,request.Port)
	fmt.Println("strDockerServer=", strDockerServer)	
	client, _ := dockerclient.NewDockerClient(strDockerServer, nil)

	strID, ok := GetContainerID(request.Params)
	if !ok {
		log.Fatal("cannot Restart Container ", )
		code=1;result="faild"
		return code,result
	}
	//fmt.Println("strID=", strID)

	nTime:=30
	err := client.RestartContainer(strID,nTime)
	if err != nil {
		log.Fatal("cannot get containers: %s", err)
		code=1;result="faild"
		return code,result
	}

	return 0,"ok"
}

func PauseContainer(request common.RequestData) (code int,result string)  {
	strDockerServer:= fmt.Sprintf("%s:%d",request.ServerIP,request.Port)
	fmt.Println("strDockerServer=", strDockerServer)	
	client, _ := dockerclient.NewDockerClient(strDockerServer, nil)

	strID, ok := GetContainerID(request.Params)
	if !ok {
		log.Fatal("cannot get containers: ")
		code=1;result="faild"
		return code,result
	}
	//fmt.Println("strID=", strID)

	err:= client.PauseContainer(strID)
	if err != nil {
		log.Fatal("cannot get containers: %s", err)
		code=1;result="faild"
		return code,result
	}

	return 0,"ok"
}

func UnpauseContainer(request common.RequestData) (code int,result string)  {
	strDockerServer:= fmt.Sprintf("%s:%d",request.ServerIP,request.Port)
	fmt.Println("strDockerServer=", strDockerServer)	
	client, _ := dockerclient.NewDockerClient(strDockerServer, nil)

	strID, ok := GetContainerID(request.Params)
	if !ok {
		log.Fatal("cannot get containers:")
		code=1;result="faild"
		return code,result
	}
	//fmt.Println("strID=", strID)

	err := client.UnpauseContainer(strID)
	if err != nil {
		log.Fatal("cannot get containers: %s", err)
		code=1;result="faild"
		return code,result
	}

	return 0,"ok"
}


func RemoveContainer(request common.RequestData) (code int,result string)  {
	strDockerServer:= fmt.Sprintf("%s:%d",request.ServerIP,request.Port)
	fmt.Println("strDockerServer=", strDockerServer)	
	client, _ := dockerclient.NewDockerClient(strDockerServer, nil)

	strID, ok := GetContainerID(request.Params)
	if !ok {
		log.Fatal("cannot get containers:")
		code=1;result="faild"
		return code,result
	}
	fmt.Println("strID=", strID)

	err := client.RemoveContainer(strID,true,true)
	if err != nil {
		log.Fatal("cannot get containers: %s", err)
		code=1;result="faild"
		return code,result
	}

	return 0,"ok"
}

func KillContainer(request common.RequestData) (code int,result string)  {
	strDockerServer:= fmt.Sprintf("%s:%d",request.ServerIP,request.Port)
	fmt.Println("strDockerServer=", strDockerServer)	
	client, _ := dockerclient.NewDockerClient(strDockerServer, nil)

	strID, ok := GetContainerID(request.Params)
	if !ok {
		log.Fatal("cannot get containers:")
		code=1;result="faild"
		return code,result
	}
	//fmt.Println("strID=", strID)

	err := client.KillContainer(strID,"5")
	if err != nil {
		log.Fatal("cannot get containers: %s", err)
		code=1;result="faild"
		return code,result
	}

	return 0,"ok"
}

/*
func ContainerExec(request map[string]interface{}) string {
	common.DisplayJson(request)
	strServerIP, _ := request["ServerIP"].(string)
	strDockerServer:= fmt.Sprintf("%s:%.0f",strServerIP,request["Port"])
	fmt.Println("strDockerServer=", strDockerServer)	
	client, _ := dockerclient.NewDockerClient(strDockerServer, nil)

	//获取项目名称
	strID, ok := request["Params"].(string)
	if !ok {
		log.Fatal("cannot get containers: ")
		return "faild"
	}
	//fmt.Println("strID=", strID)

	err := client.Exec(strID)
	if err != nil {
		log.Fatal("cannot get containers: %s", err)
		return ""
	}

	return "ok"
}
*/

func InfoContainer(request common.RequestData) (code int,result string)  {
	strDockerServer:= fmt.Sprintf("%s:%d",request.ServerIP,request.Port)
	fmt.Println("strDockerServer=", strDockerServer)	
	client, _ := dockerclient.NewDockerClient(strDockerServer, nil)


	Info, err := client.Info()
	if err != nil {
		log.Fatal("cannot get containers: %s", err)
		code=1;result="faild"
		return code,result
	}

	fmt.Println("Info=", Info)
	rets, _ := json.Marshal(Info)
	return 0,string(rets)
}
