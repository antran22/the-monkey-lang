package object

import (
	"monkey/ast"
)

// Callable

type Callable interface {
	Object
	DisplayName() string
}

// Function
const FUNCTION_OBJ = "FUNCTION"

type Function struct {
	FuncNode *ast.FunctionExpression
	Env      *Environment
}

var _ Callable = (*Function)(nil)

func (f *Function) Inspect() string {
	return f.FuncNode.String()
}

func (f *Function) IsTruthy() bool {
	return true
}

func (f *Function) Type() ObjectType {
	return FUNCTION_OBJ
}

func (f *Function) DisplayName() string {
	return f.FuncNode.DisplayName()
}

// Builtin

const BUILTIN_OBJ = "BUILTIN"

type (
	BuiltinFunc func(args ...Object) Object
	Builtin     struct {
		Fn   BuiltinFunc
		Name string
	}
)

var _ Callable = (*Builtin)(nil)

func NewBuiltin(name string, fn BuiltinFunc) *Builtin {
	return &Builtin{
		Fn:   fn,
		Name: name,
	}
}

func (b *Builtin) Inspect() string {
	return b.Name
}

func (b *Builtin) IsTruthy() bool {
	return true
}

func (b *Builtin) Type() ObjectType {
	return BUILTIN_OBJ
}

func (b *Builtin) DisplayName() string {
	return b.Name
}
