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
