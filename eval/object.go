package eval

import (
	"monkey/ast"
	"monkey/eval/object"
)

func evalArrayLiteral(node *ast.ArrayLiteral, env *object.Environment) object.Object {
	elements := evalExpressions(node.Elements, env)

	if len(elements) == 1 && object.IsError(elements[0]) {
		return elements[0]
	}

	return &object.Array{
		Elements: elements,
	}
}

func evalIndexingOperator(node *ast.IndexExpression, env *object.Environment) object.Object {
	leftExp := Eval(node.Left, env)
	if object.IsError(leftExp) {
		return leftExp
	}

	idxExp := Eval(node.Index, env)

	if object.IsError(idxExp) {
		return idxExp
	}

	if leftExp.Type() == object.ARRAY_OBJ {
		return evalArrayIndexing(leftExp.(*object.Array), idxExp)
	}

	return object.UnknownInfixOpError(leftExp, "INDEX", idxExp)
}

func evalArrayIndexing(array *object.Array, index object.Object) object.Object {
	if index.Type() != object.INTEGER_OBJ {
		return object.UnknownInfixOpError(array, "INDEX", index)
	}

	idxInt := index.(*object.Integer).Value

	if idxInt < 0 || idxInt >= len(array.Elements) {
		return object.ArrayOutOfBoundError(idxInt)
	}

	return array.Elements[idxInt]
}
