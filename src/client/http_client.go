package client

import (
	"api/common"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type RequestData struct {
	Version  string
	ServerIP string
	Port     int
	Method   string
	Params   string
}

func httpGet() {
	resp, err := http.Get("http://127.0.0.1:8080/v1/version")
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
func httpPost() {
	//var data map[string]string
	//data["imageName"]="centos:latest"
	//data["containerName"]="tomzhao"
	//post_data:=RequestData{Version:"1.0",ServerIP:"117.78.19.76",Port:4243,Method:"container/list",Params:"{\"id\":[\"13100f09aa075adf6129c57467b6564ab3f89f2e4ec8c01bbb312e8258e21ea7\"]}"}
	//post_data:=RequestData{Version:"1.0",ServerIP:"117.78.19.76",Port:4243,Method:"container/info",Params:""}
	//post_data:=RequestData{Version:"1.0",ServerIP:"117.78.19.76",Port:4243,Method:"container/inspect",Params:"13100f09aa075adf6129c57467b6564ab3f89f2e4ec8c01bbb312e8258e21ea7"}
	//post_data:=RequestData{Version:"1.0",ServerIP:"117.78.19.76",Port:4243,Method:"container/changes",Params:"{\"id\":\"13100f09aa075adf6129c57467b6564ab3f89f2e4ec8c01bbb312e8258e21ea7\"}"}
	//post_data:=RequestData{Version:"1.0",ServerIP:"192.168.1.117",Port:4243,Method:"container/stop",Params:"{\"id\":\"6e9f36168a78\"}"}
	//post_data:=RequestData{Version:"1.0",ServerIP:"192.168.1.117",Port:4243,Method:"container/restart",Params:"{\"id\":\"6e9f36168a78\"}"}
	//post_data:=RequestData{Version:"1.0",ServerIP:"192.168.1.117",Port:4243,Method:"container/pause",Params:"{\"id\":\"6e9f36168a78\"}"}
	//post_data:=RequestData{Version:"1.0",ServerIP:"192.168.1.117",Port:4243,Method:"container/unpause",Params:"{\"id\":\"6e9f36168a78\"}"}
	//post_data:=RequestData{Version:"1.0",ServerIP:"192.168.1.117",Port:4243,Method:"container/remove",Params:"{\"id\":\"3fb5c4080f37\"}"}
	//post_data:=RequestData{Version:"1.0",ServerIP:"192.168.1.117",Port:4243,Method:"container/kill",Params:"{\"id\":\"ad7be2d3c897\"}"}
	//post_data:=RequestData{Version:"1.0",ServerIP:"192.168.1.117",Port:4243,Method:"container/create",Params:""}
	//post_data:=RequestData{Version:"1.0",ServerIP:"192.168.1.117",Port:4243,Method:"image/list",Params:""}
	//post_data := RequestData{Version: "1.0", ServerIP: "10.122.75.228", Port: 4243, Method: "reg/info", Params: ""}
	//post_data:=RequestData{Version:"1.0",ServerIP:"117.78.19.76",Port:4243,Method:"reg/tag",Params:"{\"local_tag\":\"centos:latest\",\"remote_tag\":\"10.122.75.228:5000/centostest:latest\"}"}
	//post_data:=RequestData{Version:"1.0",ServerIP:"117.78.19.76",Port:4243,Method:"reg/push",Params:"{\"local_tag\":\"10.122.75.228:5000/centostest:latest\"}"}
	//post_data:=RequestData{Version:"1.0",ServerIP:"117.78.19.76",Port:4243,Method:"reg/pull",Params:"{\"local_tag\":\"10.122.75.228:5000/centostest:latest\"}"}
	//post_data:=RequestData{Version:"1.0",ServerIP:"117.78.19.76",Port:4243,Method:"image/rmi",Params:"{\"local_tag\":\"10.122.75.228:5000/centostest:latest\"}"}
	//post_data := RequestData{Version: "1.0", ServerIP: "10.122.75.228", Port: 5000, Method: "reg/delete", Params: "{\"repository\":\"centos\",\"tags\":\"latest\"}"}
	post_data := RequestData{Version: "1.0", ServerIP: "117.78.19.76", Port: 4243, Method: "reg/pushone", Params: "{\"local_tag\":\"centos:latest\",\"remote_tag\":\"10.122.75.228:5000/centostest:latest\"}"}
	//post_data:=RequestData{Version:"1.0",ServerIP:"117.78.19.76",Port:4243,Method:"reg/pullone",Params:"{\"local_tag\":\"10.122.75.228:5000/centostest:latest\"}"}
	strPostData, _ := json.Marshal(post_data)
	strTemp := "request=" + string(strPostData)
	resp, err := http.Post("http://127.0.0.1:8080/v1/reg/pushone",
		"application/x-www-form-urlencoded", strings.NewReader(strTemp))
	//"application/json",strings.NewReader(strTemp))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func httpPostForm() {
	resp, err := http.PostForm("http://127.0.0.1:8080/v1/version",
		url.Values{"key": {"Value"}, "id": {"123"}})

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}

func httpDo() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://127.0.0.1:8080/v1/version", strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

//

func GetImage(strLocalTag string) (strImage string, strTag string) {
	var strName string
	nPos := strings.Index(strLocalTag, "/")
	strName = common.SubstrAfter(strLocalTag, nPos)
	nPos = strings.Index(strName, ":")
	strImage = common.SubstrBefore(strName, nPos)
	strTag = common.SubstrAfter(strName, nPos)
	return strImage, strTag
}

func main() {
	//strImage,strTag:=GetImage("10.122.75.228:5000/centostest:latest")
	//fmt.Println("strImage=", strImage)
	//fmt.Println("strTag=", strTag)
	//return
	//httpGet()
	httpPost()
	//httpPostForm()
	//httpDo()
}
