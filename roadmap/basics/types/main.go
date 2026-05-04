package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"math/bits"
	"os"
)

func Sha256() ([32]byte, [32]byte) {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	// fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	return c1, c2
}

// func reverse(s []int) {
// 	fmt.Printf("Original: %v\n", s)
// 	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
// 		s[i], s[j] = s[j], s[i]
// 	}
// 	fmt.Printf("Reversed: %v\n", s)
// }

func reverseAnyType[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseArrayPointer(arr *[5]int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func exercise1() int {
	c1, c2 := Sha256()
	count := 0

	for i, b1 := range c1 {
		xor := b1 ^ c2[i]

		count += bits.OnesCount8(xor)
	}
	return count
}

func exercise2() {
	hashflag := flag.String("hash", "sha256", "Choose hash function: sha256, sha384, sha512")
	flag.Parse()

	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading stdin: %v\n", err)
		return
	}

	switch *hashflag {
	case "sha256":
		sum := sha256.Sum256(input)
		fmt.Printf("%x\n", sum)

	case "sha384":
		sum := sha512.Sum384(input)
		fmt.Printf("%x\n", sum)

	case "sha512":
		sum := sha512.Sum512(input)
		fmt.Printf("%x\n", sum)

	default:
		fmt.Fprintf(os.Stderr, "Unknown Hash Type %s\n", *hashflag)
	}
}

func main() {
	// Sha256()
	// reverse([]int{1, 2, 3, 4})
	arr := []int{1, 2, 3, 4, 5}
	reverseAnyType(arr)

	ex1 := exercise1()
	fmt.Printf("Exercise 1 Answer: %d\n", ex1)

	arr2 := [5]int{1, 2, 3, 4, 5}
	reverseArrayPointer(&arr2)
	fmt.Printf("Reversed array (using pointer): %v\n", arr2)
	// exercise2()
}
