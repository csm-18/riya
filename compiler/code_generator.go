package compiler

import "fmt"

func CodeGenerator(ast FileNode) {
	fmt.Println("Code generation started for file: " + ast.Filename)
}
