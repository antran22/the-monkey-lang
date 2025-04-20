package object_test

import (
	"monkey/eval/object"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntegerHashKey(t *testing.T) {
	r := require.New(t)
	one1 := object.NewInt(1)
	one2 := object.NewInt(1)

	r.Equal(one1.Hash(), one2.Hash())

	two := object.NewInt(2)

	r.NotEqual(one1.Hash(), two.Hash())
}

func TestStringHashKey(t *testing.T) {
	r := require.New(t)

	hello1 := object.NewString("Hello World")
	hello2 := object.NewString("Hello World")

	r.Equal(hello1.Hash(), hello2.Hash())

	diff1 := object.NewString("My name is Johnny")
	diff2 := object.NewString("My name is johnny")

	r.NotEqual(diff1.Value, diff2.Value)
}
