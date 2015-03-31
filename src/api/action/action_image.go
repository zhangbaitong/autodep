package action

import (
	"fmt"
	"encoding/json"
	"github.com/samalba/dockerclient"
	"api/common"
)

func ListImages(request common.RequestData) string {
	strDockerServer:= fmt.Sprintf("%s:%d",request.ServerIP,request.Port)
	fmt.Println("strDockerServer=", strDockerServer)	
	docker, _ := dockerclient.NewDockerClient(strDockerServer, nil)

	images, err := docker.ListImages()
	if err != nil {
		logger.Println("List images faild!")
	}

	ret, _ := json.Marshal(images)
	return string(ret)
}
