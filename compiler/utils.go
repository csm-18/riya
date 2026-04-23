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
