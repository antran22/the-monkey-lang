package eval_test

import (
	"fmt"
	"testing"
)

// Array

func TestArrayLiteral(t *testing.T) {
	input := `[1, true, "hello", 3 + 4]`

	evaluated := evalProgram(t, input)

	array := assertIsArrayObject(t, evaluated)

	testObject(t, array.Elements[0], 1)
	testObject(t, array.Elements[1], true)
	testObject(t, array.Elements[2], "hello")
	testObject(t, array.Elements[3], 7)
}

func TestArrayIndexingOperator(t *testing.T) {
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
			evaluated := evalProgram(t, input)
			testObject(t, evaluated, exp)
		})
	}

	for i, exp := range expectedError {
		input := fmt.Sprintf("%s[%s]", array, i)
		t.Run(fmt.Sprintf("array[%s]", i), func(t *testing.T) {
			evaluated := evalProgram(t, input)
			testErrorObject(t, evaluated, exp)
		})
	}
}

// Hash

func TestHashLiterals(t *testing.T) {
	happyCases := []happyTestCase{
		{
			`
      let two = "two";
      let hash = {
        "one": 10 - 9,
        two: 1 + 1,
        "thr" + "ee": 6 / 2,
        4: 4,
        true: 5,
        false: 6
      }
    `,
			map[any]any{
				"one":   1,
				"two":   2,
				"three": 3,
				4:       4,
				true:    5,
				false:   6,
			},
		},
	}

	testExpressionEvaluation(t, happyCases, []errorTestCase{})
}

func TestHashIndexingOperator(t *testing.T) {
	commonSetup := `
      let two = "two";
      let hash = {
        "one": 10 - 9,
        two: 1 + 1,
        "thr" + "ee": 6 / 2,
        4: 4,
        true: 5,
        false: 6
      };
  `

	happyCases := []happyTestCase{
		{
			commonSetup + `hash["one"];`,
			1,
		},
		{
			commonSetup + `hash["two"];`,
			2,
		},
		{
			commonSetup + `hash["three"];`,
			3,
		},
		{
			commonSetup + `hash[4];`,
			4,
		},
		{
			commonSetup + `hash[true];`,
			5,
		},
		{
			commonSetup + `hash[false];`,
			6,
		},
		{
			commonSetup + `hash[5];`,
			nil,
		},
	}

	testExpressionEvaluation(t, happyCases, []errorTestCase{})
}
