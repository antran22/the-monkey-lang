package parser_test

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/parser"
	"testing"

	"github.com/stretchr/testify/require"
)

// Let statement

func TestLetStatements(t *testing.T) {
	r := require.New(t)
	input := `
let x = 5;
let y = 10;
let foobar = y;
`

	tests := []struct {
		expectedIdentifier string
		value              any
	}{
		{"x", 5},
		{"y", 10},
		{"foobar", "y"},
	}

	program := makeProgram(input, t)

	r.NotNil(program)

	r.Len(program.Statements, 3)

	for i, tc := range tests {
		stmt := program.Statements[i]
		letStmt := testLetStatement(t, stmt, tc.expectedIdentifier)

		testLiteralExpression(t, letStmt.Value, tc.value)
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) *ast.LetStatement {
	r := require.New(t)

	r.Equal("let", s.TokenLiteral())
	letStmt, ok := s.(*ast.LetStatement)

	r.True(ok, "s is not *ast.LetStatement")

	r.Equal(name, letStmt.Name.Value)

	r.Equal(name, letStmt.Name.TokenLiteral())

	return letStmt
}

func TestFaultyLetStatement(t *testing.T) {
	r := require.New(t)
	input := `
let = 5;
let y 10;
let 838383;
`

	l := lexer.New(input)
	p := parser.New(l)

	p.ParseProgram()

	errors := p.Errors()

	r.Len(errors, 4)
}

// Return statement

func TestReturnStatements(t *testing.T) {
	r := require.New(t)
	input := `
return 5;
return 10;
return y;
`

	expectedReturns := []any{
		5,
		10,
		"y",
	}

	program := makeProgram(input, t)

	r.Len(program.Statements, 3)

	for i, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		r.True(ok, "stmt not *ast.ReturnStatement. got=%T", stmt)
		r.Equal("return", returnStmt.TokenLiteral())

		testLiteralExpression(t, returnStmt.Value, expectedReturns[i])
	}
}

// Block Statement

func TestBlockStatement(t *testing.T) {
	r := require.New(t)

	input := `{
let x = 1;
return 2;
}`

	program := makeProgram(input, t)
	r.Len(program.Statements, 1)

	bl, ok := program.Statements[0].(*ast.BlockStatement)

	r.Truef(ok, "not an *ast.BlockStatement")
	r.Len(bl.Statements, 2)
}
