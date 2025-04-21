package object

import (
	"fmt"
)

const RANGE_OBJ = "RANGE"

type Range struct {
	Start int
	End   int
	Step  int
}

var _ Object = (*Range)(nil)

func (r *Range) Inspect() string {
	return fmt.Sprintf("range(%d,%d,%d)", r.Start, r.End, r.Step)
}

func (r *Range) IsTruthy() bool {
	return true
}

func (r *Range) Type() ObjectType {
	return RANGE_OBJ
}

func NewImplicitRange(start, end int) *Range {
	var step int
	if start > end {
		step = -1
	} else {
		step = 1
	}
	return &Range{
		Start: start,
		End:   end,
		Step:  step,
	}
}

func NewExplicitRange(start, end, step int) (*Range, *Error) {
	if step == 0 {
		return nil, NewError("invalid range: `step` must not be zero")
	}

	if start < end && step < 0 {
		return nil, NewError("invalid range: `step` must be positive if start < end")
	}

	if start > end && step > 0 {
		return nil, NewError("invalid range: `step` must be negative if start > end")
	}

	return &Range{
		Start: start,
		End:   end,
		Step:  step,
	}, nil
}
