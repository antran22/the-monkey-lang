package eval

import (
	"fmt"
	"monkey/eval/object"
)

var builtins = map[string]*object.Builtin{}

func addBuiltIn(name string, minArgCount int, fn object.BuiltinFunc) {
	if minArgCount > 0 {
		builtins[name] = object.NewBuiltin(name, func(args ...object.Object) object.Object {
			if len(args) < minArgCount {
				return object.FuncNotEnoughArg(name, minArgCount, len(args))
			}
			return fn(args...)
		})
	} else {
		builtins[name] = object.NewBuiltin(name, fn)
	}
}

func init() {
	addBuiltIn("len", 1, func(args ...object.Object) object.Object {
		switch arg := args[0].(type) {
		case *object.String:
			return &object.Integer{Value: len(arg.Value)}
		case *object.Array:
			return &object.Integer{Value: len(arg.Elements)}
		default:
			return object.FuncArgTypeMismatch("len", 0, "STRING | ARRAY", string(arg.Type()))
		}
	})

	addBuiltIn("first", 1, func(args ...object.Object) object.Object {
		switch arg := args[0].(type) {
		case *object.String:
			return arg.Index(0)
		case *object.Array:
			return arg.Index(0)
		default:
			return object.FuncArgTypeMismatch("first", 0, "STRING | ARRAY", string(arg.Type()))
		}
	})

	addBuiltIn("last", 1, func(args ...object.Object) object.Object {
		switch arg := args[0].(type) {
		case *object.String:
			return arg.Index(len(arg.Value) - 1)
		case *object.Array:
			return arg.Index(len(arg.Elements) - 1)
		default:
			return object.FuncArgTypeMismatch("last", 0, "STRING | ARRAY", string(arg.Type()))
		}
	})

	addBuiltIn("tail", 1, func(args ...object.Object) object.Object {
		switch arg := args[0].(type) {
		case *object.String:
			lv := len(arg.Value)
			return arg.Slice(1, lv, 1)
		case *object.Array:
			lv := len(arg.Elements)
			return arg.Slice(1, lv, 1)
		default:
			return object.FuncArgTypeMismatch("tail", 0, "STRING | ARRAY", string(arg.Type()))
		}
	})

	addBuiltIn("append", 2, func(args ...object.Object) object.Object {
		rest := args[1:]
		switch arg := args[0].(type) {
		case *object.Array:
			return &object.Array{
				Elements: append(arg.Elements, rest...),
			}
		default:
			return object.FuncArgTypeMismatch("append", 0, "ARRAY", string(arg.Type()))
		}
	})

	addBuiltIn("map", 2, func(args ...object.Object) object.Object {
		if args[0].Type() != object.ARRAY_OBJ {
			return object.FuncArgTypeMismatch("map", 0, string(object.ARRAY_OBJ), string(args[0].Type()))
		}
		arr := args[0].(*object.Array)

		mapper := args[1]
		if mapper.Type() != object.FUNCTION_OBJ && mapper.Type() != object.BUILTIN_OBJ {
			return object.FuncArgTypeMismatch("map", 1, "CALLABLE", string(mapper.Type()))
		}

		result := make([]object.Object, 0, len(arr.Elements))
		for idx, value := range arr.Elements {
			mapV := evalCallable(mapper, value, object.NewInt(idx))
			result = append(result, mapV)
		}

		return &object.Array{
			Elements: result,
		}
	})

	addBuiltIn("range", 2, func(args ...object.Object) object.Object {
		startArg, endArg := args[0], args[1]

		if startArg.Type() != object.INTEGER_OBJ {
			return object.FuncArgTypeMismatch("range", 0, string(object.INTEGER_OBJ), string(startArg.Type()))
		}
		if endArg.Type() != object.INTEGER_OBJ {
			return object.FuncArgTypeMismatch("range", 1, string(object.INTEGER_OBJ), string(endArg.Type()))
		}

		start, end := startArg.(*object.Integer).Value, endArg.(*object.Integer).Value

		var step int
		if len(args) == 2 {
			if start < end {
				step = 1
			} else {
				step = -1
			}
		} else {
			stepArg := args[2]
			if stepArg.Type() != object.INTEGER_OBJ {
				return object.FuncArgTypeMismatch("range", 2, string(object.INTEGER_OBJ), string(stepArg.Type()))
			}
			step = stepArg.(*object.Integer).Value
		}

		if r, err := object.NewExplicitRange(start, end, step); err != nil {
			return err
		} else {
			return r
		}
	})

	addBuiltIn("print", 0, func(args ...object.Object) object.Object {
		for _, arg := range args {
			fmt.Print(arg.Inspect())
		}
		fmt.Println()
		return object.NULL
	})
}
