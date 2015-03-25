package action

import (
	"client"
	"fmt"
)

const (
	SEARCH             = "http://10.122.75.228:5000/v1/search"
	TAGS               = "http://10.122.75.228:5000/v1/repositories/"
	TAGS_END           = "/tags"
	NAMESPACE_DEFAULT  = "library/"
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

func ActionRegTags() string {
	ret, flag := client.GetHTTP(TAGS + NAMESPACE_DEFAULT + REPOSITORY_DEFAULT + TAGS_END)
	fmt.Println(ret)
	if !flag {
		ret = ""
	}
	return ret
}
