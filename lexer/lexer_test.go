package lexer

import (
	"monkey/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextToken(t *testing.T) {
	assert := assert.New(t)
	input := "=+(){},;"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := NewLexer(input)

	for i, tt := range tests {
		tok := l.NextToken()

		assert.Equal(tt.expectedType, tok.Type, "tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		assert.Equal(tt.expectedLiteral, tok.Literal, "tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
	}
}
