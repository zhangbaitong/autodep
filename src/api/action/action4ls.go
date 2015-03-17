package action

import (
	"fmt"
	"github.com/codeskyblue/go-sh"
)

func Actionls() string {
	session := sh.NewSession()
	session.ShowCMD = true
	err := session.Call("ssh","117.78.19.76", "touch ","tt.aa")
	if err != nil {
		fmt.Println("Server start faild error:", err)
	}
	return "go-sh"
}
