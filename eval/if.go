package eval

import (
	"monkey/ast"
	"monkey/eval/object"
)

func evalIfExpression(ie *ast.IfExpression) object.Object {
	condition := Eval(ie.Condition)
	if object.IsError(condition) {
		return condition
	}

	if condition.IsTruthy() {
		return Eval(ie.ThenBranch)
	}
	if ie.ElseBranch != nil {
		return Eval(ie.ElseBranch)
	}
	return object.NULL
}
