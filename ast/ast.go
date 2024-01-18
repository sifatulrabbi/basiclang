package ast

import (
	"basiclang/token"
)

// Basic building block of the AST
type Node interface {
	TokenLiteral() string
}

// programming statement
// i.e. `let msg = "hello world"`
type Statement interface {
	Node
	statementNode()
}

// programming Expression
// i.e. `5 + 5` or `add(5, 10)`
type Expression interface {
	Node
	expressionNode()
}

// This Program node is going to be the root node of every AST our parser produces.
// Every valid Monkey program is a series of statements.
// These statements are contained in the Program.Statements, which is just a slice of AST nodes that implement the Statement interface.
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// ensuring interface implementations
var (
	_ Node       = (*Identifier)(nil)
	_ Expression = (*Identifier)(nil)
	_ Statement  = (*LetStatement)(nil)
)

// this implements the Expression interface
type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) expressionNode() {
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// this implements the Statement interface
type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// this implements the Statement interface
type ReturnStatement struct {
	Token token.Token // token.RETURN
	Name  *Identifier
	Value Expression
}

func (ls *ReturnStatement) statementNode() {
}

func (ls *ReturnStatement) TokenLiteral() string {
	return ls.Token.Literal
}
