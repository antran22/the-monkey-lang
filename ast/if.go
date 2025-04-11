package ast

import (
	"bytes"
	"monkey/token"
)

type IfExpression struct {
	Token      token.Token
	Condition  Expression
	TrueBranch *BlockStatement
	ElseBranch *BlockStatement
}

func (i *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(i.Condition.String())
	out.WriteString(" ")
	out.WriteString(i.TrueBranch.String())

	if i.ElseBranch != nil {
		out.WriteString(" else ")
		out.WriteString(i.ElseBranch.String())
	}
}

func (i *IfExpression) TokenLiteral() string {
	return i.Token.Literal
}

func (i *IfExpression) expressionNode() {}

var _ Expression = (*IfExpression)(nil)
