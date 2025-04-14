package eval

import (
	"monkey/ast"
	"monkey/eval/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	// Statements
	case *ast.Program:
		return evalProgram(node.Statements)
	case *ast.BlockStatement:
		return evalBlockStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.ReturnStatement:
		val := Eval(node.Value)
		if object.IsError(val) {
			return val
		}
		return &object.ReturnValue{Value: val}

		// Literals
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.BooleanLiteral:
		return object.NewBoolean(node.Value)
	case *ast.NullLiteral:
		return object.NULL

	// Expressions

	case *ast.PrefixExpression:
		rightVal := Eval(node.Right)
		if object.IsError(rightVal) {
			return rightVal
		}
		return evalPrefixExpression(node.Operator, rightVal)
	case *ast.InfixExpression:
		left := Eval(node.Left)
		if object.IsError(left) {
			return left
		}

		right := Eval(node.Right)
		if object.IsError(right) {
			return right
		}

		return evalInfixExpression(left, node.Operator, right)

	// Control flows
	case *ast.IfExpression:
		return evalIfExpression(node)

	default:
		return object.NewErrorf("unable to evaluate expression (%T) %v", node, node)
	}
}

func evalProgram(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, stmt := range stmts {
		result = Eval(stmt)
		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}
	return result
}

func evalBlockStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, stmt := range stmts {
		result = Eval(stmt)
		if result == nil {
			continue
		}

		t := result.Type()

		if t == object.RETURN_OBJ || t == object.ERROR_OBJ {
			return result
		}
	}
	return result
}
