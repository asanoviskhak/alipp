package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/asanoviskhak/alipp/src/lexer"
	"github.com/asanoviskhak/alipp/src/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		currentLine := scanner.Text()
		lexerInstance := lexer.New(currentLine)

		for tok := lexerInstance.NextToken(); tok.Type != token.EOF; tok = lexerInstance.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}