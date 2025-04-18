package eval

import (
	"monkey/ast"
	"monkey/eval/object"
)

func evalFunctionExpression(node *ast.FunctionExpression, env *object.Environment) object.Object {
	result := &object.Function{
		FuncNode: node,
		Env:      env,
	}
	if len(node.Name) > 0 {
		env.Set(node.Name, result)
	}
	return result
}

func evalCallable(callable object.Object, args ...object.Object) object.Object {
	argCount := len(args)
	if argCount == 1 && object.IsError(args[0]) {
		return args[0]
	}

	var result object.Object

	switch callable := callable.(type) {
	case *object.Function:
		if argCount < len(callable.FuncNode.Parameters) {
			return object.FuncNotEnoughArg(callable.DisplayName(), len(callable.FuncNode.Parameters), argCount)
		}
		evaluationEnv := extendFunctionEnv(callable, args)
		result = Eval(callable.FuncNode.Body, evaluationEnv)
	case *object.Builtin:
		result = callable.Fn(args...)
	default:
		return object.NotACallable(callable)
	}

	if returnValue, ok := result.(*object.ReturnValue); ok {
		return returnValue.Value
	}
	return result
}

func evalFunctionCall(node *ast.CallExpression, env *object.Environment) object.Object {
	callable := Eval(node.Function, env)
	if object.IsError(callable) {
		return callable
	}
	args := evalExpressions(node.Arguments, env)

	return evalCallable(callable, args...)
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
