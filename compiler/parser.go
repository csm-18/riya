// make ast from tokens
package compiler

func ParseTokensToAST(tokens []Token, srcFile SourceFile) ASTNode {
	var ast ASTNode

	ast = FileNode{
		Filename: srcFile.Filename,
		Children: nil,
	}

	x := 0
	for x < len(tokens) {
		if tokens[x].Type == TokenTypes["Keyword"] && tokens[x].Value == "fun" {
			//parse function definition
		} else {
			PrintErrorWithPosition(srcFile, tokens[x].Index, "Unexpected token: "+tokens[x].Value)
		}
		x += 1
	}

	return ast
}
