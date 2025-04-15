package eval_test

import (
	"fmt"
	"testing"
)

func TestArrayLiteral(t *testing.T) {
	input := `[1, true, "hello", 3 + 4]`

	evaluated := evalProgram(input)

	array := assertIsArrayObject(t, evaluated)

	testObject(t, array.Elements[0], 1)
	testObject(t, array.Elements[1], true)
	testObject(t, array.Elements[2], "hello")
	testObject(t, array.Elements[3], 7)
}

func TestIndexingOperator(t *testing.T) {
	array := `[1, true, "hello", 3 + 4]`

	expectedOutput := map[string]any{
		"0":     1,
		"1":     true,
		"2":     "hello",
		"3":     7,
		"1 + 2": 7,
	}

	expectedError := map[string]string{
		"-1":    "index out of bound: -1",
		"4":     "index out of bound: 4",
		`"xyz"`: "unsupported operation: ARRAY `INDEX` STRING",
		`false`: "unsupported operation: ARRAY `INDEX` BOOLEAN",
	}

	for i, exp := range expectedOutput {
		input := fmt.Sprintf("%s[%s]", array, i)
		t.Run(fmt.Sprintf("array[%s]", i), func(t *testing.T) {
			evaluated := evalProgram(input)
			testObject(t, evaluated, exp)
		})
	}

	for i, exp := range expectedError {
		input := fmt.Sprintf("%s[%s]", array, i)
		t.Run(fmt.Sprintf("array[%s]", i), func(t *testing.T) {
			evaluated := evalProgram(input)
			testErrorObject(t, evaluated, exp)
		})
	}
}
