package object_test

import (
	"monkey/eval/object"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArray(t *testing.T) {
	r := require.New(t)

	arr := &object.Array{
		Elements: []object.Object{
			object.TRUE,
			object.NewInt(2),
			object.NewInt(4),
		},
	}

	r.True(arr.IsTruthy())
	r.Equal(object.TRUE, arr.Index(0))

	r.Equal(2, arr.Index(1).(*object.Integer).Value)
	r.Equal(4, arr.Index(2).(*object.Integer).Value)

	arr2 := arr.Slice(0, 1, 1).(*object.Array)

	r.Equal(1, len(arr2.Elements))
}
