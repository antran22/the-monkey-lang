package ast

import "monkey/token"

// Integer Literal
type IntegerLiteral struct {
	Token token.Token
	Value int
}

var _ Expression = (*IntegerLiteral)(nil)

func (i *IntegerLiteral) TokenLiteral() string { return i.Token.Literal }
func (i *IntegerLiteral) String() string       { return i.Token.Literal }

// Boolean literal

type BooleanLiteral struct {
	Token token.Token
	Value bool
}

var _ Expression = (*BooleanLiteral)(nil)

func (i *BooleanLiteral) TokenLiteral() string { return i.Token.Literal }
func (i *BooleanLiteral) String() string       { return i.Token.Literal }

// String literal

type StringLiteral struct {
	Token token.Token
	Value string
}

var _ Expression = (*StringLiteral)(nil)

func (i *StringLiteral) TokenLiteral() string { return i.Token.Literal }
func (i *StringLiteral) String() string       { return `"` + i.Token.Literal + `"` }

// Null literal

type NullLiteral struct {
	Token token.Token
}

var _ Expression = (*NullLiteral)(nil)

func (n *NullLiteral) String() string {
	return "null"
}

func (n *NullLiteral) TokenLiteral() string { return n.Token.Literal }
