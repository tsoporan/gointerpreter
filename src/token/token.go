/*
 * A token is effectively a string that can be easily categorized
 */

package token

type TokenType string

type Token struct {
	Type    TokenType
	Value   string
}

// Define token types
const (
	ILLEGAL = "ILLEGAL" // Fallback when no match
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // x, y, add
	INT   = "INT" // 1

	// Operators
	ASSIGN = "="
	PLUS   = "+"

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
)