package parser

import (
	"testing"

	"basiclang/ast"
	"basiclang/lexer"
	"basiclang/token"
)

func TestLestStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 3241299;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatal("p.ParserProgram() retrned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got: %d \n",
			len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}
	for i, tt := range tests {
		statement := program.Statements[i]
		testLestStatement(t, statement, tt.expectedIdentifier)
	}
}

func testLestStatement(t *testing.T, s ast.Statement, name string) {
	if s.TokenLiteral() != "let" {
		t.Errorf("statement name should be 'let', got: %v", s.TokenLiteral())
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s it not *ast.LetStatement. got %T\n", s)
	}

	if letStmt.Token.Type != token.LET {
		t.Errorf("expected TokenType is LET, got: %s\n", letStmt.Token.Type)
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value is not  %s, got: %s\n", name, letStmt.Name.Value)
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got: %s", name, letStmt.Name)
	}
}
