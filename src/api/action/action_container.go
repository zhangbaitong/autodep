package action

import (
	"github.com/samalba/dockerclient"
	"log"
	"fmt"
	"api/common"
)

func ActionListContainers(request map[string]interface{}) string {
	common.DisplayJson(request)
	strServerIP, _ := request["ServerIP"].(string)
	strDockerServer:= fmt.Sprintf("%s:%.0f",strServerIP,request["Port"])
	client, _ := dockerclient.NewDockerClient(strDockerServer, nil)

	// Get only running containers
	containers, err := client.ListContainers(true, false, "")
	if err != nil {
		log.Fatal("cannot get containers: %s", err)
	}

	fmt.Println("containers :", containers)
	return containers
}
