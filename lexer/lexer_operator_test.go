package lexer_test

import (
	"monkey/token"
	"testing"
)

func TestLexerOperator(t *testing.T) {
	input := "=+(){},;"
	tokens := []tokenTestExp{
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

	testLexerTokenization(t, input, tokens)
}

func TestDigraphOperator(t *testing.T) {
	input := `
2 == 2;
2 <= 2;
2 >= 2;
2 != 2;
`
	tokens := []tokenTestExp{
		{token.INT, "2"},
		{token.EQ, "=="},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},

		{token.INT, "2"},
		{token.LE, "<="},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},

		{token.INT, "2"},
		{token.GE, ">="},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},

		{token.INT, "2"},
		{token.NOT_EQ, "!="},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},
	}

	testLexerTokenization(t, input, tokens)
}
