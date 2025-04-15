package parser_test

import (
	"monkey/ast"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParsingArrayLiterals(t *testing.T) {
	r := require.New(t)

	input := "[1, 2 * 2, 3 + 3]"

	program := makeProgram(input, t)

	stmt := testExpressionStatement(t, program.Statements[0])

	array, ok := stmt.Expression.(*ast.ArrayLiteral)

	r.Truef(ok, "stmt.Expression is not *ast.ArrayLiteral, got %T (%v)", stmt.Expression, stmt.Expression)

	r.Len(array.Elements, 3)

	testIntegerLiteral(t, array.Elements[0], 1)
	testInfixExpression(t, array.Elements[1], 2, "*", 2)
	testInfixExpression(t, array.Elements[2], 3, "+", 3)
}

func TestParsingIndexingOperator(t *testing.T) {
	r := require.New(t)
	input := "[1, 2, 3][1]"

	program := makeProgram(input, t)

	stmt := testExpressionStatement(t, program.Statements[0])

	indexExp, ok := stmt.Expression.(*ast.IndexExpression)
	r.Truef(ok, "stmt.Expression is not *ast.IndexExpression, got %T (%v)", stmt.Expression, stmt.Expression)

	_, ok = (indexExp.Left).(*ast.ArrayLiteral)
	r.Truef(ok, "left is not *ast.ArrayLiteral, got %T (%v)", indexExp.Left, indexExp.Left)

	testLiteralExpression(t, indexExp.Index, 1)
}
