package lexer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var testCases = []lexerTestCase{
	testCase1,
	testCase2,
	testCase3,
}

func TestNextToken(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test_case_%d", i+1), func(tt *testing.T) {
			assert := require.New(tt)
			l := NewLexer(testCase.input)
			for i, tokenTest := range testCase.tokens {
				tok := l.NextToken()
				assert.Equal(tokenTest.expectedType, tok.Type, "tests[%d] - tokentype wrong. expected=%q, got=%q, literal=%q", i, tokenTest.expectedType, tok.Type, tok.Literal)
				assert.Equal(tokenTest.expectedLiteral, tok.Literal, "tests[%d] - literal wrong. expected=%q, got=%q", i, tokenTest.expectedLiteral, tok.Literal)
			}
		})
	}
}
