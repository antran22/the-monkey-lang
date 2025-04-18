package ast

import "monkey/token"

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) String() string {
	return i.Value
}

func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

var _ Node = (*Identifier)(nil)
