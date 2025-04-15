package parser_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIdentifierExpression(t *testing.T) {
	r := require.New(t)

	input := "foobar;"
	program := makeProgram(input, t)

	r.Len(program.Statements, 1)
	stmt := testExpressionStatement(t, program.Statements[0])
	testIdentifier(t, stmt.Expression, "foobar")
}

func TestIntegerLiteralExpression(t *testing.T) {
	r := require.New(t)

	input := "5;"
	program := makeProgram(input, t)

	r.Len(program.Statements, 1)
	stmt := testExpressionStatement(t, program.Statements[0])
	testLiteralExpression(t, stmt.Expression, 5)
}

func TestStringLiteralExpression(t *testing.T) {
	r := require.New(t)

	input := "\"abcdef\";"
	program := makeProgram(input, t)

	r.Len(program.Statements, 1)
	stmt := testExpressionStatement(t, program.Statements[0])
	testLiteralExpression(t, stmt.Expression, []byte("abcdef"))
}

func TestBooleanLiteralExpression(t *testing.T) {
	testCases := []struct {
		input string
		value bool
	}{
		{"true;", true},
		{"false;", false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("boolean_test_%t", tc.value), func(st *testing.T) {
			program := makeProgram(tc.input, t)
			require.Len(t, program.Statements, 1)

			stmt := testExpressionStatement(t, program.Statements[0])

			testLiteralExpression(t, stmt.Expression, tc.value)
		})
	}
}

func TestNullLiteralExpression(t *testing.T) {
	r := require.New(t)

	input := "null;"
	program := makeProgram(input, t)

	r.Len(program.Statements, 1)
	stmt := testExpressionStatement(t, program.Statements[0])
	testLiteralExpression(t, stmt.Expression, nil)
}
