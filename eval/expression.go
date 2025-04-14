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
	truthy := value.IsTruthy()
	return object.NewBoolean(!truthy)
}

func evalPrefixMinusOperator(value object.Object) object.Object {
	t := value.Type()

	switch t {
	case object.INTEGER_OBJ:
		iv := value.(*object.Integer).Value
		return &object.Integer{Value: -iv}
	default:
		return object.NewErrorf("unsupported operation: - %s", t)
	}
}

// infix expression

func evalInfixExpression(left object.Object, operator ast.Operator, right object.Object) object.Object {
	lt, rt := left.Type(), right.Type()
	switch {
	case lt == object.INTEGER_OBJ && rt == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(left.(*object.Integer), operator, right.(*object.Integer))
	case lt == object.BOOLEAN_OBJ && rt == object.BOOLEAN_OBJ:
		return evalBooleanInfixExpression(left.(*object.Boolean), operator, right.(*object.Boolean))
	default:
		return object.NewErrorf("unsupported operation: %s %s %s", lt, operator, rt)
	}
}

func evalBooleanInfixExpression(left *object.Boolean, operator ast.Operator, right *object.Boolean) object.Object {
	switch operator {
	case ast.OP_EQ:
		return object.NewBoolean(left == right)
	case ast.OP_NEQ:
		return object.NewBoolean(left != right)
	case ast.OP_AND:
		return object.NewBoolean(left.Value && right.Value)
	case ast.OP_OR:
		return object.NewBoolean(left.Value || right.Value)
	default:
		return object.NewErrorf("unsupported operation: BOOLEAN %s BOOLEAN", operator)
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
	case ast.OP_BITWISE_AND:
		return &object.Integer{Value: lv & rv}
	case ast.OP_BITWISE_OR:
		return &object.Integer{Value: lv | rv}
	case ast.OP_BITWISE_XOR:
		return &object.Integer{Value: lv ^ rv}

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
		return object.NewErrorf("unsupported operation: INTEGER %s INTEGER", operator)
	}
}
