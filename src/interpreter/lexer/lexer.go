/*
 * The lexer is responsible for the "lexical analysis", turning the source code into tokens, also known
 * as tokenizer or scanner.
 */

package lexer

import "interpreter/token"

type Lexer struct {
	input            string
	position         int  // Points to current char
	nextReadPosition int  // Points to next char
	current          byte // Current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	// Only supports ascii for now
	if l.nextReadPosition >= len(l.input) {
		l.current = 0 // End
	} else {
		l.current = l.input[l.nextReadPosition] // set to the next char
	}

	// Move position pointers
	l.position = l.nextReadPosition
	l.nextReadPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.current {
	case '=':
		tok = newToken(token.ASSIGN, l.current)
	case ';':
		tok = newToken(token.SEMICOLON, l.current)
	case '(':
		tok = newToken(token.LPAREN, l.current)
	case ')':
		tok = newToken(token.RPAREN, l.current)
	case ',':
		tok = newToken(token.COMMA, l.current)
	case '+':
		tok = newToken(token.PLUS, l.current)
	case '{':
		tok = newToken(token.LBRACE, l.current)
	case '}':
		tok = newToken(token.RBRACE, l.current)
  case '!':
    tok = newToken(token.BANG, l.current)
  case '-':
    tok = newToken(token.MINUS, l.current)
  case '/':
    tok = newToken(token.SLASH, l.current)
  case '*':
    tok = newToken(token.ASTERISK, l.current)
  case '<':
    tok = newToken(token.LT, l.current)
  case '>':
    tok = newToken(token.GT, l.current)
	case 0:
		tok.Value = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.current) {
			tok.Value = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Value)
			return tok
		} else if isDigit(l.current) {
			tok.Value = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.current)
		}
	}

	l.readChar()

	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Value: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	start := l.position

	for isLetter(l.current) {
		l.readChar()
	}

	return l.input[start:l.position]
}

func (l *Lexer) readNumber() string {
	start := l.position

	for isDigit(l.current) {
		l.readChar()
	}

	return l.input[start:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) skipWhiteSpace() {
	// cur := l.current

	// Skip over any whitespace
	for l.current == ' ' || l.current == '\t' || l.current == '\n' || l.current == '\r' {
		l.readChar()
	}
}
