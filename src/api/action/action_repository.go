package action

import (
	"api/common"
	"client"
	"encoding/json"
	"fmt"
	"log"
	"strings"
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
	//ip, _ = common.Config().String("autodep", "ip")
	//common.Log().Println("common inint ip - ", ip)
	logger = common.Log()
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

func ActionRegTags(ns []string, rep []string,ServerIP string) string {
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

func routineSearch(image *Image, ch chan string,ServerIP string) {
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
		go routineSearch(&repo.Results[i], ch,request.ServerIP)
	}

	for i := 0; i < repo.Num_results; i++ {
		v := <-ch
		logger.Println("Received tag is ", i, v)
		code=1;result="faild"
		return code,result
	}

	rets, _ := json.Marshal(repo)
	return 0,string(rets)
}
