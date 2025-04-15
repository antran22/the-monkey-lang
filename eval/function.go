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
		case *object.Array:
			return &object.Integer{Value: len(arg.Elements)}
		default:
			return object.FuncArgTypeMismatch("len", 0, "STRING | ARRAY", string(arg.Type()))
		}
	}),

	"first": object.NewBuiltin("first", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return object.NewErrorf("wrong number of argument for `first`, expected 1, got %d", len(args))
		}
		switch arg := args[0].(type) {
		case *object.String:
			v := arg.Value
			if len(v) == 0 {
				return object.ArrayOutOfBoundError(0)
			}
			return &object.String{
				Value: string(v[0]),
			}
		case *object.Array:
			v := arg.Elements
			if len(v) == 0 {
				return object.ArrayOutOfBoundError(0)
			}
			return v[0]
		default:
			return object.FuncArgTypeMismatch("last", 0, "STRING | ARRAY", string(arg.Type()))
		}
	}),

	"last": object.NewBuiltin("last", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return object.NewErrorf("wrong number of argument for `last`, expected 1, got %d", len(args))
		}
		switch arg := args[0].(type) {
		case *object.String:
			v := arg.Value
			lv := len(v)
			if lv == 0 {
				return object.ArrayOutOfBoundError(0)
			}
			return &object.String{
				Value: string(v[lv-1]),
			}
		case *object.Array:
			v := arg.Elements
			lv := len(v)
			if lv == 0 {
				return object.ArrayOutOfBoundError(0)
			}
			return v[lv-1]
		default:
			return object.FuncArgTypeMismatch("last", 0, "STRING | ARRAY", string(arg.Type()))
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
	argCount := len(args)
	if argCount == 1 && object.IsError(args[0]) {
		return args[0]
	}

	var result object.Object
	switch callable := callable.(type) {
	case *object.Function:
		if argCount != len(callable.FuncNode.Parameters) {
			return object.FuncArgCountMismatch(callable.DisplayName(), len(callable.FuncNode.Parameters), argCount)
		}
		evaluationEnv := extendFunctionEnv(callable, args)
		result = Eval(callable.FuncNode.Body, evaluationEnv)
	case *object.Builtin:
		result = callable.Fn(args...)
	default:
		return object.NotACallable(node.Function)
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
