package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/divxvid/monkey-interpreter/lexer"
	"github.com/divxvid/monkey-interpreter/token"
)

const PROMPT = ">>> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		hasInput := scanner.Scan()

		if !hasInput {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
