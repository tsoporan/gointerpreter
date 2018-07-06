/*
  Read-Eval-Print-Loop
*/

package repl

import (
	"bufio"
	"fmt"
	"interpreter/lexer"
	"interpreter/token"
	"io"
)

const PROMPT = ">>> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for { // While
		fmt.Printf(PROMPT)

		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		lex := lexer.New(line)

    // Iterate and print tokens until meeting end of file
		for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
