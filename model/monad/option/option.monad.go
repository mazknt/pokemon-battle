package option

type None struct {
	error error
}

type Monad[T any] struct {
	Kind  string
	Value T
}

func OfNone[T any]() Monad[T] {
	var zeroValue T
	return Monad[T]{
		Value: zeroValue,
		Kind:  "none",
	}
}

func OfSome[T any](value T) Monad[T] {
	return Monad[T]{
		Value: value,
		Kind:  "some",
	}
}

type fn[T any, K any] func(T) K

func Map[T any, K any](option Monad[T], fn fn[T, K]) Monad[K] {
	if option.Kind == "kind" {
		return OfSome[K](fn(option.Value))
	}
	return OfNone[K]()
}
