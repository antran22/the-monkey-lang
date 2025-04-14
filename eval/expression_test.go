package eval_test

import (
	"fmt"
	"testing"
)

func TestEvalBooleanOperator(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		// prefix expression
		{"!false", true},
		{"!true", false},
		{"!0", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},

		// infix compare expression
		{"1 < 2", true},
		{"1 > 2", false},
		{"2 <= 2", true},
		{"1 >= 2", false},
		{"1 == 1", true},
		{"1 != 2", true},

		{"true == true", true},
		{"true == false", false},
		{"true != false", true},
		{"true != true", false},

		{"true && true", true},
		{"true && false", false},
		{"false && false", false},
		{"false && true", false},

		{"true || true", true},
		{"true || false", true},
		{"false || true", true},
		{"false || false ", false},

		// mixed
		{"!(1 == 2)", true},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("eval bool %s", tc.input), func(st *testing.T) {
			evaluated := evalProgram(tc.input)
			testBooleanObject(t, evaluated, tc.expected)
		})
	}
}

func TestEvalIntegerExpression(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"5", 5},
		{"10", 10},
		// prefix expression
		{"-10", -10},
		{"--5", 5},
		// infix expression
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
		{"1 & 1", 1},
		{"1 & 0", 0},
		{"1 ^ 1", 0},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("eval int %d", tc.expected), func(t *testing.T) {
			evaluated := evalProgram(tc.input)
			testIntegerObject(t, evaluated, tc.expected)
		})
	}
}

func TestEvalStringExpression(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{`"hello"`, "hello"},
		{`"hello" + "hi"`, "hellohi"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("eval string %s", tc.expected), func(t *testing.T) {
			evaluated := evalProgram(tc.input)
			testStringObject(t, evaluated, tc.expected)
		})
	}
}
