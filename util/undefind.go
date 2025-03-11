package util

import (
	"reflect"

	FP "github.com/IBM/fp-go/function"
	O "github.com/IBM/fp-go/option"
)

func IsUndefined[T any](param T) bool {
	var zero T
	return reflect.DeepEqual(zero, param)
}

func Undefined[T any](opt O.Option[T]) T {
	var zero T
	return FP.Pipe1(
		opt,
		O.Fold(
			func() T { return zero },
			func(el T) T { return el },
		),
	)
}
