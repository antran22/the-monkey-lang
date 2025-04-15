package parser_test

import (
	"monkey/ast"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIfExpression(t *testing.T) {
	r := require.New(t)

	input := `if (x < y) { x }`

	program := makeProgram(input, t)
	r.Len(program.Statements, 1)
	stmt := testExpressionStatement(t, program.Statements[0])

	exp, ok := stmt.Expression.(*ast.IfExpression)
	r.True(ok, "exp is not *ast.IfExpression")

	testInfixExpression(t, exp.Condition, "x", ast.OP_LT, "y")

	r.NotNil(exp.ThenBranch)
	r.Len(exp.ThenBranch.Statements, 1)
	stmtExp := testExpressionStatement(t, exp.ThenBranch.Statements[0])
	testIdentifier(t, stmtExp.Expression, "x")

	r.Nil(exp.ElseBranch)
}

func TestIfElseExpression(t *testing.T) {
	r := require.New(t)

	input := `if (x >= y) { x } else { y }`

	program := makeProgram(input, t)
	r.Len(program.Statements, 1)
	stmt := testExpressionStatement(t, program.Statements[0])

	exp, ok := stmt.Expression.(*ast.IfExpression)
	r.True(ok, "exp is not *ast.IfExpression")

	testInfixExpression(t, exp.Condition, "x", ast.OP_GE, "y")

	r.NotNil(exp.ThenBranch)
	r.Len(exp.ThenBranch.Statements, 1)
	thenStmtExp := testExpressionStatement(t, exp.ThenBranch.Statements[0])
	testIdentifier(t, thenStmtExp.Expression, "x")

	r.NotNil(exp.ElseBranch)
	r.Len(exp.ElseBranch.Statements, 1)
	elseStmtExp := testExpressionStatement(t, exp.ElseBranch.Statements[0])
	testIdentifier(t, elseStmtExp.Expression, "y")
}
