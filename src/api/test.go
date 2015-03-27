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
