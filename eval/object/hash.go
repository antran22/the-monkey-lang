package object

import (
	"bytes"
)

const HASH_OBJ ObjectType = "HASH"

type HashPair struct {
	Key   Object
	Value Object
}

type Hash struct {
	Pairs map[HashKey]HashPair
}

var _ Object = (*Hash)(nil)

func (h *Hash) Inspect() string {
	var out bytes.Buffer

	out.WriteString("{")
	isFirst := false
	for _, pair := range h.Pairs {
		if !isFirst {
			isFirst = true
		} else {
			out.WriteString(", ")
		}
		out.WriteString(pair.Key.Inspect())
		out.WriteString(":")
		out.WriteString(pair.Value.Inspect())
	}
	out.WriteString("}")
	return out.String()
}

func (h *Hash) IsTruthy() bool {
	return len(h.Pairs) > 0
}

func (h *Hash) Type() ObjectType {
	return HASH_OBJ
}
