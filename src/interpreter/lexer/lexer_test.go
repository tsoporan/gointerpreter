/* Test the Lexer */

package lexer

import (
	"interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	
	let add = fn(x, y) {
		x + y;
	};

	let result = add(five, ten);
  !-/*3;
  6 < 10 > 2
	`

	tests := []struct {
		expectedType  token.TokenType
		expectedValue string
	}{
		// Assign five
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		// Assign ten
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},

		// Function
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},

		// Result
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},


    // !-/*3;
    {token.BANG, "!"},
    {token.MINUS, "-"},
    {token.SLASH, "/"},
    {token.ASTERISK, "*"},
    {token.INT, "3"},
    {token.SEMICOLON, ";"},

    // 6 < 10 > 2
    {token.INT, "6"},
    {token.LT, "<"},
    {token.INT, "10"},
    {token.GT, ">"},
    {token.INT, "2"},

		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - TokenType wrong, expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Value != tt.expectedValue {
			t.Fatalf("tests[%d] - TokenValue wrong, expected=%q, got=%q", i, tt.expectedValue, tok.Value)

		}
	}
}
