package eval

import (
	"monkey/ast"
	"monkey/eval/object"
)

func evalProgram(stmts []ast.Statement, env *object.Environment) object.Object {
	var result object.Object

	for _, stmt := range stmts {
		result = Eval(stmt, env)
		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}
	return result
}

func evalBlockStatements(stmts []ast.Statement, env *object.Environment) object.Object {
	var result object.Object

	for _, stmt := range stmts {
		result = Eval(stmt, env)
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
