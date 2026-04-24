package compiler

// for compiling a source file and all its dependencies
func Compile(filename string) {
	srcFiles := ResolveImports(filename)
	outputFiles := make([]SourceFile, len(srcFiles))
	for i, srcFile := range srcFiles {
		outputFiles[i] = CompileFile(srcFile)
	}

}

// for compiling a single source file
func CompileFile(srcFile SourceFile) SourceFile {
	var outputFile SourceFile
	return outputFile
}
