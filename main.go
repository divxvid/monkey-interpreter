package main

import (
	"os"

	"github.com/divxvid/monkey-interpreter/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
