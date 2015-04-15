package main

import (
	"fmt"
)

func main() {
	test(1, 2)
}

func test(args ...int) {
	test2(args...)
}

func test2(args ...int) {
	fmt.Println(args[0])
	fmt.Println(args[1])
}

func test3() {
	session := sh.NewSession()
	session.ShowCMD = true
	out, err := session.Command("ssh", "117.78.19.76", "cd /home/docker/fig/o2omember && fig stop && fig rm").Output()
	fmt.Println("out=", string(out), "err=", err)
	fmt.Println("os out---", *os.Stdout)
}
