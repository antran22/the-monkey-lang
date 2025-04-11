package ast

import (
	"bytes"
	"monkey/token"
)

type Statement interface {
	Node
	statementNode()
}

// LetStatement

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

var _ Statement = (*LetStatement)(nil) // interface check

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral())
	out.WriteString(" ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// ReturnStatement

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

var _ Statement = (*ReturnStatement)(nil) // interface check

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral())
	out.WriteString(" ")

	if rs.Value != nil {
		out.WriteString(rs.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// BlockStatement

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

var _ Statement = (*BlockStatement)(nil) // interface check

func (b *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range b.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (b *BlockStatement) TokenLiteral() string {
	return b.Token.Literal
}

func (b *BlockStatement) statementNode() {}
