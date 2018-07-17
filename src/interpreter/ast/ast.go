/*
  AST (Abstract Syntax Tree)
*/

package ast

import "interpreter/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node // Has to provide a TokenLiteral
	statementNode()
}

type Expression interface {
	Node // " "
	expressionNode()
}

type Identifier struct {
	Token token.Token // IDENT token
	Value string
}

type LetStatement struct {
	Token token.Token // LET token
	Name  *Identifier
	Value Expression
}

type ReturnStatement struct {
	Token       token.Token // RETURN token
	ReturnValue Expression
}

/*
  Each program is a series of statements, which is a slice of AST nodes which
	implement the Statement interface
*/
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

// Implements Expression, Identifiers can eventually produce values (RHS)
func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Value }

// Implements Statement
func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Value }

// Return Statment
func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Value }
