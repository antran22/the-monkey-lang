package object

import "fmt"

const ERROR_OBJ = "ERROR"

type Error struct {
	Message string
}

var _ Object = (*Error)(nil)

func (e *Error) Type() ObjectType { return ERROR_OBJ }

func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}

func (e *Error) IsTruthy() bool {
	return true
}

func NewError(message string) *Error {
	return &Error{
		Message: message,
	}
}

func NewErrorf(message string, a ...any) *Error {
	return NewError(fmt.Sprintf(message, a...))
}

func IsError(obj Object) bool {
	return obj != nil && obj.Type() == ERROR_OBJ
}
