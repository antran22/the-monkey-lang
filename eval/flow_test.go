package eval_test

import (
	"fmt"
	"testing"
)

func TestReturnStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"return 10;", 10},
		{"return 10; 9", 10},
		{"return 2 * 5; 9", 10},
		{"9; return 2 * 5; 9", 10},
		{
			`
if (10 > 1) {
  if (10 > 1) {
    return 10;
  }
  return 1;
}`, 10,
		},
	}

	for _, tc := range tests {
		evaluated := evalProgram(tc.input)
		testIntegerObject(t, evaluated, tc.expected)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{"5 + true;", "unsupported operation: INTEGER + BOOLEAN"},
		{"5 + true; 5", "unsupported operation: INTEGER + BOOLEAN"},
		{"-true", "unsupported operation: - BOOLEAN"},
		{"true + false", "unsupported operation: BOOLEAN + BOOLEAN"},
		{"5; true + false; 5", "unsupported operation: BOOLEAN + BOOLEAN"},
		{"if (10 > 1) { true + false; }", "unsupported operation: BOOLEAN + BOOLEAN"},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("error handling %s", tc.input), func(t *testing.T) {
			evaluated := evalProgram(tc.input)
			testErrorObject(t, evaluated, tc.expectedMessage)
		})
	}
}
