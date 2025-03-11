package result

import (
	"errors"
	"log"
)

type Err struct {
	error
}
type Result[T any] struct {
	value T
	error error
	IsOk  bool
}

func OfErr[T any](err error) Result[T] {
	var zeroValue T
	return Result[T]{
		value: zeroValue,
		error: err,
		IsOk:  false,
	}
}

func OfOk[T any](value T) Result[T] {
	return Result[T]{
		value: value,
		error: errors.New(""),
		IsOk:  true,
	}
}

func Unwrap[T any](result Result[T]) T {
	if !result.IsOk {
		log.Println("Error:", result.error)
		var zeroValue T
		return zeroValue
	}
	return result.value
}
