package eval_test

import (
	"monkey/eval"
	"monkey/eval/object"
	"monkey/lexer"
	"monkey/parser"
	"testing"

	"github.com/stretchr/testify/require"
)

func evalProgram(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return eval.Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int) {
	r := require.New(t)
	result, ok := obj.(*object.Integer)

	r.Truef(ok, "obj is not *object.Integer, got %T (%v)", obj, obj)

	r.Equal(expected, result.Value)
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) {
	r := require.New(t)
	result, ok := obj.(*object.Boolean)

	r.Truef(ok, "obj is not *object.Boolean, got %T (%v)", obj, obj)
	r.Equal(expected, result.Value)
}
