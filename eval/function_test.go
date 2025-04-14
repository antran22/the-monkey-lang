package eval_test

import (
	"fmt"
	"monkey/eval/object"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFunctionObject(t *testing.T) {
	r := require.New(t)

	input := "fn(x) {x + 2}"
	evaluated := evalProgram(input)

	fn, ok := evaluated.(*object.Function)
	r.Truef(ok, "fn is not *object.Function, got %T (%v)", evaluated, evaluated)

	r.Len(fn.FuncNode.Parameters, 1)

	r.Equal(fn.FuncNode.Parameters[0].String(), "x")

	r.Equal("{(x + 2)}", fn.FuncNode.Body.String())
}

func TestFunctionCall(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"let identity = fn(x) { x; }; identity(5);", 5},
		{"let identity = fn(x) { return x; }; identity(5);", 5},
		{"let double = fn(x) { x * 2; }; double(5);", 10},
		{"let add = fn(x, y) { x + y; }; add(5, 5);", 10},
		{"let add = fn(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20},
		{"fn(x) { x; }(5)", 5},
		{`
let adder = fn(x) {
  return fn(y) {
    return x + y;
  }
}
let addTwo = adder(2);
addTwo(5);
`, 7},

		{`
let applyFn= fn(func, a, b) {
  return func(a, b);
}
let multiply = fn(x, y) {
  return x * y;
}
applyFn(multiply, 5, 2);
`, 10},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("function call %s", tc.input), func(t *testing.T) {
			evaluated := evalProgram(tc.input)
			testIntegerObject(t, evaluated, tc.expected)
		})
	}
}
