package object

import "fmt"

type ObjectType string

const (
	INTEGER_OBJ ObjectType = "INTEGER"
	BOOLEAN_OBJ ObjectType = "BOOLEAN"
	NULL_OBJ    ObjectType = "NULL"
)

type Object interface {
	Type() ObjectType
	Truthy() bool
	Inspect() string
}

// Integer

type Integer struct {
	Value int
}

var _ Object = (*Integer)(nil)

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

func (i *Integer) Truthy() bool {
	return i.Value != 0
}

// Boolean

type Boolean struct {
	Value bool
}

var _ Object = (*Boolean)(nil)

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

func (b *Boolean) Truthy() bool {
	return b.Value
}

var (
	TRUE  = &Boolean{Value: true}
	FALSE = &Boolean{Value: false}
)

func NewBoolean(value bool) *Boolean {
	if value {
		return TRUE
	}
	return FALSE
}

// Null

var NULL = &Null{}

var _ Object = (*Null)(nil)

type Null struct{}

func (n *Null) Truthy() bool {
	return false
}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }
