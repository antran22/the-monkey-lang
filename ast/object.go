package ast

import (
	"bytes"
	"monkey/token"
)

// Array Literal

type ArrayLiteral struct {
	Token    token.Token
	Elements []Expression
}

var _ Expression = (*ArrayLiteral)(nil)

func (a *ArrayLiteral) String() string {
	var out bytes.Buffer
	out.WriteString("[")
	for i, el := range a.Elements {
		if i > 0 {
			out.WriteString(", ")
		}
		out.WriteString(el.String())
	}
	out.WriteString("]")

	return out.String()
}

func (a *ArrayLiteral) TokenLiteral() string {
	return a.Token.Literal
}

// Indexing Expression

type IndexExpression struct {
	Token token.Token
	Left  Expression
	Index Expression
}

var _ Expression = (*IndexExpression)(nil)

func (i *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(i.Left.String())
	out.WriteString("[")
	out.WriteString(i.Index.String())
	out.WriteString("])")

	return out.String()
}

func (i *IndexExpression) TokenLiteral() string {
	return i.Token.Literal
}
