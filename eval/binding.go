package eval

import (
	"monkey/ast"
	"monkey/eval/object"
)

func evalLetStatement(stmt *ast.LetStatement, env *object.Environment) object.Object {
	val := Eval(stmt.Value, env)
	if object.IsError(val) {
		return val
	}
	env.Set(stmt.Name.Value, val)
	return val
}

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}
	return object.NewErrorf("identifier not found: %s", node.Value)
}
