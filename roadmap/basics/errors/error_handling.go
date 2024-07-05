package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("filename.txt")

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Opened a file %v", f)
	}
}
