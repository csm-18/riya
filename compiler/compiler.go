package compiler

import "fmt"

// for compiling a source file and all its dependencies
func Compile(filename string) {
	srcFiles := ResolveImports(filename)

	if len(srcFiles) == 1 {
		CompileFile(srcFiles[0])
	}

}

// for compiling a single source file
func CompileFile(srcFile SourceFile) {
	fmt.Println("Compiling... ")
}
