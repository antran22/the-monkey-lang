package eval_test

import (
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
	for _, tc := range happyCases {
		t.Run(tc.input, func(t *testing.T) {
			evaluated := evalProgram(t, tc.input)
			testObject(t, evaluated, tc.expected)
		})
	}

	for _, tc := range errorCases {
		t.Run(tc.input, func(t *testing.T) {
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

	default:
		t.Fatalf("no tester for value %v of type %T", obj, obj)
	}
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
