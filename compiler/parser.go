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
	x += 1 // skip 'fun' keyword

	var functionNode = FunctionNode{
		Name:        "",
		Parameters:  []FunctionParameterNode{},
		ReturnTypes: []string{},
		Body:        []ASTNode{},
	}
	if x < len(tokens) && tokens[x].Type != TokenTypes["Identifier"] {
		PrintErrorWithPosition(srcFile, tokens[x].Index, "Expected function name after 'fun'")
	} else {
		functionNode.Name = tokens[x].Value
		x += 1
	}

	//for now skip parameters and return types parsing, just check for '()' after function name
	if tokens[x].Type != TokenTypes["LeftParen"] && x+1 < len(tokens) && tokens[x+1].Type != TokenTypes["RightParen"] {
		PrintErrorWithPosition(srcFile, tokens[x].Index, "Expected '()' after function name")
	} else {
		x += 2
	}

	if tokens[x].Type != TokenTypes["LeftBrace"] {
		PrintErrorWithPosition(srcFile, tokens[x].Index, "Expected '{' to start function body")
	} else {
		x += 1
	}

	//for now skip function body parsing, just check for '}' at the end of function body
	if x < len(tokens) && tokens[x].Type != TokenTypes["RightBrace"] {
		PrintErrorWithPosition(srcFile, tokens[x].Index, "Expected '}' to end function body")
	} else {
		x += 1
	}

	return functionNode, x
}
