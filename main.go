package main

import (
	"flag"
	"monkey/repl"
	"os"
)

func main() {
	flag.Parse()

	if flag.NArg() > 0 {
		script := flag.Arg(0)
		repl.Interpret(script)
	} else {
		repl.Start(os.Stdin, os.Stdout)
	}
}
