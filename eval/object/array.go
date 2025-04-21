package object

import (
	"bytes"
)

const ARRAY_OBJ ObjectType = "ARRAY"

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

func (a *Array) Index(i int) Object {
	if i < 0 || i >= len(a.Elements) {
		return ArrayOutOfBoundError(i)
	}
	return a.Elements[i]
}

func (a *Array) Slice(start, end, step int) Object {
	n := len(a.Elements)

	if start < 0 || start > n {
		return ArrayOutOfBoundError(start)
	}

	if end < -1 || end > n {
		return ArrayOutOfBoundError(end)
	}

	var result []Object
	dir := 1
	if start > end {
		dir = -1
	}

	if start <= end && step == 1 {
		result = a.Elements[start:end]
	} else {
		for i := start; i*dir < end*dir; i += step {
			result = append(result, a.Elements[i])
		}
	}

	return &Array{
		Elements: result,
	}
}
