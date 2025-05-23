package eval_test

import (
	"testing"
)

func TestEvalBooleanOperator(t *testing.T) {
	testCases := []happyTestCase{
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

	testExpressionEvaluation(t, testCases, []errorTestCase{})
}

func TestEvalIntegerExpression(t *testing.T) {
	testCases := []happyTestCase{
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

	testExpressionEvaluation(t, testCases, []errorTestCase{})
}

func TestNullExpression(t *testing.T) {
	testCases := []happyTestCase{
		{`"hello" == null`, false},
		{`"hello" != null`, true},

		{`1 == null`, false},
		{`1 != null`, true},

		{`true == null`, false},
		{`true != null`, true},

		{`false == null`, false},
		{`false != null`, true},

		{`[1, 2] == null`, false},
		{`[1, 2] != null`, true},

		{`null == null`, true},
		{`null != null`, false},
	}
	testExpressionEvaluation(t, testCases, []errorTestCase{})
}

func TestEvalStringExpression(t *testing.T) {
	testCases := []happyTestCase{
		{`"hello"`, "hello"},
		{`"hello" + "hi"`, "hellohi"},
	}

	testExpressionEvaluation(t, testCases, []errorTestCase{})
}

func TestEvalArrayExpression(t *testing.T) {
	happyCases := []happyTestCase{
		{`["hello", 1, 2] + [3]`, []any{"hello", 1, 2, 3}},
		{`["hello", 1, 2] + []`, []any{"hello", 1, 2}},
	}

	errorCases := []errorTestCase{
		{`["hello", 1, 2] + 3`, "unsupported operation: ARRAY `+` INTEGER"},
		{`["hello", 1, 2] + "hello"`, "unsupported operation: ARRAY `+` STRING"},
	}
	testExpressionEvaluation(t, happyCases, errorCases)
}

func TestEvalHashExpression(t *testing.T) {
	happyCases := []happyTestCase{
		{`{"hello": 1} + {"hi": true}`, map[any]any{"hello": 1, "hi": true}},
		{`{"hello": 1} + {"hello": 3}`, map[any]any{"hello": 3}},
	}

	errorCases := []errorTestCase{
		{`{"hello": 1} + 3`, "unsupported operation: HASH `+` INTEGER"},
	}
	testExpressionEvaluation(t, happyCases, errorCases)
}
