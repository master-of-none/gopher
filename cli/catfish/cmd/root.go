/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// catCmd represents the cat command
var catfish = &cobra.Command{
	Use:   "cat [filename]",
	Short: "Display contents of file",
	Long:  `A simple implementation of the cat command to display the content of a file.`,
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		showNonBlankLineNum, _ := cmd.Flags().GetBool("number-notblank")
		showLineNum, _ := cmd.Flags().GetBool("number")
		file, err := os.Open(filename)

		if err != nil {
			fmt.Printf("Error in opening the file %v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		lineNumber := 1
		for scanner.Scan() {
			// fmt.Println(scanner.Text())
			line := scanner.Text()
			if showNonBlankLineNum && len(line) > 0 {
				fmt.Printf("%5d  %s\n", lineNumber, line)
				lineNumber++
			} else if showLineNum {
				fmt.Printf("%5d %s\n", lineNumber, line)
				lineNumber++
			} else if !showNonBlankLineNum {
				fmt.Println(line)
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error in reading the file %v\n", err)
		}
	},
}

func GetCatCmd() *cobra.Command {
	return catfish
}

func Execute() {
	if err := catfish.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
