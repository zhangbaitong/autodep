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

func ActionRegSearch(q []string, n []string, page []string) string {
	if len(q) == 0 || len(n) || 0 && len(page) || 0 {
		return ActionRegList()
	}

	url := fmt.Sprintf(SEARCH+"?q=%s&page=%s&n=%s", q[0], page[0], n[0])
	ret, flag := client.GetHTTP(url)
	fmt.Println(ret)
	if !flag {
		ret = ""
	}
	return ret
}
