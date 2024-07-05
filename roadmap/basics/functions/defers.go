package main

import "fmt"

func trace(s string) string {
	fmt.Println("Entering here: ", s)
	return s
}

func un(s string) {
	fmt.Println("Leaving: ", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("Inside a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("Inside b")
	a()
}
func main() {
	b()
}
