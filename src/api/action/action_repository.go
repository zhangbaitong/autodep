package action

import (
	"api/common"
	"client"
	"fmt"
)

const (
	SEARCH             = "http://%s:5000/v1/search"
	TAGS               = "http://%s:5000/v1/repositories/%s/%s/tags"
	NAMESPACE_DEFAULT  = "library"
	REPOSITORY_DEFAULT = "registry"
)

var ip string

func init() {
	ip, _ = common.Config().String("autodep", "ip")
	common.Log().Println("common inint ip - ", ip)
}

func ActionRegList() string {
	ret, flag := client.GetHTTP(fmt.Sprintf(SEARCH, ip))
	fmt.Println(ret)
	if !flag {
		ret = ""
	}
	return ret
}

func ActionRegTags(ns []string, rep []string) string {
	var url string
	if len(ns) > 0 && len(rep) > 0 {
		url = fmt.Sprintf(TAGS, ip, ns[0], rep[0])
	} else {
		url = fmt.Sprintf(TAGS, ip, NAMESPACE_DEFAULT, REPOSITORY_DEFAULT)
	}
	ret, flag := client.GetHTTP(url)
	fmt.Println(ret)
	if !flag {
		ret = ""
	}
	return ret
}

func ActionRegSearch(q []string, n []string, page []string) string {
	if len(q) == 0 || len(n) == 0 || len(page) == 0 {
		return ActionRegList()
	}
	//ip, _ := common.Config().String("autodep", "ip")
	url := fmt.Sprintf(SEARCH+"?q=%s&page=%s&n=%s", ip, q[0], page[0], n[0])
	ret, flag := client.GetHTTP(url)
	fmt.Println(ret)
	if !flag {
		ret = ""
	}
	return ret
}
