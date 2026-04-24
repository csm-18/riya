package compiler

import (
	"fmt"
	"os"
)

func ReadFileToString(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: \n %v\n", err)
		os.Exit(1)
	}
	return string(data)
}

func LinePositionFromIndex(content string, index int) (line int, column int) {
	line = 1
	column = 1

	for i := 0; i < index && i < len(content); i++ {
		if content[i] == '\n' {
			line++
			column = 1
		} else {
			column++
		}
	}

	return line, column
}

func PrintErrorWithPosition(srcFile SourceFile, index int, message string) {
	line, column := LinePositionFromIndex(srcFile.Content, index)
	fmt.Printf("Error in file '%s' at line %d, column %d:\n %s\n", srcFile.Filename, line, column, message)
	os.Exit(1)
}
