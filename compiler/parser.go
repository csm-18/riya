// make ast from tokens
package compiler

func ParseTokensToAST(tokens []Token, srcFile SourceFile) ASTNode {
	var ast = FileNode{
		Filename: srcFile.Filename,
		Children: []ASTNode{},
	}

	x := 0
	for x < len(tokens) {
		if tokens[x].Type == TokenTypes["Keyword"] && tokens[x].Value == "fun" {
			//parse function definition
			functionNode, newIndex := ParseFunctionNode(tokens[:], x, srcFile)
			ast.Children = append(ast.Children, functionNode)
			x = newIndex
			continue
		} else {
			PrintErrorWithPosition(srcFile, tokens[x].Index, "Unexpected token: "+tokens[x].Value)
		}
		x += 1
	}

	return ast
}

func ParseFunctionNode(tokens []Token, x int, srcFile SourceFile) (ASTNode, int) {
	newIndex := 0
	return nil, newIndex
}
