package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	s, sep := "",""
	for i, arg := range os.Args[1:]{
		s += sep + arg
		sep = " "
		fmt.Println("index: ", i)
		fmt.Println("arg: " + arg)
	}
	fmt.Println(s)
	fmt.Println(os.Args[0])
	fmt.Println(strings.Join(os.Args[1:], " "))
}
