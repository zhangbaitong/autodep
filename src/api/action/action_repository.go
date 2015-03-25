package action

import (
	"client"
)

const (
	SEARCH = "http://10.122.75.228:5000/v1/search"
)

func ActionRegList() string {
	ret, flag := GetHTTP(SEARCH)
	fmt.Println(ret)
	if !flag {
		ret = ""
	}
	return ret
}
