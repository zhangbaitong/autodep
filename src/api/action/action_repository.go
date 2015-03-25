package action

import (
	"client"
	"fmt"
)

const (
	SEARCH             = "http://10.122.75.228:5000/v1/search"
	TAGS               = "http://10.122.75.228:5000/v1/repositories/%s/%s/tags"
	NAMESPACE_DEFAULT  = "library"
	REPOSITORY_DEFAULT = "registry"
)

func ActionRegList() string {
	ret, flag := client.GetHTTP(SEARCH)
	fmt.Println(ret)
	if !flag {
		ret = ""
	}
	return ret
}

func ActionRegTags(ns []string, rep []string) string {
	var url string
	if len(ns) > 0 && len(rep) > 0 {
		url = fmt.Sprintf(TAGS, ns[0], rep[0])
	} else {
		url = fmt.Sprintf(TAGS, NAMESPACE_DEFAULT, REPOSITORY_DEFAULT)
	}
	ret, flag := client.GetHTTP(url)
	fmt.Println(ret)
	if !flag {
		ret = ""
	}
	return ret
}
