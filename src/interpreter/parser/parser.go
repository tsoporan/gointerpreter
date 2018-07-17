/*
  Parser, the parser is responsible for building out the AST based on the input
  Also called "syntactic anaylsis"
*/

package parser

import (
	"fmt"
	"interpreter/ast"
	"interpreter/lexer"
	"interpreter/token"
)

type Parser struct {
	l *lexer.Lexer // Pointer to instance of Lexer

	curToken  token.Token
	peekToken token.Token

	errors []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{}, // Empty init errors
	}

	// Read in first two tokens to set both cur and next
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("Expected next token to be %s, got %s instead.", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

/*
  Constructs the program statements
*/
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

/*
  Parse statement based on token type
*/
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
  case token.RETURN:
    return p.parseReturnStatement()
	default:
		return nil
	}
}

/*
  Parse a LET statement
*/
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Value}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

/*
  Parse a RETURN statement
*/
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
  stmt := &ast.ReturnStatement{Token: p.curToken}

  p.nextToken()

  for !p.curTokenIs(token.SEMICOLON) {
    p.nextToken()
  }

  return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

/*
  Assertion to ensure that the next token is expected, otherwise
  this results in a parsing error
*/
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}
