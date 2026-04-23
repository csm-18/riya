package compiler

import "fmt"

func Compile(filename string) {
	source := ReadFileToString(filename)
	fmt.Println(source)
}
