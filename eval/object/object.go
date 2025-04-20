package object

type ObjectType string

type HashKey struct {
	Type  ObjectType
	Value uint64
}

type Object interface {
	Type() ObjectType
	IsTruthy() bool
	Inspect() string
}

type Hashable interface {
	Hash() HashKey
}
