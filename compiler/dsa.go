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

var Keywords = []string{"fun"}

// AST (Abstract Syntax Tree) data structures
type ASTNode interface {
	node()
}

type FileNode struct {
	Filename string
	Children []ASTNode
}

func (n FileNode) node() {}

type FunctionNode struct {
	Name        string
	Parameters  []FunctionParameterNode
	ReturnTypes []string
	Body        []ASTNode
}

func (n FunctionNode) node() {}

type FunctionParameterNode struct {
	Name string
	Type string
}

func (n FunctionParameterNode) node() {}
