package ast

import "monkey/token"

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (e *ExpressionStatement) TokenLiteral() string {
	return e.Token.Literal
}

func (e *ExpressionStatement) statementNode() {
	panic("unimplemented")
}

// interface check
var _ Statement = (*ExpressionStatement)(nil)
