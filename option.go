package options

type Option[T any] interface {
	IsSome() bool
	IsSomeAnd(pred Predicate[T]) bool
	IsNone() bool
	IsNoneOr(pred Predicate[T]) bool
	Expect(msg string) T
	Unwrap() T
	UnwrapOrElse(fn func() T) T
	UnwrapOrDefault() T
	OkOr(err error) Result[T]
	OkOrElse(fn func() error) Result[T]
	Or(optb Option[T]) Option[T]
	OrElse(fn func() Option[T]) Option[T]
}

func None[T any]() Option[T] {
	return &option[T]{value: nil}
}

func Some[T any](val T) Option[T] {
	return &option[T]{value: &val}
}
