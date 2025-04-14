package lexer_test

import (
	"monkey/token"
	"testing"
)

func TestLexerString(t *testing.T) {
	input := `
let a = "hello world";
`
	tokens := []tokenTestExp{
		{token.LET, "let"},
		{token.IDENT, "a"},
		{token.ASSIGN, "="},
		{token.STRING, "hello world"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	testLexerTokenization(t, input, tokens)
}
