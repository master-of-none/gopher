package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFileName := flag.String("csv", "problem.csv", "must be csv file")

	flag.Parse()

	file, err := os.Open(*csvFileName)

	if err != nil {
		exit(fmt.Sprintf("Failed to open this file. Error is %v\n", err))
		os.Exit(1)
	}

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()

	if err != nil {
		exit(fmt.Sprintf("Failed to read CSV file %v\n", err))
	}

	problems := parseLines(lines)
	fmt.Println(problems)
}

type problem struct {
	question string
	answer   string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   line[1],
		}
	}

	return ret
}
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
