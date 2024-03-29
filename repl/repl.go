package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/divxvid/monkey-interpreter/evaluator"
	"github.com/divxvid/monkey-interpreter/lexer"
	"github.com/divxvid/monkey-interpreter/object"
	"github.com/divxvid/monkey-interpreter/parser"
)

const PROMPT = ">>> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Print(PROMPT)
		hasInput := scanner.Scan()

		if !hasInput {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserError(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserError(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
