package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/eval"
	"monkey/eval/object"
	"monkey/lexer"
	"monkey/parser"
	"os/user"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(in)

	env := object.NewEnvironment()

	fmt.Fprintf(out, "Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Fprintf(out, "Type in some commands please\n")
	for {
		fmt.Fprint(out, PROMPT)

		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := eval.Eval(program, env)

		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")

	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
