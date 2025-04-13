package eval

import (
	"monkey/ast"
	"monkey/eval/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	// Statements
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.BlockStatement:
		return evalStatements(node.Statements)

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
		return evalPrefixExpression(node.Operator, rightVal)
	case *ast.InfixExpression:
		left, right := Eval(node.Left), Eval(node.Right)
		return evalInfixExpression(left, node.Operator, right)

	}
	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, stmt := range stmts {
		result = Eval(stmt)
	}
	return result
}
