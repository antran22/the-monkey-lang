package ast

import (
	"bytes"
	"monkey/token"
)

type Expression interface {
	Node
	expressionNode()
}

// Expression Statement

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (e *ExpressionStatement) String() string {
	return e.Expression.String()
}

func (e *ExpressionStatement) TokenLiteral() string {
	return e.Token.Literal
}

func (e *ExpressionStatement) statementNode() {}

var _ Statement = (*ExpressionStatement)(nil)

// Prefix Expression

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (p *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(p.Operator)
	out.WriteString(p.Right.String())
	out.WriteString(")")

	return out.String()
}

func (p *PrefixExpression) TokenLiteral() string {
	return p.Token.Literal
}

func (p *PrefixExpression) expressionNode() {}

var _ Expression = (*PrefixExpression)(nil)

// Infix Expression

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (i *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(i.Left.String())
	out.WriteString(" " + i.Operator + " ")
	out.WriteString(i.Right.String())
	out.WriteString(")")

	return out.String()
}

func (i *InfixExpression) TokenLiteral() string {
	return i.Token.Literal
}

func (i *InfixExpression) expressionNode() {}

var _ Expression = (*InfixExpression)(nil)
