package main

import (
	"fmt"
	"os"
	"strconv"
)

func exercise1() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func exercise2() {
	var s string
	for i := 1; i < len(os.Args); i++ {
		s += "\n" + strconv.Itoa(i) + " " + os.Args[i]
	}
	fmt.Println(s)
}
func main() {
	exercise1()
	exercise2()
}
