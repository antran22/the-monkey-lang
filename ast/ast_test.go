package ast

import (
	"monkey/token"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	r := require.New(t)

	p := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	r.Equal("let myVar = anotherVar;", p.String())
}
