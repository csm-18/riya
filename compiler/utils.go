package compiler

import (
	"fmt"
	"os"
	"strings"
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

func DebugPrintTokens(tokens []Token) {
	fmt.Println("=== TOKENS ===")
	for i, token := range tokens {
		fmt.Printf("[%d] Type: %-15s | Value: %-10s | Index: %d\n", i, token.Type, "'"+token.Value+"'", token.Index)
	}
	fmt.Println("==============")
}

func DebugPrintAST(node ASTNode) {
	fmt.Println("=== AST ===")
	printASTNode(node, 0)
	fmt.Println("===========")
}

func printASTNode(node ASTNode, indent int) {
	indentStr := strings.Repeat("  ", indent)

	switch n := node.(type) {
	case FileNode:
		fmt.Printf("%sFileNode: %s\n", indentStr, n.Filename)
		for _, child := range n.Children {
			printASTNode(child, indent+1)
		}
	case FunctionNode:
		fmt.Printf("%sFunctionNode: %s\n", indentStr, n.Name)
		fmt.Printf("%s  Parameters:\n", indentStr)
		for _, param := range n.Parameters {
			printASTNode(param, indent+2)
		}
		fmt.Printf("%s  ReturnTypes: %v\n", indentStr, n.ReturnTypes)
		fmt.Printf("%s  Body:\n", indentStr)
		for _, stmt := range n.Body {
			printASTNode(stmt, indent+2)
		}
	case FunctionParameterNode:
		fmt.Printf("%sParameter: %s (%s)\n", indentStr, n.Name, n.Type)
	default:
		fmt.Printf("%sUnknown AST Node\n", indentStr)
	}
}
