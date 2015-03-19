package main


import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "strings"
    "encoding/json"
)
type RequestData struct
{
        Version         string
        ServerIP        string
        Port             int
        Method         string
        Params         string
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
                post_data:=RequestData{"1.0","117.78.19.76",4243,"fig/create","TestValue"}
                strPostData, _ := json.Marshal(post_data)
                strPostData="{\"Version\":\"1.0\",\"ServerIP\":\"117.78.19.76\","Port":\"4243\",\"Method\":\"create_fig\",\"Params\":{\"fig_data\":\"nginx:\n  image: centos6\/nginx:2015-02-08\n  ports:\n    - \"8807:8807\"\n    - \"9090:9090\"\n  links:\n    - mysql\n  volumes:\n    - \/data\/esp\/code\/nginx:\/data\/esp\/code\/nginx\n    - \/data\/espreport\/code:\/data\/espreport\/code\n  command: #&fig_directory&#\/startup\/nginx\/start.sh\r\nmysql:\n  image: centos6\/mysql:2015-01-09\n  ports:\n    - \"3306:3306\"\n  volumes:\n    - \/data\/esp\/mysql:\/home\/databases\/mysql\/data\n    - \/home\/docker\/fig\/startup\/mysql:\/home\/docker\/fig\/startup\/mysql\n  command: #&fig_directory&#\/startup\/mysql\/start.sh\r\n\",\"commands\":{\"nginx\":\"#!\/bin\/bash\nservice nginx restart\ntail -f nginx\",\"mysql\":\"#!\/bin\/bash\nservice mysqld restart\ntail -f mysql\"}}}"
                strTemp:="request="+string(strPostData)
                //strTemp:=string(strPostData)
	resp, err := http.Post("http://127.0.0.1:8080/v1/fig/create",
		"application/x-www-form-urlencoded",strings.NewReader(strTemp))
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

func main() {
	//httpGet()
	httpPost()
	//httpPostForm()
	//httpDo()
}
