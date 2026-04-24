package compiler

// for compiling a source file and all its dependencies
func Compile(filename string) {
	//read source file
	content := ReadFileToString(filename)
	srcFile := SourceFile{
		Filename: filename,
		Content:  content,
	}

}
