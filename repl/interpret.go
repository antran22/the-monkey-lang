package repl

import (
	"fmt"
	"monkey/eval"
	"monkey/eval/object"
	"monkey/lexer"
	"monkey/parser"
	"os"
)

func Interpret(script string) {
	file, err := os.ReadFile(script)
	if err != nil {
		panic(err)
	}

	l := lexer.New(string(file))
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		printParserErrors(os.Stdout, p.Errors())
	}

	env := object.NewEnvironment()
	evaluated := eval.Eval(program, env)

	if object.IsError(evaluated) {
		fmt.Println(evaluated.Inspect())
	}
}
