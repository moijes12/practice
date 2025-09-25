// Check for duplicate lines and print lines and files
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lineCount := make(map[string]int)  // A frequency counter for lines
	fileMap := make(map[string]string) // A map of lines against the files in which they occur

	files := os.Args[1:]
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
		}
		countLines(f, lineCount, fileMap)
		f.Close()
	}

	for line, n := range lineCount {
		if n > 1 {
			fmt.Printf("line = %s\tCount = %d\tfiles = %s\n", line, n, fileMap[line])
		}
	}
}

func countLines(f *os.File, lineCount map[string]int, fileMap map[string]string) {
	input := bufio.NewScanner(f)
	fileName := f.Name()
	for input.Scan() {
		line := input.Text()
		lineCount[line]++
		if lineCount[line] == 1 {
			fileMap[line] = fileName
		} else {
			fileMap[line] += "," + fileName
		}
	}
}
