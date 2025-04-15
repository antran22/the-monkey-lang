package object

import "bytes"

const ARRAY_OBJ = "ARRAY"

type Array struct {
	Elements []Object
}

var _ Object = (*Array)(nil)

func (a *Array) Inspect() string {
	var out bytes.Buffer
	out.WriteString("[")

	for i, e := range a.Elements {
		if i > 0 {
			out.WriteString(", ")
		}
		out.WriteString(e.Inspect())
	}
	out.WriteString("]")

	return out.String()
}

func (a *Array) IsTruthy() bool {
	return len(a.Elements) > 0
}

func (a *Array) Type() ObjectType {
	return ARRAY_OBJ
}
