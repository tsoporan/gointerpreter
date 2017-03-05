package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal String
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ","

	LPAREN  = "("
	RPARENT = ")"
	LBRACE  = "{"
	RBRACE  = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
