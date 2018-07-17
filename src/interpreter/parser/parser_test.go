/*
  Parser tests
*/

package parser

import (
	"interpreter/ast"
	"interpreter/lexer"
	"testing"
)

type ExpectedIdentifier struct {
	name string
}

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 6;
let foo = 121212;
  `

	Lexer := lexer.New(input)
	Parser := New(Lexer)

	program := Parser.ParseProgram()
	checkParserErrors(t, Parser)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	statementCount := len(program.Statements)

	if statementCount != 3 {
		t.Fatalf("program.Statements does not contain three statements, got=%d", statementCount)
	}

	tests := []ExpectedIdentifier{
		{"x"},
		{"y"},
		{"foo"},
	}

	for i, expectedIdent := range tests {
		stmt := program.Statements[i]

		if !testLetStatement(t, stmt, expectedIdent.name) {
			return
		}
	}

}

func TestReturnStatements(t *testing.T) {
	input := `
return 1;
return 2;
return 123123131;
`

	Lexer := lexer.New(input)
	Parser := New(Lexer)

	program := Parser.ParseProgram()
	checkParserErrors(t, Parser)

	statementCount := len(program.Statements)

	if statementCount != 3 {
		t.Fatalf("program.Statements does not contain three statements, got=%d", statementCount)
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement, got=%T", stmt)
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got=%q", returnStmt.TokenLiteral())
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let', got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("s not *ast.LetStatement, got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s', got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s', got=%s", name, letStmt.Name)
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("Parser has %d errors", len(errors))

	for _, msg := range errors {
		t.Errorf("Parser error: %q", msg)
	}

	t.FailNow()
}
