package parser_test

import (
	"monkey/ast"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFunctionExpression(t *testing.T) {
	r := require.New(t)

	input := `fn(a, b) { a + b; }`

	program := makeProgram(input, t)
	r.Len(program.Statements, 1)
	stmt := testExpression(t, program.Statements[0])

	exp, ok := stmt.Expression.(*ast.FunctionExpression)
	r.True(ok, "exp is not *ast.FunctionExpression")

	r.Equal("fn", exp.TokenLiteral())

	r.Len(exp.Parameters, 2)
	testLiteralExpression(t, exp.Parameters[0], "a")
	testLiteralExpression(t, exp.Parameters[1], "b")

	r.NotNil(exp.Body)
	r.Len(exp.Body.Statements, 1)

	expStmt, ok := exp.Body.Statements[0].(*ast.ExpressionStatement)
	r.True(ok, "inner exp is not *ast.ExpressionStatement")

	testInfixExpression(t, expStmt.Expression, "a", "+", "b")
}

func TestFunctionParameterParsing(t *testing.T) {
	testCases := []struct {
		input          string
		expectedParams []string
	}{
		{"fn() {};", []string{}},
		{"fn(x) {};", []string{"x"}},
		{"fn(x, y, z) {};", []string{"x", "y", "z"}},
	}

	for _, tc := range testCases {
		r := require.New(t)
		program := makeProgram(tc.input, t)

		r.Len(program.Statements, 1)
		stmt := testExpression(t, program.Statements[0])

		exp, ok := stmt.Expression.(*ast.FunctionExpression)
		r.True(ok, "exp is not *ast.FunctionExpression")

		r.Equal(len(tc.expectedParams), len(exp.Parameters))

		for i, ident := range tc.expectedParams {
			testLiteralExpression(t, exp.Parameters[i], ident)
		}
	}
}

// Function calling

func TestFunctionCallExpression(t *testing.T) {
	r := require.New(t)
	input := "add(1, 2 * 3, 4 + 5)"

	program := makeProgram(input, t)
	r.Len(program.Statements, 1)
	stmt := testExpression(t, program.Statements[0])

	exp, ok := stmt.Expression.(*ast.CallExpression)
	r.True(ok, "exp is not *ast.CallExpression")

	testIdentifier(t, exp.Function, "add")

	r.Len(exp.Arguments, 3)

	testLiteralExpression(t, exp.Arguments[0], 1)
	testInfixExpression(t, exp.Arguments[1], 2, "*", 3)
	testInfixExpression(t, exp.Arguments[2], 4, "+", 5)
}
