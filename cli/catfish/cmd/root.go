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
	Args:  cobra.ExactArgs(1),
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
	return rootCmd
}

func init() {
	rootCmd.AddCommand(catfish)

	// Here you will define your flags and configuration settings.
	catfish.Flags().BoolP("number-notblank", "b", false, "Number of non blank output lines")
	catfish.Flags().BoolP("number", "n", false, "Number output lines")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// catCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// catCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

