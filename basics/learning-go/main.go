package main

import "fmt"

func main() {
	// x := 10

	// if x > 5 {
	// 	fmt.Println(x)
	// 	x := 5
	// 	fmt.Println(x)
	// }
	// fmt.Println(x)

	evenVals := []int{0, 1, 2, 3, 4, 5}

	for i, v := range evenVals {
		fmt.Println(i, v)
	}
}
