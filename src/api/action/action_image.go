package action

import (
	"api/common"
	"encoding/json"
	"fmt"
	"github.com/samalba/dockerclient"
	"strings"
	"time"
)

type CreateImage struct {
	Template   string
	Image_name string
	Code_path  string
	Creator    string
	Remark     string
}

func ListImages(request common.RequestData) (code int, result string) {
	strDockerServer := fmt.Sprintf("%s:%d", request.ServerIP, request.Port)
	fmt.Println("strDockerServer=", strDockerServer)
	docker, _ := dockerclient.NewDockerClient(strDockerServer, nil)

	code = 0
	images, err := docker.ListImages()
	if err != nil {
		logger.Println("List images faild!")
		code = 1
		result = ""
		return code, result
	}

	strRet, _ := json.Marshal(images)
	result = string(strRet)
	return code, result
}

func CreateImage(request common.RequestData) (code int, result string) {
	var image CreateImage
	err := json.Unmarshal([]byte(request.Params), &image)
	if err != nil {
		logger.Println("json data decode faild :", err)
	}

	//参数校验
	if image.Template == "" {
		code = 1
		result = "template cann't be empty "
		return code, result
	}

	if image.Image_name == "" {
		code = 1
		result = "image_name cann't be empty "
		return code, result
	}

	if image.Code_path == "" {
		code = 1
		result = "code_path cann't be empty "
		return code, result
	}

	if image.Creator == "" {
		code = 1
		result = "creator cann't be empty "
		return code, result
	}

	if "" == createDockerfile(image.Template, image.Code_path) {
		code = 1
		result = "dockerfile_directory is empty,create_dockerfile error"
		return code, result
	}

	return code, result
}

func createDockerfile(template, codePath string) string {
	dockerfile, _ := common.Config().String("image", "dockerfile")
	datetime := time.Now().Format("2006-01-02")
	folder := dockerfile + "/" + datetime + "/" + template
	pos := strings.LastIndex(codePath, "/")
	codePathPrev := common.SubstrBefore(folder, pos)
	if "" == codePathPrev {
		codePathPrev = "/"
	}
	relativePath := "." + common.SubstrAfter(folder, pos)
	addContent := "\n" + "ADD  " + relativePath + "  /data/" + template + "_code"

	//读取模版，生成目标Dockerfile文件
	templateContent := common.ReadFile(dockerfile + "/" + template + "/Dockerfile")
	newContent := addContent(templateContent, "EXPOSE,CMD", addContent)
	createFile(folder+"/Dockerfile", newContent)
	createFile(codePathPrev+"/Dockerfile", newContent)

	return codePathPrev
}

//创建文件并写入内容，如文件已存在，覆盖旧文件
func createFile(filePath, strData string) (code int, result string) {
	if strData == "" {
		return 1, "content不能为空"
	}

	pos := strings.LastIndex(filePath, "/")
	fileDirectory := substrBefore(filePath, pos)
	if !common.IsDirExists(fileDirectory) {
		err := os.MkdirAll(fileDirectory, os.ModePerm) //生成多级目录
		if err != nil {
			fmt.Println("创建目录("+fileDirectory+")失败：", err)
			return 1, "创建目录(" + fileDirectory + ")失败"
		}
	}

	if !common.SaveFile(filePath, strData) {
		return 1, "创建文件失败"
	}

	return 0, "创建文件成功"
}

func addContent(oldContent, addFlag, addContent string) string {
	if oldContent == "" || addFlag == "" || addContent == "" {
		fmt.Println("oldContent、addFlag、addContent不能为空")
		return ""
	}

	for _, f := range strings.Split(addFlag, ",") {

		if f == "" {
			continue
		}

		pos := strings.Index(oldContent, f)

		if pos == -1 {
			continue
		}
		return common.SubstrBefore(oldContent, pos) + addContent + common.SubstrAfter(oldContent, pos-1)
	}

	return ""
}
