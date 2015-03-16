package action

import "github.com/codeskyblue/go-sh"

func Actionls() string {
	sh.Command("echo", "hello\tworld").Command("cut", "-f2").Run()
	sh.Command("ls").Run()
	return "go-sh"
}
