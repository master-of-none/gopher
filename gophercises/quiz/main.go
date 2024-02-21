package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFileName := flag.String("csv", "problem.csv", "must be csv file")

	flag.Parse()

	file, err := os.Open(*csvFileName)

	if err != nil {
		fmt.Printf("Failed to open this file. Error is %v\n", err)
		os.Exit(1)
	}
	_ = file
}
