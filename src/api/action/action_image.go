package action

import (
	"encoding/json"
	"github.com/samalba/dockerclient"
)

func ActionImages() string {
	docker, _ := dockerclient.NewDockerClient("117.78.19.76:4243", nil)

	images, err := docker.ListImages()
	if err != nil {
		logger.Println("List images faild!")
	}

	ret, _ := json.Marshal(images)
	return string(ret)
}

// func Actionversion() string {
// 	docker, _ := dockerclient.NewDockerClient("127.0.0.1:4243", nil)

// 	// Get only running containers
// 	version, err := docker.Version()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("version type", version.Version)
// 	return version.Version
// }
