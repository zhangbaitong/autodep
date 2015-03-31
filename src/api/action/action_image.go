package action

import (
	"fmt"
	"encoding/json"
	"github.com/samalba/dockerclient"
	"api/common"
)

func ListImages(request common.RequestData) (code int,result string) {
	strDockerServer:= fmt.Sprintf("%s:%d",request.ServerIP,request.Port)
	fmt.Println("strDockerServer=", strDockerServer)	
	docker, _ := dockerclient.NewDockerClient(strDockerServer, nil)

	code=0
	images, err := docker.ListImages()
	if err != nil {
		logger.Println("List images faild!")
		code=1;result=""
		return code,result
	}

	strRet, _ := json.Marshal(images)
	result =string(strRet)
	return code,result
}
