package eval

import (
	"monkey/ast"
	"monkey/eval/object"
)

// prefix expression

func evalPrefixExpression(operator ast.Operator, value object.Object) object.Object {
	switch operator {
	case ast.OP_NEGATE:
		return evalPrefixBangOperator(value)
	case ast.OP_MINUS:
		return evalPrefixMinusOperator(value)
	}

	return object.NULL
}

func evalPrefixBangOperator(value object.Object) object.Object {
	truthy := value.Truthy()
	return object.NewBoolean(!truthy)
}

func evalPrefixMinusOperator(value object.Object) object.Object {
	switch value.Type() {
	case object.INTEGER_OBJ:
		iv := value.(*object.Integer).Value
		return &object.Integer{Value: -iv}
	case object.BOOLEAN_OBJ:
		if value == object.TRUE {
			return &object.Integer{Value: -1}
		}
		return &object.Integer{Value: 0}
	default:
		return object.NULL
	}
}

// infix expression

func evalInfixExpression(left object.Object, operator ast.Operator, right object.Object) object.Object {
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(left.(*object.Integer), operator, right.(*object.Integer))
	case left.Type() == object.BOOLEAN_OBJ && right.Type() == object.BOOLEAN_OBJ:
		return evalBooleanInfixExpression(left.(*object.Boolean), operator, right.(*object.Boolean))
	default:
		return object.NULL
	}
}

func evalBooleanInfixExpression(left *object.Boolean, operator ast.Operator, right *object.Boolean) object.Object {
	lv, rv := left.Value, right.Value
	switch operator {
	case ast.OP_EQ:
		return object.NewBoolean(lv == rv)
	case ast.OP_NEQ:
		return object.NewBoolean(lv != rv)
	default:
		return object.NULL
	}
}

func evalIntegerInfixExpression(left *object.Integer, operator ast.Operator, right *object.Integer) object.Object {
	lv, rv := left.Value, right.Value
	switch operator {
	case ast.OP_PLUS:
		return &object.Integer{Value: lv + rv}
	case ast.OP_MINUS:
		return &object.Integer{Value: lv - rv}
	case ast.OP_MULTIPLY:
		return &object.Integer{Value: lv * rv}
	case ast.OP_DIVIDE:
		return &object.Integer{Value: lv / rv}
	case ast.OP_EQ:
		return object.NewBoolean(lv == rv)
	case ast.OP_NEQ:
		return object.NewBoolean(lv != rv)
	case ast.OP_GT:
		return object.NewBoolean(lv > rv)
	case ast.OP_LT:
		return object.NewBoolean(lv < rv)
	case ast.OP_GE:
		return object.NewBoolean(lv >= rv)
	case ast.OP_LE:
		return object.NewBoolean(lv <= rv)
	default:
		return object.NULL
	}
}
