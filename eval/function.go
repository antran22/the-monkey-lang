package eval

import (
	"monkey/ast"
	"monkey/eval/object"
)

var builtins = map[string]*object.Builtin{
	"len": object.NewBuiltin("len", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return object.NewErrorf("wrong number of argument for `len`, expected 1, got %d", len(args))
		}
		switch arg := args[0].(type) {
		case *object.String:
			return &object.Integer{Value: len(arg.Value)}
		default:
			return object.NewErrorf("unsupported argument type for `len`: %s", arg.Type())
		}
	}),
}

func evalFunctionExpression(node *ast.FunctionExpression, env *object.Environment) object.Object {
	return &object.Function{
		FuncNode: node,
		Env:      env,
	}
}

func evalFunctionCall(node *ast.CallExpression, env *object.Environment) object.Object {
	callable := Eval(node.Function, env)
	if object.IsError(callable) {
		return callable
	}
	args := evalExpressions(node.Arguments, env)
	if len(args) == 1 && object.IsError(args[0]) {
		return args[0]
	}

	var result object.Object
	switch callable := callable.(type) {
	case *object.Function:
		evaluationEnv := extendFunctionEnv(callable, args)
		result = Eval(callable.FuncNode.Body, evaluationEnv)
	case *object.Builtin:
		result = callable.Fn(args...)
	default:
		result = object.NewErrorf("%s is not a callable, type: %T", node.Function.String(), callable)
	}

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
