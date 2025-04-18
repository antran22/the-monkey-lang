package eval_test

import (
	"testing"
)

func TestBuiltinLen(t *testing.T) {
	happyCases := []happyTestCase{
		{`len("hello")`, 5},
		{`len("hello", "hi")`, 5},
		{`len("hello world")`, 11},
		{`len([])`, 0},
		{`len([1, 2, "hello"])`, 3},
	}

	errorCases := []errorTestCase{
		{`len()`, "too few argument for `len`, expected at least 1, got 0"},
		{`len(1)`, "wrong value type for argument #0 for `len`, expected STRING | ARRAY, got INTEGER"},
		{`len(false)`, "wrong value type for argument #0 for `len`, expected STRING | ARRAY, got BOOLEAN"},
	}

	testExpressionEvaluation(t, happyCases, errorCases)
}

func TestBuiltinFirst(t *testing.T) {
	happyCases := []happyTestCase{
		{`first([1])`, 1},
		{`first("h")`, "h"},
		{`first([1, 2])`, 1},
		{`first("hello")`, "h"},
		{`first("hello", "hi")`, "h"},
	}

	errorCases := []errorTestCase{
		{`first()`, "too few argument for `first`, expected at least 1, got 0"},
		{`first(1)`, "wrong value type for argument #0 for `first`, expected STRING | ARRAY, got INTEGER"},
		{`first(false)`, "wrong value type for argument #0 for `first`, expected STRING | ARRAY, got BOOLEAN"},
		{`first([])`, "index out of bound: 0"},
		{`first("")`, "index out of bound: 0"},
	}

	testExpressionEvaluation(t, happyCases, errorCases)
}

func TestBuiltinLast(t *testing.T) {
	happyCases := []happyTestCase{
		{`last([1])`, 1},
		{`last("h")`, "h"},
		{`last([1, 2])`, 2},
		{`last("hello")`, "o"},
		{`last("hello", "hi")`, "o"},
	}

	errorCases := []errorTestCase{
		{`last()`, "too few argument for `last`, expected at least 1, got 0"},
		{`last(1)`, "wrong value type for argument #0 for `last`, expected STRING | ARRAY, got INTEGER"},
		{`last(false)`, "wrong value type for argument #0 for `last`, expected STRING | ARRAY, got BOOLEAN"},
		{`last([])`, "index out of bound: -1"},
		{`last("")`, "index out of bound: -1"},
	}

	testExpressionEvaluation(t, happyCases, errorCases)
}

func TestBuiltinTail(t *testing.T) {
	happyCases := []happyTestCase{
		{`tail([1])`, []any{}},
		{`tail("h")`, ""},
		{`tail([1, 2])`, []any{2}},
		{`tail("hello")`, "ello"},
		{`tail("hello", "hi")`, "ello"},
	}

	errorCases := []errorTestCase{
		{`tail()`, "too few argument for `tail`, expected at least 1, got 0"},
		{`tail(1)`, "wrong value type for argument #0 for `tail`, expected STRING | ARRAY, got INTEGER"},
		{`tail(false)`, "wrong value type for argument #0 for `tail`, expected STRING | ARRAY, got BOOLEAN"},
		{`tail([])`, "index out of bound: 1"},
		{`tail("")`, "index out of bound: 1"},
	}

	testExpressionEvaluation(t, happyCases, errorCases)
}

func TestBuiltinAppend(t *testing.T) {
	happyCases := []happyTestCase{
		{`append([], 1)`, []any{1}},
		{`append([1], 2)`, []any{1, 2}},
		{`append([1, 2], 3, 4)`, []any{1, 2, 3, 4}},
		{`append([1, 2], true, "hello")`, []any{1, 2, true, "hello"}},
	}

	errorCases := []errorTestCase{
		{`append()`, "too few argument for `append`, expected at least 2, got 0"},
		{`append([])`, "too few argument for `append`, expected at least 2, got 1"},
		{`append(1, 1)`, "wrong value type for argument #0 for `append`, expected ARRAY, got INTEGER"},
		{`append(false, 2)`, "wrong value type for argument #0 for `append`, expected ARRAY, got BOOLEAN"},
		{`append("hello", 3)`, "wrong value type for argument #0 for `append`, expected ARRAY, got STRING"},
	}

	testExpressionEvaluation(t, happyCases, errorCases)
}

func TestBuiltinMap(t *testing.T) {
	happyCases := []happyTestCase{
		{`map([], fn(){return 1;})`, []any{}},
		{`map([true, false], fn () {return 1;})`, []any{1, 1}},
		{`map([1, 2], fn (a) {return a * 2;})`, []any{2, 4}},
		{`map([true, false], fn (a, idx) {return idx;})`, []any{0, 1}},
	}

	errorCases := []errorTestCase{}

	testExpressionEvaluation(t, happyCases, errorCases)
}
