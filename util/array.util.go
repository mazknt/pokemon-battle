package util

import (
	"reflect"

	O "github.com/IBM/fp-go/option"
)

func IsExist[T any](slice []T) O.Option[T] {
	if len(slice) == 0 {
		return O.None[T]()
	}
	return O.Some(slice[0])
}

func GetValueOfElement[T any](slice []*T) []T {
	clients := make([]T, len(slice)) // 値のスライスを作成
	for i, clientPtr := range slice {
		if clientPtr != nil { // nil チェック（安全のため）
			clients[i] = *clientPtr // ポインタをデリファレンスして値を格納
		}
	}
	return clients
}

func Push[T any](slice []T) func(element T) []T {
	return func(element T) []T {
		return append(slice, element)
	}
}

func Update[T any](slice []T, oldElement T) func(newElement T) []T {
	return func(newElement T) []T {
		for index, el := range slice {
			reflect.DeepEqual(el, oldElement)
			slice[index] = newElement
		}
		return slice
	}
}
