/*
 * A token is effectively a string that can be easily categorized
 */

package token

type TokenType string

type Token struct {
	Type  TokenType
	Value string
}

// Define token types
const (
	ILLEGAL = "ILLEGAL" // Fallback when no match
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // x, y, add
	INT   = "INT"   // 1

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok { // Check if keyword
		return tok
	}
	return IDENT
}
