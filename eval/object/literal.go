package object

import (
	"fmt"
	"hash/fnv"
)

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

func NewString(v string) *String {
	return &String{
		Value: v,
	}
}

var (
	_ Object   = (*String)(nil)
	_ Hashable = (*String)(nil)
)

func (s *String) Inspect() string {
	return `"` + s.Value + `"`
}

func (s *String) Type() ObjectType {
	return STRING_OBJ
}

func (s *String) IsTruthy() bool {
	return len(s.Value) > 0
}

func (s *String) Hash() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{
		Type:  s.Type(),
		Value: h.Sum64(),
	}
}

func (a *String) Index(i int) Object {
	if i < 0 || i >= len(a.Value) {
		return ArrayOutOfBoundError(i)
	}
	return &String{Value: string(a.Value[i])}
}

func (a *String) Slice(start, end, step int) Object {
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

var (
	_ Object   = (*Integer)(nil)
	_ Hashable = (*Integer)(nil)
)

func NewInt(value int) *Integer {
	return &Integer{
		Value: value,
	}
}

func (i *Integer) Hash() HashKey {
	return HashKey{
		Type:  INTEGER_OBJ,
		Value: uint64(i.Value),
	}
}

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

var (
	_ Object   = (*Boolean)(nil)
	_ Hashable = (*Boolean)(nil)
)

func (i *Boolean) Hash() HashKey {
	var value uint64
	if i.Value {
		value = 0
	} else {
		value = 1
	}
	return HashKey{
		Type:  BOOLEAN_OBJ,
		Value: value,
	}
}

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

type Null struct{}

var (
	_ Object   = (*Null)(nil)
	_ Hashable = (*Null)(nil)
)

func (n *Null) Hash() HashKey {
	return HashKey{
		Type:  NULL_OBJ,
		Value: 0,
	}
}

func (n *Null) IsTruthy() bool {
	return false
}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }
