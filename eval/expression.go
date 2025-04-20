package eval

import (
	"maps"
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
		return object.UnknownPrefixOpError("-", value)
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
	case lt == object.STRING_OBJ && rt == object.STRING_OBJ:
		return evalStringInfixExpression(left.(*object.String), operator, right.(*object.String))
	case lt == object.ARRAY_OBJ && rt == object.ARRAY_OBJ:
		return evalArrayInfixExpression(left.(*object.Array), operator, right.(*object.Array))
	case lt == object.HASH_OBJ && rt == object.HASH_OBJ:
		return evalHashInfixExpression(left.(*object.Hash), operator, right.(*object.Hash))
	default:
		return object.UnknownInfixOpError(left, string(operator), right)
	}
}

func evalHashInfixExpression(left *object.Hash, operator ast.Operator, right *object.Hash) object.Object {
	if operator != ast.OP_PLUS {
		return object.UnknownInfixOpError(left, string(operator), right)
	}
	r := map[object.HashKey]object.HashPair{}
	maps.Copy(r, left.Pairs)
	maps.Copy(r, right.Pairs)
	return &object.Hash{
		Pairs: r,
	}
}

func evalArrayInfixExpression(left *object.Array, operator ast.Operator, right *object.Array) object.Object {
	if operator != ast.OP_PLUS {
		return object.UnknownInfixOpError(left, string(operator), right)
	}

	return &object.Array{
		Elements: append(left.Elements, right.Elements...),
	}
}

func evalStringInfixExpression(left *object.String, operator ast.Operator, right *object.String) object.Object {
	switch operator {
	case ast.OP_EQ:
		return object.NewBoolean(left.Value == right.Value)
	case ast.OP_NEQ:
		return object.NewBoolean(left.Value != right.Value)
	case ast.OP_PLUS:
		return &object.String{Value: left.Value + right.Value}
	default:
		return object.UnknownInfixOpError(left, string(operator), right)
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
		return object.UnknownInfixOpError(left, string(operator), right)
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
		return object.UnknownInfixOpError(left, string(operator), right)
	}
}
