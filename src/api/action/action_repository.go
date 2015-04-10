package action

import (
	"api/common"
	"client"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"net/http"
	"io/ioutil"
)

const (
	SEARCH             = "http://%s:5000/v1/search"
	TAGS               = "http://%s:5000/v1/repositories/%s/%s/tags"
	NAMESPACE_DEFAULT  = "library"
	REPOSITORY_DEFAULT = "registry"
)

//var ip string

var logger *log.Logger

func init() {
	if logger==nil{
		logger = common.Log()
	}
}

func search(url string) string {
	ret, flag := client.GetHTTP(url)
	if !flag {
		ret = ""
	}
	return ret
}

func ActionRegList(ServerIP string) string {
	url := fmt.Sprintf(SEARCH, ServerIP)
	return search(url)
}

func ActionRegTags(ns []string, rep []string, ServerIP string) string {
	var url string
	if len(ns) > 0 && len(rep) > 0 {
		url = fmt.Sprintf(TAGS, ServerIP, ns[0], rep[0])
	} else {
		url = fmt.Sprintf(TAGS, ServerIP, NAMESPACE_DEFAULT, REPOSITORY_DEFAULT)
	}
	return search(url)
}

/*
func ActionRegSearch(q []string, n []string, page []string,ServerIP string) string {
	if len(q) == 0 || len(n) == 0 || len(page) == 0 {
		return ActionRegList()
	}
	url := fmt.Sprintf(SEARCH+"?q=%s&page=%s&n=%s", ServerIP, q[0], page[0], n[0])
	return search(url)
}
*/
type Image struct {
	Description string
	Name        string
	Tag         string
}
type Repository struct {
	Num_results int
	Query       string
	Results     []Image
}

func (image *Image) getNS() string {
	return strings.Split(image.Name, "/")[0]
}

func (image *Image) getName() string {
	return strings.Split(image.Name, "/")[1]
}

func routineSearch(image *Image, ch chan string, ServerIP string) {
	url := fmt.Sprintf(TAGS, ServerIP, image.getNS(), image.getName())
	retTag := search(url)

	var imageTag map[string]interface{}
	err := json.Unmarshal([]byte(retTag), &imageTag)
	if err != nil {
		logger.Println("json data decode failed :", err)
	}

	for k, v := range imageTag {
		fmt.Println(k, v)
		image.Tag = k
	}

	logger.Println(image.Tag)
	ch <- retTag
}

func ActionAllInfo(request common.RequestData) (code int,result string) {
	ret_json := ActionRegList(request.ServerIP)

	var repo Repository
	err := json.Unmarshal([]byte(ret_json), &repo)
	if err != nil {
		logger.Println("json data decode faild :", err)
		code=1;result="faild"
		return code,result
	}
	logger.Println("Method ActionAllInfo Num_results : ", repo.Num_results)
	ch := make(chan string, repo.Num_results)
	for i := 0; i < len(repo.Results); i++ {
		go routineSearch(&repo.Results[i], ch, request.ServerIP)
	}

	for i := 0; i < repo.Num_results; i++ {
		v := <-ch
		logger.Println("Received tag is ", i, v)
	}

	rets, _ := json.Marshal(repo)
	return 0,string(rets)
}

func GetTag(params string) (strLocalTag string, strRemoteTag string) {
	var req interface{}
	err := json.Unmarshal([]byte(params), &req)
	if err != nil {
		return "", ""
	}
	data, _ := req.(map[string]interface{})
	strLocalTag, ok := data["local_tag"].(string)
	if !ok {
		strLocalTag=""
	}
	strRemoteTag, ok = data["remote_tag"].(string)
	if !ok {
		strRemoteTag=""
	}
	return strLocalTag, strRemoteTag
}

func RegTag(request common.RequestData) (code int,result string) {
	strLocalTag, strRemoteTag := GetTag(request.Params)
	if len(strLocalTag)==0 || len(strRemoteTag)==0 {
		return 1, "faild"
	}
	logger.Println("strLocalTag=", strLocalTag)
	logger.Println("strRemoteTag=", strRemoteTag)

	strCMD:=fmt.Sprintf("docker tag %s %s",strLocalTag,strRemoteTag)
	ret, out := common.ExecRemoteDocker(request.ServerIP, strCMD)
	logger.Println("out=", string(out))
	if ret > 0 {
		fmt.Println("docker tag up is error!")
		code = 1
	} else {
		code = 0
	}
	if strings.Contains(out,"No such id"){
		code=1
	}

	return code, string(out)
}

func RegPush(request common.RequestData) (code int,result string) {
	strLocalTag, _ := GetTag(request.Params)
	if len(strLocalTag)==0{
		return 1, "faild"
	}
	strCMD:=fmt.Sprintf("docker push %s",strLocalTag)
	ret, out := common.ExecRemoteDocker(request.ServerIP,strCMD)
	if ret > 0 {
		fmt.Println("exec docker push  is error!")
		code = 1
	} else {
		code = 0
	}

	if strings.Contains(out,"Error:"){
		code=1
	}

	return code, out
}

func RegPull(request common.RequestData) (code int,result string) {
	strLocalTag, _ := GetTag(request.Params)
	if len(strLocalTag)==0{
		return 1, "faild"
	}

	strCMD:=fmt.Sprintf("docker pull %s",strLocalTag)
	ret, out := common.ExecRemoteDocker(request.ServerIP,strCMD)
	if ret > 0 {
		fmt.Println("exec docker push  is error!")
		code = 1
	} else {
		code = 0
	}

	if strings.Contains(out,"Error:"){
		code=1
	}

	return code, out
}

func GetRepository(params string) (strRepository string, strTags string) {
	var req interface{}
	err := json.Unmarshal([]byte(params), &req)
	if err != nil {
		return "", ""
	}
	data, _ := req.(map[string]interface{})
	strRepository, ok := data["repository"].(string)
	if !ok {
		strRepository=""
	}
	strTags, ok = data["tags"].(string)
	if !ok {
		strTags=""
	}
	return strRepository, strTags
}

func RegDelete(request common.RequestData) (code int,result string) {	
	strRepository, strTags := GetRepository(request.Params)
	if len(strRepository)==0 || len(strTags)==0 {
		return 1, "faild"
	}
	logger.Println("strRepository=", strRepository)
	logger.Println("strTags=", strTags)

	//strURL:=fmt.Sprintf("http://%s:%d/v1/repositories/library/%s/tags/%s",request.ServerIP,request.Port,strRepository,strTags)
	if len(strTags)>0 {	
		strURL:=fmt.Sprintf("http://%s:%d/v1/repositories/%s/tags/%s",request.ServerIP,request.Port,strRepository,strTags)
		logger.Println("strURL=", strURL)
		req, err := http.NewRequest("DELETE", strURL,nil)
		if err != nil {
			return 1,"faild"
		}

		resp, err := http.DefaultClient.Do(req)	
		if err != nil {
			return 1,"faild"
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return 1,"faild"
		}

		logger.Println("body=", string(body))
		if resp.StatusCode == 404 {
			return 1,"faild"
		}
		if resp.StatusCode >= 400 {
			return 1,"faild"
		}
	}

	strURL:=fmt.Sprintf("http://%s:%d/v1/repositories/%s/",request.ServerIP,request.Port,strRepository)
	logger.Println("strURL=", strURL)
	req, err := http.NewRequest("DELETE", strURL,nil)
	if err != nil {
		return 1,"faild"
	}

	resp, err := http.DefaultClient.Do(req)	
	if err != nil {
		return 1,"faild"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 1,"faild"
	}
	logger.Println("body=", string(body))

	return 0, string(body)
}
