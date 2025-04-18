package object

import "fmt"

const (
	INTEGER_OBJ ObjectType = "INTEGER"
	BOOLEAN_OBJ ObjectType = "BOOLEAN"
	NULL_OBJ    ObjectType = "NULL"
	STRING_OBJ  ObjectType = "STRING"
)

// String

type String struct {
	Value string
}

var _ Object = (*String)(nil)

func (s *String) Inspect() string {
	return `"` + s.Value + `"`
}

func (s *String) Type() ObjectType {
	return STRING_OBJ
}

func (s *String) IsTruthy() bool {
	return len(s.Value) > 0
}

func (a *String) Index(i int) Object {
	if i < 0 || i >= len(a.Value) {
		return ArrayOutOfBoundError(i)
	}
	return &String{Value: string(a.Value[i])}
}

func (a *String) Slice(start, end int) Object {
	n := len(a.Value)

	if start < 0 || start > n {
		return ArrayOutOfBoundError(start)
	}

	if end < 0 || end > n {
		return ArrayOutOfBoundError(start)
	}

	if start > end+1 {
		return NewErrorf("cannot take slice from %d to %d", start, end)
	}

	return &String{
		Value: a.Value[start:end],
	}
}

// Integer

type Integer struct {
	Value int
}

func NewInt(value int) *Integer {
	return &Integer{
		Value: value,
	}
}

var _ Object = (*Integer)(nil)

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

func (i *Integer) IsTruthy() bool {
	return i.Value != 0
}

// Boolean

type Boolean struct {
	Value bool
}

var (
	TRUE  = &Boolean{Value: true}
	FALSE = &Boolean{Value: false}
)

var _ Object = (*Boolean)(nil)

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

func (b *Boolean) IsTruthy() bool {
	return b.Value
}

func NewBoolean(value bool) *Boolean {
	if value {
		return TRUE
	}
	return FALSE
}

func NewFromObject(obj Object) *Boolean {
	if obj == TRUE || obj == FALSE {
		return obj.(*Boolean)
	}

	return NewBoolean(obj.IsTruthy())
}

// Null

var NULL = &Null{}

var _ Object = (*Null)(nil)

type Null struct{}

func (n *Null) IsTruthy() bool {
	return false
}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }
