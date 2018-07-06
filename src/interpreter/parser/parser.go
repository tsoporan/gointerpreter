/*
  Parser, the parser is responsible for building out the AST based on the input
  Also called "syntactic anaylsis"
*/

package parser

import (
  "interpreter/ast"
  "interpreter/lexer"
  "interpreter/token"
)

type Parser struct {
  l *lexer.Lexer // Pointer to instance of Lexer

  curToken token.Token
  peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
  p := &Parser{l: l}

  // Read in two tokens to set both cur and next
  p.nextToken()
  p.nextToken()

  return p
}

func (p *Parser) nextToken() {
  p.curToken = p.peekToken
  p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
  return nil
}
