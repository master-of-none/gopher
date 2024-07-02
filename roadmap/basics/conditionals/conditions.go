package main

import "fmt"

func main() {
	nums := [5]int{1, 2, 3, 4, 5}

	for _, n := range nums {
		if n%2 == 0 {
			fmt.Printf("%d is Even\n", n)
		} else {
			fmt.Printf("%d is odd\n", n)
		}
	}
}
