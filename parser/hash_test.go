package parser_test

import (
	"monkey/ast"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseEmptyHashLiteral(t *testing.T) {
	r := require.New(t)

	input := `{}`

	program := makeProgram(input, t)

	stmt := testExpressionStatement(t, program.Statements[0])

	hash, ok := stmt.Expression.(*ast.HashLiteral)
	r.True(ok)

	r.Len(hash.Pairs, 0)
}

func TestParseHashLiteral(t *testing.T) {
	r := require.New(t)

	inputs := []string{
		`{"one": 1, "two": 2, "three": 3}`,
		`{"one": 1, "two": 2, "three": 3,}`, // trailing comma
	}

	expected := []expectedKVPair{
		{[]byte("one"), 1},
		{[]byte("two"), 2},
		{[]byte("three"), 3},
	}

	for _, input := range inputs {
		program := makeProgram(input, t)

		stmt := testExpressionStatement(t, program.Statements[0])

		hash, ok := stmt.Expression.(*ast.HashLiteral)
		r.True(ok)

		r.Len(hash.Pairs, 3)

		testLiteralExpression(t, hash, expected)
	}
}

func TestParseHashLiteralWithExpression(t *testing.T) {
	r := require.New(t)

	input := `{
    "one": 0 + 1,
    "two": 4 / 2,
    "three": 1 * 3
  }`

	program := makeProgram(input, t)

	stmt := testExpressionStatement(t, program.Statements[0])

	hash, ok := stmt.Expression.(*ast.HashLiteral)
	r.True(ok)

	r.Len(hash.Pairs, 3)

	tests := []struct {
		key    []byte
		tester func(ast.Expression)
	}{
		{
			[]byte("one"), func(e ast.Expression) {
				testInfixExpression(t, e, 0, "+", 1)
			},
		},
		{
			[]byte("two"), func(e ast.Expression) {
				testInfixExpression(t, e, 4, "/", 2)
			},
		},
		{
			[]byte("three"), func(e ast.Expression) {
				testInfixExpression(t, e, 1, "*", 3)
			},
		},
	}

	for idx, pair := range hash.Pairs {
		testLiteralExpression(t, pair.Key, tests[idx].key)

		tester := tests[idx].tester
		tester(pair.Value)
	}
}
