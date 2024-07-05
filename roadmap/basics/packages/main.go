package main

import "fmt"
import "math1/math12"

func main() {
	xs := []float64{1, 2, 3, 4, 5}
	avg := math12.Average(xs)
	fmt.Println(avg)
}
