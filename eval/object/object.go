package object

type ObjectType string

type Object interface {
	Type() ObjectType
	IsTruthy() bool
	Inspect() string
}
