package eval

import (
	"monkey/ast"
	"monkey/eval/object"
)

func evalFunctionExpression(node *ast.FunctionExpression, env *object.Environment) object.Object {
	return &object.Function{
		FuncNode: node,
		Env:      env,
	}
}

func evalFunctionCall(node *ast.CallExpression, env *object.Environment) object.Object {
	function := Eval(node.Function, env)
	if object.IsError(function) {
		return function
	}
	args := evalExpressions(node.Arguments, env)
	if len(args) == 1 && object.IsError(args[0]) {
		return args[0]
	}

	fn, ok := function.(*object.Function)
	if !ok {
		return object.NewErrorf("expression is not a callable: %s", function.Inspect())
	}

	evaluationEnv := extendFunctionEnv(fn, args)

	result := Eval(fn.FuncNode.Body, evaluationEnv)

	if returnValue, ok := result.(*object.ReturnValue); ok {
		return returnValue.Value
	}
	return result
}

func extendFunctionEnv(
	fn *object.Function,
	args []object.Object,
) *object.Environment {
	env := object.NewWrappedEnvironment(fn.Env)

	for paramIdx, param := range fn.FuncNode.Parameters {
		env.Set(param.Value, args[paramIdx])
	}

	return env
}
