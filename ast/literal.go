package ast

import "monkey/token"

// Integer Literal
type IntegerLiteral struct {
	Token token.Token
	Value int
}

func (i *IntegerLiteral) expressionNode()      {}
func (i *IntegerLiteral) TokenLiteral() string { return i.Token.Literal }
func (i *IntegerLiteral) String() string       { return i.Token.Literal }

var _ Node = (*IntegerLiteral)(nil)

// Boolean interal

type BooleanLiteral struct {
	Token token.Token
	Value bool
}

func (i *BooleanLiteral) expressionNode()      {}
func (i *BooleanLiteral) TokenLiteral() string { return i.Token.Literal }
func (i *BooleanLiteral) String() string       { return i.Token.Literal }

var _ Node = (*BooleanLiteral)(nil)
