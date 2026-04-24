// lexical analysis and tokenization of source code
package compiler

import "unicode"

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
		} else if unicode.IsLetter(rune(srcFile.Content[x])) {
			identifier := ""

			y := x
			for y < len(srcFile.Content) && (unicode.IsLetter(rune(srcFile.Content[y])) || unicode.IsDigit(rune(srcFile.Content[y]))) || srcFile.Content[y] == '_' {
				identifier += string(srcFile.Content[y])
				y += 1
			}

			isKeyword := false
			z := 0
			for z < len(Keywords) {
				if identifier == Keywords[z] {
					isKeyword = true
					break
				}
				z += 1
			}
			if isKeyword {
				token := Token{Type: TokenTypes["Keyword"], Value: identifier, Index: x}
				tokens = append(tokens, token)
			} else {
				token := Token{Type: TokenTypes["Identifier"], Value: identifier, Index: x}
				tokens = append(tokens, token)
			}

			x = y
			continue
		} else if srcFile.Content[x] == ' ' || srcFile.Content[x] == '\n' || srcFile.Content[x] == '\t' {
			x += 1
			continue
		} else {
			PrintErrorWithPosition(srcFile, x, "Undefined token!")
		}
		x += 1
	}
	return tokens
}
