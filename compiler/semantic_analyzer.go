package compiler

import (
	"fmt"
	"os"
)

func SemanticAnalyzer(ast FileNode) {
	if ast.NodeKind != "FileNode" {
		fmt.Println("Expected FileNode as AST root node!")
		os.Exit(1)
	}

	if len(ast.Children) == 0 {
		fmt.Println("No functions defined in the source file!")
		os.Exit(1)
	} else if len(ast.Children) > 1 {
		fmt.Println("Multiple functions defined in the source file! Only one function is allowed for now.")
		os.Exit(1)
	} else if ast.Children[0].(FunctionNode).Name != "main" {
		fmt.Println("The only function defined in the source file must be named 'main'!")
		os.Exit(1)
	}

	fmt.Println("Semantic analysis passed successfully!")

}
