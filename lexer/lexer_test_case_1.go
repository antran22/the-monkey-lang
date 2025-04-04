package lexer

import "monkey/token"

var testCase1 = lexerTestCase{
	input: "=+(){},;",
	tokens: []tokenTestExp{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	},
}
