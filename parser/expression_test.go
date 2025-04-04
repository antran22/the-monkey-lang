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

func assertIsExpression(s ast.Statement, r *require.Assertions) *ast.ExpressionStatement {
	stmt, ok := s.(*ast.ExpressionStatement)
	r.True(ok, "not ast.ExpressionStatement. got=%T", s)
	return stmt
}

func TestIdentifierExpression(t *testing.T) {
	r := require.New(t)

	input := "foobar;"
	program := makeProgram(input, t)

	r.Len(program.Statements, 1)

	stmt := assertIsExpression(program.Statements[0], r)

	ident, ok := stmt.Expression.(*ast.Identifier)
	r.True(ok, "expr is not *ast.Identifier. got=%T", ident)

	r.Equal("foobar", ident.Value)
	r.Equal("foobar", ident.TokenLiteral())
}

func TestIntegerLiteralExpression(t *testing.T) {
	r := require.New(t)

	input := "5;"
	program := makeProgram(input, t)

	r.Len(program.Statements, 1)

	stmt := assertIsExpression(program.Statements[0], r)

	ident, ok := stmt.Expression.(*ast.IntegerLiteral)
	r.True(ok, "expr is not *ast.IntegerLiteral. got=%T", ident)
}

func assertIntegerLiteral(il ast.Expression, value int, r *require.Assertions) {
	integ, ok := il.(*ast.IntegerLiteral)
	r.True(ok, "il not *ast.IntegerLiteral. got=%T", il)

	r.Equal(value, integ.Value)

	r.Equal(fmt.Sprintf("%d", value), integ.TokenLiteral())
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
	}

	for i, tt := range prefixTests {
		t.Run(fmt.Sprintf("prefix_%d", i), func(pt *testing.T) {
			r := require.New(pt)
			program := makeProgram(tt.input, t)

			r.Len(program.Statements, 1)

			stmt := assertIsExpression(program.Statements[0], r)

			exp, ok := stmt.Expression.(*ast.PrefixExpression)
			r.True(ok, "exp is not *ast.PrefixExpression. got=%T", stmt.Expression)
			r.Equal(tt.operator, exp.Operator)

			assertIntegerLiteral(exp.Right, tt.integerValue, r)
		})
	}
}
