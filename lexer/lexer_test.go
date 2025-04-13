package lexer_test

import (
	"monkey/lexer"
	"monkey/token"
	"testing"

	"github.com/stretchr/testify/require"
)

type tokenTestExp struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func testLexerTokenization(t *testing.T, input string, tokens []tokenTestExp) {
	assert := require.New(t)
	l := lexer.New(input)
	for i, tokenTest := range tokens {
		tok := l.NextToken()
		assert.Equal(tokenTest.expectedType, tok.Type, "tests[%d] - tokentype wrong. expected=%q, got=%q, literal=%q", i, tokenTest.expectedType, tok.Type, tok.Literal)
		assert.Equal(tokenTest.expectedLiteral, tok.Literal, "tests[%d] - literal wrong. expected=%q, got=%q", i, tokenTest.expectedLiteral, tok.Literal)
	}
}
