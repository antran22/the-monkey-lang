package eval_test

import "testing"

func TestIfElseExpression(t *testing.T) {
	happyCases := []happyTestCase{
		{"if (true) {10}", 10},
		{"if (false) {10}", nil},
		{"if (1) {10}", 10},
		{"if (0) {10}", nil},
		{"if (1 < 2) {10}", 10},
		{"if (1 > 2) {10}", nil},
		{"if (1 < 2) {10} else {20}", 10},
		{"if (1 > 2) {10} else {20}", 20},
		{"if (1 > 2 || 3 < 4) {10} else {20}", 10},
		{"if (1 > 2 && 3 < 4) {10} else {20}", 20},
	}

	testExpressionEvaluation(t, happyCases, []errorTestCase{})
}
