// all data structures used in the compiler are defined here.
package compiler

type SourceFile struct {
	Filename string
	Content  string
}

type Token struct {
	Type  string
	Value string
	Index int
}

var TokenTypes = map[string]string{
	"Identifier":    "Identifier",
	"Keyword":       "Keyword",
	"StringLiteral": "StringLiteral",
	"NumberLiteral": "NumberLiteral",
	"Comma":         "Comma",
	"Semicolon":     "Semicolon",
	"LeftParen":     "LeftParen",
	"RightParen":    "RightParen",
	"LeftBrace":     "LeftBrace",
	"RightBrace":    "RightBrace",
}
