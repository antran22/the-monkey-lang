package eval

import (
	"monkey/ast"
	"monkey/eval/object"
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	// Statements
	case *ast.Program:
		return evalProgram(node.Statements, env)
	case *ast.BlockStatement:
		return evalBlockStatements(node.Statements, env)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.ReturnStatement:
		val := Eval(node.Value, env)
		if object.IsError(val) {
			return val
		}
		return &object.ReturnValue{Value: val}
	case *ast.LetStatement:
		return evalLetStatement(node, env)

	// Literals
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.BooleanLiteral:
		return object.NewBoolean(node.Value)
	case *ast.NullLiteral:
		return object.NULL

	// Expressions

	case *ast.PrefixExpression:
		rightVal := Eval(node.Right, env)
		if object.IsError(rightVal) {
			return rightVal
		}
		return evalPrefixExpression(node.Operator, rightVal)
	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		if object.IsError(left) {
			return left
		}

		right := Eval(node.Right, env)
		if object.IsError(right) {
			return right
		}

		return evalInfixExpression(left, node.Operator, right)

	case *ast.FunctionExpression:
		return evalFunctionExpression(node, env)

	case *ast.CallExpression:
		return evalFunctionCall(node, env)

	// Control flows
	case *ast.IfExpression:
		return evalIfExpression(node, env)

	default:
		return object.NewErrorf("unable to evaluate expression (%T) %v", node, node)
	}
}

func evalExpressions(exps []ast.Expression, env *object.Environment) (result []object.Object) {
	for _, e := range exps {
		evaluated := Eval(e, env)
		if object.IsError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}
	return result
}
