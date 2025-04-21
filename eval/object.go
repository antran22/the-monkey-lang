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

func evalHashLiteral(node *ast.HashLiteral, env *object.Environment) object.Object {
	res := &object.Hash{
		Pairs: map[object.HashKey]object.HashPair{},
	}

	for _, astPair := range node.Pairs {
		keyObj := Eval(astPair.Key, env)
		if object.IsError(keyObj) {
			return keyObj
		}

		keyObjH, ok := keyObj.(object.Hashable)
		if !ok {
			return object.TypeNotHashable(keyObj.Type())
		}

		valObj := Eval(astPair.Value, env)

		if object.IsError(valObj) {
			return valObj
		}

		hashKey := keyObjH.Hash()

		res.Pairs[hashKey] = object.HashPair{
			Key:   keyObj,
			Value: valObj,
		}
	}
	return res
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

	switch leftExp.Type() {
	case object.ARRAY_OBJ:
		return evalArrayIndexing(leftExp.(*object.Array), idxExp)
	case object.HASH_OBJ:
		return evalHashIndexing(leftExp.(*object.Hash), idxExp)
	default:
		return object.UnknownInfixOpError(leftExp, "INDEX", idxExp)

	}
}

func evalArrayIndexing(array *object.Array, index object.Object) object.Object {
	switch index.Type() {
	case object.INTEGER_OBJ:
		idxInt := index.(*object.Integer).Value
		return array.Index(idxInt)
	case object.RANGE_OBJ:
		r := index.(*object.Range)
		return array.Slice(r.Start, r.End, r.Step)
	default:
		return object.UnknownInfixOpError(array, "INDEX", index)
	}
}

func evalHashIndexing(hash *object.Hash, index object.Object) object.Object {
	hashable, ok := index.(object.Hashable)
	if !ok {
		return object.TypeNotHashable(index.Type())
	}
	hashKey := hashable.Hash()
	if pair, ok := hash.Pairs[hashKey]; !ok {
		return object.NULL
	} else {
		return pair.Value
	}
}
