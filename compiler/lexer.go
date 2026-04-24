// lexical analysis and tokenization of source code
package compiler

func GetTokens(srcFile SourceFile) []Token {
	var tokens []Token

	x := 0
	for x < len(srcFile.Content) {
		if srcFile.Content[x] == '(' {
			token := Token{Type: TokenTypes["LeftParen"], Value: "(", Index: x}
			tokens = append(tokens, token)
			x += 1
			continue
		} else if srcFile.Content[x] == ')' {
			token := Token{Type: TokenTypes["RightParen"], Value: ")", Index: x}
			tokens = append(tokens, token)
			x += 1
			continue
		} else if srcFile.Content[x] == '{' {
			token := Token{Type: TokenTypes["LeftBrace"], Value: "{", Index: x}
			tokens = append(tokens, token)
			x += 1
			continue
		} else if srcFile.Content[x] == '}' {
			token := Token{Type: TokenTypes["RightBrace"], Value: "}", Index: x}
			tokens = append(tokens, token)
			x += 1
			continue
		}
		x += 1
	}
	return tokens
}
