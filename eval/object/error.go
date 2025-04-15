package object

import (
	"fmt"
	"monkey/ast"
)

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

// common error format

func UnknownInfixOpError(left Object, operator string, right Object) *Error {
	return NewErrorf("unsupported operation: %s `%s` %s", left.Type(), operator, right.Type())
}

func UnknownPrefixOpError(operator string, right Object) *Error {
	return NewErrorf("unsupported operation: `%s` %s", operator, right.Type())
}

func InvalidExpressionError(node ast.Node) *Error {
	return NewErrorf("unable to evaluate expression: (%T) %v", node, node)
}

// array error format
func ArrayOutOfBoundError(idx int) *Error {
	return NewErrorf("index out of bound: %d", idx)
}

// function error format
func FuncArgCountMismatch(name string, expectedCount, argCount int) *Error {
	return NewErrorf("wrong number of argument for `%s`, expected %d, got %d", name, expectedCount, argCount)
}

func NotACallable(node ast.Node) *Error {
	return NewErrorf("%s is not a callable, type: %T", node.String(), node)
}

func FuncArgTypeMismatch(name string, argNum int, expectedType string, actualType string) *Error {
	return NewErrorf("wrong value type for argument #%d for `%s`, expected %s, got %s", argNum, name, expectedType, actualType)
}
