package main

import (
	"bytes"
	"catfish/cmd"
	"os"
	"path/filepath"
	"testing"
)

func readFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func runCatCommand(args []string) (string, error) {

	buf := new(bytes.Buffer)
	catCmd := cmd.GetCatCmd()
	catCmd.SetOut(buf)
	catCmd.SetErr(buf)
	catCmd.SetArgs(args[1:])

	err := catCmd.Execute()
	return buf.String(), err
}

func TestCatCommand(t *testing.T) {
	tests := []struct {
		inputFile    string
		expectedFile string
		args         []string
		description  string
	}{
		{"file1.txt", "file1.txt", []string{"catfish", "cat", "input/file1.txt"}, "cat without flags"},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			expected, err := readFile(filepath.Join("expected", test.expectedFile))
			if err != nil {
				t.Fatalf("failed to read expected file: %v", err)
			}

			output, err := runCatCommand(test.args)
			if err != nil {
				t.Fatalf("failed to run cat command: %v", err)
			}

			if output != expected {
				t.Errorf("output does not match expected\nOutput:\n%s\nExpected:\n%s", output, expected)
			}
		})
	}
}
