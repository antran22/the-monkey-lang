package lexer

import "monkey/token"

type tokenTestExp struct {
	expectedType    token.TokenType
	expectedLiteral string
}

type lexerTestCase struct {
	input  string
	tokens []tokenTestExp
}
