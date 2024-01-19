package ast

import (
	"bytes"

	"basiclang/token"
)

// Basic building block of the AST
type Node interface {
	TokenLiteral() string
	String() string
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

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// ensuring interface implementations
var (
	_ Node       = (*Identifier)(nil)
	_ Expression = (*Identifier)(nil)
	_ Statement  = (*LetStatement)(nil)
	_ Statement  = (*ReturnStatement)(nil)
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

func (i *Identifier) String() string {
	return i.Value
}

// this implements the Statement interface
type LetStatement struct {
	// the token that we've got from the parser
	Token token.Token // token.LET
	// name is the actual statement that we've parsed out from the source code.
	Name *Identifier
	// value is the identifier next to it or for this case the name of the variable.
	Value Expression
}

func (ls *LetStatement) statementNode() {
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// this implements the Statement interface
type ReturnStatement struct {
	Token token.Token // token.RETURN
	Name  *Identifier
	Value Expression
}

func (rs *ReturnStatement) statementNode() {
}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")
	if rs.Value != nil {
		out.WriteString(rs.Value.String())
	}
	out.WriteString(";")

	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {
}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression == nil {
		return ""
	}
	return es.Expression.String()
}
