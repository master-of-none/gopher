package main

import "flag"

func main() {
	csvFileName := flag.String("csv", "problems.csv", "must be csv file")

	flag.Parse()
	_ = csvFileName

}
