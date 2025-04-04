package parser_test

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/parser"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReturnStatements(t *testing.T) {
	r := require.New(t)
	input := `
return 5;
return 10;
return 993322;
`
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()

	checkParserErrors(t, p)

	r.Equal(3, len(program.Statements))

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		r.True(ok, "stmt not *ast.ReturnStatement. got=%T", stmt)
		r.Equal("return", returnStmt.TokenLiteral())
	}
}
