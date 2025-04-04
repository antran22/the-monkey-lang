package parser_test

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/parser"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLetStatements(t *testing.T) {
	r := require.New(t)
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`

	l := lexer.NewLexer(input)
	p := parser.NewParser(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

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

	l := lexer.NewLexer(input)
	p := parser.NewParser(l)

	p.ParseProgram()

	errors := p.Errors()

	r.Len(errors, 3)
}
