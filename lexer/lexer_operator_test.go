package lexer_test

import (
	"monkey/token"
	"testing"
)

func TestLexerOperator(t *testing.T) {
	input := "=+(){}[],;&^|:"
	tokens := []tokenTestExp{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.LBRACKET, "["},
		{token.RBRACKET, "]"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.B_AND, "&"},
		{token.XOR, "^"},
		{token.B_OR, "|"},
		{token.COLON, ":"},
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
2 .. 2;
true && true;
false || true;
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

		{token.INT, "2"},
		{token.D_DOT, ".."},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},

		{token.TRUE, "true"},
		{token.AND, "&&"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},

		{token.FALSE, "false"},
		{token.OR, "||"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
	}

	testLexerTokenization(t, input, tokens)
}
