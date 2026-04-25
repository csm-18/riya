package compiler

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

	DebugPrintTokens(tokens)

	//parse tokens to AST
	ast := ParseTokensToAST(tokens, srcFile)

	DebugPrintAST(ast)

	//semantic analysis
	SemanticAnalyzer(ast)

	//code generation
	CodeGenerator(ast)
}
