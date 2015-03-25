package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//get,post,postform
func GetHTTP(url string) (ret string, flag bool) {
	fmt.Println("HTTP GET : ", url)
	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		return err.Error(), false
	}
	if response.StatusCode != 200 {
		return string(response.StatusCode), false
	}
	body, _ := ioutil.ReadAll(response.Body)
	return string(body), true
}

//http.Client和http.NewRequest
func RequestHTTP() {
	client := &http.Client{}
	reqest, _ := http.NewRequest("GET", "http://www.baidu.com", nil)

	reqest.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	reqest.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
	reqest.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	reqest.Header.Set("Cache-Control", "max-age=0")
	reqest.Header.Set("Connection", "keep-alive")

	response, _ := client.Do(reqest)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		bodystr := string(body)
		fmt.Println(bodystr)
	}
}

//func main() {
//	GetRegList()
//}
