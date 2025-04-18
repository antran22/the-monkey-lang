package eval_test

import (
	"testing"
)

func TestLetStatements(t *testing.T) {
	happyCases := []happyTestCase{
		{"let a = 5; a;", 5},
		{"let a = 5 * 5; a;", 25},
		{"let a = 5; let b = a; b;", 5},
		{"let a = 5; let b = a; let c = a + b + 5; c;", 15},
	}

	errorCases := []errorTestCase{
		{"foo", "identifier not found: foo"},
		{"let foo = foo", "identifier not found: foo"},
	}

	testExpressionEvaluation(t, happyCases, errorCases)
}
