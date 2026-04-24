// this file contains the logic for resolving imports in the source code.
package compiler

func ResolveImports(filename string) []SourceFile {
	content := ReadFileToString(filename)
	srcFile := SourceFile{
		Filename: filename,
		Content:  content,
	}

	return []SourceFile{srcFile}
}
