package parser_test

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/parser"
	"testing"

	"github.com/stretchr/testify/require"
)

func makeProgram(input string, t *testing.T) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	return program
}

func checkParserErrors(t *testing.T, p *parser.Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func testExpression(s ast.Statement, r *require.Assertions) *ast.ExpressionStatement {
	stmt, ok := s.(*ast.ExpressionStatement)
	r.True(ok, "not ast.ExpressionStatement. got=%T", s)
	return stmt
}

func testLiteralExpression(t *testing.T, exp ast.Expression, expected any) {
	switch v := expected.(type) {
	case int:
		testIntegerLiteral(t, exp, v)
		return
	case int64:
		testIntegerLiteral(t, exp, int(v))
		return
	case string:
		testIdentifier(t, exp, v)
		return
	case bool:
		testBooleanLiteral(t, exp, v)
		return
	default:
		t.Fatalf("no test function for expected value %#v", expected)
	}
}

func testIdentifier(t *testing.T, exp ast.Expression, value string) {
	r := require.New(t)

	ident, ok := exp.(*ast.Identifier)

	r.True(ok, "exp is not *ast.Identifier")

	r.Equal(ident.Value, value)

	r.Equal(ident.TokenLiteral(), value)
}

func testIntegerLiteral(t *testing.T, exp ast.Expression, value int) {
	r := require.New(t)

	integ, ok := exp.(*ast.IntegerLiteral)
	r.Truef(ok, "exp not *ast.IntegerLiteral. got=%T", exp)

	r.Equal(value, integ.Value)

	r.Equal(fmt.Sprintf("%d", value), integ.TokenLiteral())
}

func testBooleanLiteral(t *testing.T, exp ast.Expression, value bool) {
	r := require.New(t)

	b, ok := exp.(*ast.BooleanLiteral)

	r.Truef(ok, "exp not *ast.BooleanLiteral. got=%T", exp)

	r.Equal(value, b.Value)

	r.Equal(fmt.Sprintf("%t", value), b.TokenLiteral())
}

func testInfixExpression(t *testing.T, exp ast.Expression, left interface{}, operator string, right interface{}) {
	r := require.New(t)
	opExp, ok := exp.(*ast.InfixExpression)
	r.True(ok, "exp is not an ast.InfixExpression")

	testLiteralExpression(t, opExp.Left, left)

	r.Equal(operator, opExp.Operator)

	testLiteralExpression(t, opExp.Right, right)
}
