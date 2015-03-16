package action

import (
	"bufio"
	_ "fmt"
	"os/exec"
)

func Actionecho() string {
	cmd := exec.Command("echo", "zhangbaitong")
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()
	r := bufio.NewReader(stdout)
	line, _, _ := r.ReadLine()
	//fmt.Println(string(line))
	return string(line)
}
