package opt

import (
	"errors"
)

type Optional[T any] struct {
	value *T
}

func MakeOptional[T any](value *T) Optional[T] {
	return Optional[T]{value: value}
}

func MakeEmpty[T any]() Optional[T] {
	var nothing T
	return MakeOptional(&nothing)
}

func (o* Optional[T]) Unwrap() (T, error) {
	if o.value == nil {
		var nothing T
		return nothing, errors.New("nothing")
	}
	return *o.value, nil
}

func (o* Optional[T]) ForceUnwrap() T {
	if o.value == nil {
		panic("Attemped to force unwrap, but value was nil")
	}
	return *o.value
}

func If[T any, R any](optional Optional[T], handler func(T) R) Optional[R] {
	if item, err := optional.Unwrap(); err == nil {
		returnVal := handler(item)
		return MakeOptional[R](&returnVal)
	}
	return MakeEmpty[R]()
}
