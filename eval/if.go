package eval

import (
	"monkey/ast"
	"monkey/eval/object"
)

func evalIfExpression(ie *ast.IfExpression, env *object.Environment) object.Object {
	condition := Eval(ie.Condition, env)
	if object.IsError(condition) {
		return condition
	}

	if condition.IsTruthy() {
		return Eval(ie.ThenBranch, env)
	}
	if ie.ElseBranch != nil {
		return Eval(ie.ElseBranch, env)
	}
	return object.NULL
}
