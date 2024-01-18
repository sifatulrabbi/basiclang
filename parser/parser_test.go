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

let x 5;
let = 10;
let 3241299;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatal("p.ParserProgram() retrned nil")
	}
	checkParserErrors(t, p)

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

func TestReturnStatements(t *testing.T) {
	input := `
return 5;
return 10;
return 1923423;
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatal("p.ParserProgram() retrned nil")
	}
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt is not *ast.ReturnStatement, got: %T\n", returnStmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral() is not 'return', got: %s\n",
				returnStmt.TokenLiteral())
		}
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

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors\n", len(errors))
	for _, msg := range errors {
		t.Errorf("parser errors: %q\n", msg)
	}
	t.FailNow()
}
