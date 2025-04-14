package eval_test

import (
	"fmt"
	"testing"
)

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"let a = 5; a;", 5},
		{"let a = 5 * 5; a;", 25},
		{"let a = 5; let b = a; b;", 5},
		{"let a = 5; let b = a; let c = a + b + 5; c;", 15},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("let statement %s", tc.input), func(t *testing.T) {
			evaluated := evalProgram(tc.input)
			testIntegerObject(t, evaluated, tc.expected)
		})
	}
}

func TestBindingErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{"foo", "identifier not found: foo"},
		{"let foo = foo", "identifier not found: foo"},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("error handling %s", tc.input), func(t *testing.T) {
			evaluated := evalProgram(tc.input)
			testErrorObject(t, evaluated, tc.expectedMessage)
		})
	}
}
