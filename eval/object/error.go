package object

import (
	"fmt"
	"hash/fnv"
	"monkey/ast"
)

const ERROR_OBJ = "ERROR"

type Error struct {
	Message string
}

var (
	_ Object   = (*Error)(nil)
	_ Hashable = (*Error)(nil)
)

func (e *Error) Type() ObjectType { return ERROR_OBJ }

func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}

func (e *Error) IsTruthy() bool {
	return true
}

func (e *Error) Hash() HashKey {
	h := fnv.New64a()
	h.Write([]byte(e.Message))
	return HashKey{
		Type:  ERROR_OBJ,
		Value: h.Sum64(),
	}
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

// array & hash error format
func ArrayOutOfBoundError(idx int) *Error {
	return NewErrorf("index out of bound: %d", idx)
}

func TypeNotHashable(t ObjectType) *Error {
	return NewErrorf("value type not usable as hash key: %s", t)
}

// function error format
func FuncNotEnoughArg(name string, expectedCount, argCount int) *Error {
	return NewErrorf("too few argument for `%s`, expected at least %d, got %d", name, expectedCount, argCount)
}

func NotACallable(obj Object) *Error {
	return NewErrorf("not a callable: %s", obj.Inspect())
}

func FuncArgTypeMismatch(name string, argNum int, expectedType string, actualType string) *Error {
	return NewErrorf("wrong value type for argument #%d for `%s`, expected %s, got %s", argNum, name, expectedType, actualType)
}
