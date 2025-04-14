package eval_test

import "testing"

func TestIfElseExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected any
	}{
		{"if (true) {10}", 10},
		{"if (false) {10}", nil},
		{"if (1) {10}", 10},
		{"if (0) {10}", nil},
		{"if (1 < 2) {10}", 10},
		{"if (1 > 2) {10}", nil},
		{"if (1 < 2) {10} else {20}", 10},
		{"if (1 > 2) {10} else {20}", 20},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			evaluated := evalProgram(tc.input)
			if integer, ok := tc.expected.(int); ok {
				testIntegerObject(t, evaluated, integer)
			} else {
				testNullObject(t, evaluated)
			}
		})
	}
}
