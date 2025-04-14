package object

const RETURN_OBJ ObjectType = "RETURN"

type ReturnValue struct {
	Value Object
}

var _ Object = (*ReturnValue)(nil)

func (r *ReturnValue) Inspect() string {
	return r.Value.Inspect()
}

func (r *ReturnValue) IsTruthy() bool {
	return r.Value.IsTruthy()
}

func (r *ReturnValue) Type() ObjectType {
	return RETURN_OBJ
}
