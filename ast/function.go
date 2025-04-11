package ast

import (
	"bytes"
	"monkey/token"
	"strings"
)

type FunctionExpression struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

var _ Expression = (*FunctionExpression)(nil)

func (f *FunctionExpression) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(f.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(f.Body.String())

	return out.String()
}

func (f *FunctionExpression) TokenLiteral() string {
	return f.Token.Literal
}

func (f *FunctionExpression) expressionNode() {}

// function call expression

type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (c *CallExpression) String() string {
	var out bytes.Buffer

	out.WriteString(c.Function.String())
	out.WriteString("(")
	for i, a := range c.Arguments {
		if i > 0 {
			out.WriteString(", ")
		}
		out.WriteString(a.String())
	}
	out.WriteString(")")

	return out.String()
}

func (c *CallExpression) TokenLiteral() string {
	return c.Token.Literal
}

func (c *CallExpression) expressionNode() {}

var _ Expression = (*CallExpression)(nil)
