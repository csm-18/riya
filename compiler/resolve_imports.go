package compiler

import "fmt"

func ResolveImports(filename string) {
	content := ReadFileToString(filename)
	srcFile := SourceFile{
		Filename: filename,
		Content:  content,
	}

	fmt.Println("resolving imports...")
	fmt.Printf("%+#v \n", srcFile)
}
