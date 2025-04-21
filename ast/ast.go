package ast

import "bytes"

type Node interface {
	TokenLiteral() string
	String() string
}

type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

var _ Node = (*Program)(nil)

// operator

type Operator string

const (
	// arithmetic
	OP_PLUS        Operator = "+"
	OP_MINUS       Operator = "-"
	OP_MULTIPLY    Operator = "*"
	OP_DIVIDE      Operator = "/"
	OP_BITWISE_AND Operator = "&"
	OP_BITWISE_OR  Operator = "|"
	OP_BITWISE_XOR Operator = "^"

	// comparative
	OP_LT  Operator = "<"
	OP_GT  Operator = ">"
	OP_LE  Operator = "<="
	OP_GE  Operator = ">="
	OP_EQ  Operator = "=="
	OP_NEQ Operator = "!="

	OP_RANGE Operator = ".."

	// logic
	OP_NEGATE Operator = "!"
	OP_AND    Operator = "&&"
	OP_OR     Operator = "||"
)
