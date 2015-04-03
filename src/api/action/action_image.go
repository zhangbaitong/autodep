package action

import (
	"api/common"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/codeskyblue/go-sh"
	"github.com/samalba/dockerclient"
	"log"
	"os"
	"strings"
	"time"
)

type CreateImageStru struct {
	Template   string
	Image_name string
	Code_path  string
	Creator    string
	Remark     string
}

type ListCreateImage struct {
	Image_name string
	Start_time string
	End_time   string
}

type QueryImages struct {
	Image_name  string
	Creator     string
	Create_time string
	Remark      string
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

func ListCreateImages(request common.RequestData) (code int, result string) {
	var image ListCreateImage
	err := json.Unmarshal([]byte(request.Params), &image)
	if err != nil {
		logger.Println("json data decode faild :", err)
	}

	where := "where 1=1 "

	//参数校验

	if image.Image_name != "" {
		where += fmt.Sprintf(" and image_name like '%%%s%%' ", image.Image_name)
	}

	if image.Start_time != "" {
		where += fmt.Sprintf(" and create_time > '%s' ", image.Start_time)
	}

	if image.End_time != "" {
		where += fmt.Sprintf(" and create_time < '%s' ", image.End_time)
	}

	fmt.Println("where", where)

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
		return 1, "faild"
	}
	defer db.Close()

	rows, err := db.Query("select image_name,creator,create_time,remark from images " + where)
	if err != nil {
		log.Fatal(err)
		return 1, "faild"
	}
	defer rows.Close()
	var images []QueryImages = make([]QueryImages, 0)
	for rows.Next() {
		var i QueryImages
		rows.Scan(&i.Image_name, &i.Creator, &i.Create_time, &i.Remark)

		images = append(images, i)
	}

	strImages, err := json.Marshal(images)
	if err != nil {
		log.Fatal(err)
		return 1, "faild"
	}

	return 0, string(strImages)
}

func CreateImage(request common.RequestData) (code int, result string) {
	var image CreateImageStru
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

	dockerfileDirectory := createDockerfile(image.Template, image.Code_path)

	if "" == dockerfileDirectory {
		code = 1
		result = "dockerfile_directory is empty,create_dockerfile error"
		return code, result
	}
	fmt.Println("build镜像,生成Dockerfile成功")

	_, buildErr := buildImage(image.Image_name, dockerfileDirectory)
	if buildErr != "" {
		code = 1
		result = "build image err:" + buildErr
		return code, result
	}
	fmt.Println("build镜像成功")

	code, result = saveImageToDb(image)
	if code != 0 {
		return code, result
	}
	fmt.Println("build镜像，增加数据成功")

	return 0, "build镜像"
}

func saveImageToDb(params CreateImageStru) (code int, result string) {

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return 1, "连接数据库失败"
	}
	stmt, err := tx.Prepare("insert into images(image_name,creator,create_time,remark) values(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return 1, "sql语句有错误"
	}
	defer stmt.Close()

	_, err = stmt.Exec(params.Image_name, params.Creator, time.Now().Format("2006-01-02"), params.Remark)

	if err != nil {
		log.Fatal("参数1", err)
		return 1, "执行sql出错"
	}

	tx.Commit()

	return 0, "增加数据成功"
}

func buildImage(imageName, dockerfileDirectory string) (ret int, err string) {
	fmt.Println("imageName", imageName)
	fmt.Println("dockerfileDirectory", dockerfileDirectory)
	//return common.Execsh("build image error", "docker build -t  "+imageName+"  "+dockerfileDirectory)
	sh.Command("docker", "build", "-t", imageName, dockerfileDirectory).Run()

	return 0, ""
}

func addNewContent(oldContent, addFlag, addContent string) string {
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

func createDockerfile(template, codePath string) string {
	dockerfile, _ := common.Config().String("image", "dockerfile_template")
	datetime := time.Now().Format("2006-01-02")
	folder := dockerfile + "/" + datetime + "/" + template
	pos := strings.LastIndex(codePath, "/")
	codePathPrev := common.SubstrBefore(codePath, pos)
	if "" == codePathPrev {
		codePathPrev = "/"
	}
	relativePath := "./" + common.SubstrAfter(codePath, pos)
	addContent := "\n" + "ADD  " + relativePath + "  /data/" + template + "_code" + "\n"

	//读取模版，生成目标Dockerfile文件
	templateContent := common.ReadFile(dockerfile + "/" + template + "/Dockerfile")
	newContent := addNewContent(templateContent, "EXPOSE,CMD", addContent)
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
	fileDirectory := common.SubstrBefore(filePath, pos)
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
