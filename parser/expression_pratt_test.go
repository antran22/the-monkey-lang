package parser_test

import (
	"fmt"
	"monkey/ast"
	"testing"

	"github.com/stretchr/testify/require"
)

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

			stmt := testExpression(t, program.Statements[0])

			exp, ok := stmt.Expression.(*ast.PrefixExpression)
			r.True(ok, "exp is not *ast.PrefixExpression. got=%T", stmt.Expression)
			r.Equal(ast.Operator(tt.operator), exp.Operator)

			testIntegerLiteral(pt, exp.Right, tt.integerValue)
		})
	}
}

func TestParsingInfixExpressions(t *testing.T) {
	infixTest := []struct {
		input      string
		leftValue  any
		operator   string
		rightValue any
	}{
		{"5 + 5", 5, "+", 5},
		{"5 - 5", 5, "-", 5},
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
		{"5 ^ 5;", 5, "^", 5},
		{"5 & 5;", 5, "&", 5},
		{"5 | 5;", 5, "|", 5},
		{"true == true", true, "==", true},
		{"true && true", true, "&&", true},
		{"true || true", true, "||", true},
		{"true != false", true, "!=", false},
		{"false == false", false, "==", false},
	}

	for i, tt := range infixTest {
		t.Run(fmt.Sprintf("prefix_%d", i), func(pt *testing.T) {
			r := require.New(pt)
			program := makeProgram(tt.input, t)

			r.Len(program.Statements, 1)

			stmt := testExpression(t, program.Statements[0])

			testInfixExpression(pt, stmt.Expression, tt.leftValue, ast.Operator(tt.operator), tt.rightValue)
		})
	}
}

func TestParsingPrecedence(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"true != true",
			"(true != true)",
		},
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"!-a",
			"(!(-a))",
		},
		{
			"a + b + c",
			"((a + b) + c)",
		},
		{
			"a + b - c",
			"((a + b) - c)",
		},
		{
			"a * b * c",
			"((a * b) * c)",
		},
		{
			"a * b / c",
			"((a * b) / c)",
		},
		{
			"a + b / c",
			"(a + (b / c))",
		},
		{
			"a + b * c + d / e - f",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"3 + 4; -5 * 5",
			"(3 + 4)((-5) * 5)",
		},
		{
			"5 > 4 == 3 < 4",
			"((5 > 4) == (3 < 4))",
		},
		{
			"5 >= 4 == 3 <= 4",
			"((5 >= 4) == (3 <= 4))",
		},
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			"true",
			"true",
		},
		{
			"false",
			"false",
		},
		{
			"3 > 5 == false",
			"((3 > 5) == false)",
		},
		{
			"3 < 5 == true",
			"((3 < 5) == true)",
		},
		{
			"1 + (2 + 3) + 4",
			"((1 + (2 + 3)) + 4)",
		},
		{
			"(5 + 5) * 2",
			"((5 + 5) * 2)",
		},
		{
			"2 / (5 + 5)",
			"(2 / (5 + 5))",
		},
		{
			"-(5 + 5)",
			"(-(5 + 5))",
		},
		{
			"!(true == true)",
			"(!(true == true))",
		},
		{
			"a + add(b * c) + d",
			"((a + add((b * c))) + d)",
		},
		{
			"add(a, b, 1, 2 * 3, 4 + 5, add(6, 7 * 8))",
			"add(a, b, 1, (2 * 3), (4 + 5), add(6, (7 * 8)))",
		},
		{
			"add(a + b + c * d / f + g)",
			"add((((a + b) + ((c * d) / f)) + g))",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("precedence test %d", i), func(it *testing.T) {
			r := require.New(it)

			pr := makeProgram(tt.input, it)

			actual := pr.String()

			r.Equal(tt.expected, actual)
		})
	}
}
