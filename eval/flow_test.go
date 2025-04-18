package eval_test

import (
	"testing"
)

func TestReturnStatement(t *testing.T) {
	happyCases := []happyTestCase{
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

	errorCases := []errorTestCase{
		{"5 + true;", "unsupported operation: INTEGER `+` BOOLEAN"},
		{"5 + true; 5", "unsupported operation: INTEGER `+` BOOLEAN"},
		{`"hello" - "hi"`, "unsupported operation: STRING `-` STRING"},
		{"-true", "unsupported operation: `-` BOOLEAN"},
		{"true + false", "unsupported operation: BOOLEAN `+` BOOLEAN"},
		{"5; true + false; 5", "unsupported operation: BOOLEAN `+` BOOLEAN"},
		{"if (10 > 1) { true + false; }", "unsupported operation: BOOLEAN `+` BOOLEAN"},
	}

	testExpressionEvaluation(t, happyCases, errorCases)
}
