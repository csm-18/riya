package compiler

// for compiling a source file and all its dependencies
func Compile(filename string) {
	srcFiles := ResolveImports(filename)
	outputFiles := make([]SourceFile, len(srcFiles))
	for i, srcFile := range srcFiles {
		outputFiles[i] = CompileFile(srcFile)
	}

	//create all output files
	//build exe from output files

}

// for compiling a single source file
func CompileFile(srcFile SourceFile) SourceFile {
	var outputFile SourceFile
	outputFile.Filename = srcFile.Filename[:len(srcFile.Filename)-4] + ".asm"

	return outputFile
}
