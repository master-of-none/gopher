package main

import "fmt"

func main() {
	var nums [4]int
	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	nums[3] = 4

	for n := range nums {
		fmt.Printf("%d", n)
	}
}
