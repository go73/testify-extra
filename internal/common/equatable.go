package common

type Equatable[T any] interface {
	Equal(T) bool
}
