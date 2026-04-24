package compiler

import "fmt"

// for compiling a source file and all its dependencies
func Compile(filename string) {
	//read source file
	content := ReadFileToString(filename)
	srcFile := SourceFile{
		Filename: filename,
		Content:  content,
	}

	//get tokens
	tokens := GetTokens(srcFile)

	//debug: print tokens
	fmt.Println("=== TOKENS ===")
	for i, token := range tokens {
		fmt.Printf("[%d] Type: %-15s | Value: %-10s | Index: %d\n", i, token.Type, "'"+token.Value+"'", token.Index)
	}
	fmt.Println("==============")
	//debug end

	//parse tokens to AST

}
