package ast

import "monkey/token"

// Integer Literal
type IntegerLiteral struct {
	Token token.Token
	Value int
}

var _ Expression = (*IntegerLiteral)(nil)

func (i *IntegerLiteral) expressionNode()      {}
func (i *IntegerLiteral) TokenLiteral() string { return i.Token.Literal }
func (i *IntegerLiteral) String() string       { return i.Token.Literal }

// Boolean interal

type BooleanLiteral struct {
	Token token.Token
	Value bool
}

var _ Expression = (*BooleanLiteral)(nil)

func (i *BooleanLiteral) expressionNode()      {}
func (i *BooleanLiteral) TokenLiteral() string { return i.Token.Literal }
func (i *BooleanLiteral) String() string       { return i.Token.Literal }

// Null literal

type NullLiteral struct {
	Token token.Token
}

var _ Expression = (*NullLiteral)(nil)

func (n *NullLiteral) expressionNode() {}
func (n *NullLiteral) String() string {
	return "null"
}

func (n *NullLiteral) TokenLiteral() string { return n.Token.Literal }
