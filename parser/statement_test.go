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
let foobar = 838383;
`

	program := makeProgram(input, t)

	r.NotNil(program)

	r.Len(program.Statements, 3)

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		assertLetStatement(r, stmt, tt.expectedIdentifier)
	}
}

func assertLetStatement(r *require.Assertions, s ast.Statement, name string) {
	r.Equal("let", s.TokenLiteral())
	letStmt, ok := s.(*ast.LetStatement)

	r.True(ok, "s is not *ast.LetStatement")

	r.Equal(name, letStmt.Name.Value)

	r.Equal(name, letStmt.Name.TokenLiteral())
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
return 993322;
`

	program := makeProgram(input, t)

	r.Len(program.Statements, 3)

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		r.True(ok, "stmt not *ast.ReturnStatement. got=%T", stmt)
		r.Equal("return", returnStmt.TokenLiteral())
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
