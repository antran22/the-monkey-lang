package eval_test

import (
	"errors"
	"fmt"
	"monkey/eval"
	"monkey/eval/object"
	"monkey/lexer"
	"monkey/parser"
	"testing"

	"github.com/stretchr/testify/require"
)

func evalProgram(t *testing.T, input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	errors := p.Errors()
	if len(errors) != 0 {
		t.Errorf("parser has %d errors", len(errors))
		for _, msg := range errors {
			t.Errorf("parser error: %q", msg)
		}
		t.FailNow()
	}

	env := object.NewEnvironment()

	return eval.Eval(program, env)
}

type happyTestCase struct {
	input    string
	expected any
}

type errorTestCase struct {
	input           string
	expectedMessage string
}

func testExpressionEvaluation(t *testing.T, happyCases []happyTestCase, errorCases []errorTestCase) {
	for i, tc := range happyCases {
		t.Run(fmt.Sprintf("happy case #%d", i), func(t *testing.T) {
			evaluated := evalProgram(t, tc.input)
			testObject(t, evaluated, tc.expected)
		})
	}

	for i, tc := range errorCases {
		t.Run(fmt.Sprintf("error case #%d", i), func(t *testing.T) {
			evaluated := evalProgram(t, tc.input)
			testErrorObject(t, evaluated, tc.expectedMessage)
		})
	}
}

func testObject(t *testing.T, obj object.Object, expected any) {
	switch expected := expected.(type) {
	case int:
		testIntegerObject(t, obj, expected)
	case nil:
		testNullObject(t, obj)
	case string:
		testStringObject(t, obj, expected)
	case bool:
		testBooleanObject(t, obj, expected)
	case []any:
		testArrayObject(t, obj, expected)
	case map[any]any:
		testHashObject(t, obj, expected)
	case *object.Range:
		testRangeObject(t, obj, expected)
	default:
		t.Fatalf("no tester for value %v of type %T", obj, obj)
	}
}

func unwrapPrimitiveObj(obj object.Object) (any, error) {
	switch obj := obj.(type) {
	case *object.Boolean:
		return obj.Value, nil
	case *object.Integer:
		return obj.Value, nil
	case *object.String:
		return obj.Value, nil
	case *object.Null:
		return nil, nil
	}
	return nil, errors.New("invalid primitive object")
}

func testIntegerObject(t *testing.T, obj object.Object, expected int) {
	r := require.New(t)
	result, ok := obj.(*object.Integer)
	r.Truef(ok, "obj is not *object.Integer, got %T (%v)", obj, obj)

	r.Equal(expected, result.Value)
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) {
	r := require.New(t)
	result, ok := obj.(*object.Boolean)
	r.Truef(ok, "obj is not *object.Boolean, got %T (%v)", obj, obj)

	r.Equal(expected, result.Value)
}

func testArrayObject(t *testing.T, obj object.Object, expected []any) {
	r := require.New(t)
	array, ok := obj.(*object.Array)
	r.Truef(ok, "obj is not *object.Array, got %T (%v)", obj, obj)

	r.Len(array.Elements, len(expected))

	for i, el := range array.Elements {
		testObject(t, el, expected[i])
	}
}

func testHashObject(t *testing.T, obj object.Object, expected map[any]any) {
	r := require.New(t)

	hash, ok := obj.(*object.Hash)
	r.Truef(ok, "obj is not *object.Hash, got %T (%v)", obj, obj)

	r.Equal(len(expected), len(hash.Pairs))
	for _, pair := range hash.Pairs {
		keyPrim, err := unwrapPrimitiveObj(pair.Key)
		r.Nil(err)
		testObject(t, pair.Value, expected[keyPrim])
	}
}

func testRangeObject(t *testing.T, obj object.Object, expected *object.Range) {
	r := require.New(t)

	ran, ok := obj.(*object.Range)
	r.Truef(ok, "obj is not *object.Range, got %T (%v)", obj, obj)

	r.Equal(expected.Start, ran.Start)
	r.Equal(expected.End, ran.End)
	r.Equal(expected.Step, ran.Step)
}

func testStringObject(t *testing.T, obj object.Object, expected string) {
	r := require.New(t)
	result, ok := obj.(*object.String)

	r.Truef(ok, "obj is not *object.String, got %T (%v)", obj, obj)
	r.Equal(expected, result.Value)
}

func testNullObject(t *testing.T, obj object.Object) {
	require.Equal(t, object.NULL, obj)
}

func testErrorObject(t *testing.T, obj object.Object, message string) {
	r := require.New(t)
	errObj, ok := obj.(*object.Error)

	r.Truef(ok, "obj is not *object.Error, got %T (%v)", obj, obj)

	r.Equal(message, errObj.Message)
}

func assertIsArrayObject(t *testing.T, obj object.Object) *object.Array {
	result, ok := obj.(*object.Array)

	require.Truef(t, ok, "obj is not *object.Array, got %T (%v)", obj, obj)

	return result
}
