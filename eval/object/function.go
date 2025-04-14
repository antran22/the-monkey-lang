package object

import (
	"monkey/ast"
)

const FUNCTION_OBJ = "FUNCTION"

type Function struct {
	FuncNode *ast.FunctionExpression
	Env      *Environment
}

var _ Object = (*Function)(nil)

func (f *Function) Inspect() string {
	return f.FuncNode.String()
}

func (f *Function) IsTruthy() bool {
	return true
}

func (f *Function) Type() ObjectType {
	return FUNCTION_OBJ
}

const BUILTIN_OBJ = "BUILTIN"

type (
	BuiltinFunc func(args ...Object) Object
	Builtin     struct {
		Fn   BuiltinFunc
		Name string
	}
)

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

var _ Object = (*Builtin)(nil)
