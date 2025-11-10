package options

type Result[T any] interface {
	Ok() Option[T]
	Err() Option[error]
	IsOk() bool
	IsOkAnd(pred Predicate[T]) bool
	IsError() bool
	IsErrorAnd(pred Predicate[error]) bool
	Unwrap() T
	UnwrapErr() error
	UnwrapOr(value T) T
	UnwrapOrElse(fn func() T) T
	UnwrapOrDefault() T
}

func ResultFromReturn[T any](value T, err error) Result[T] {
	if err != nil {
		return Err[T](err)
	}
	return &result[T]{
		value: &value,
		err:   nil,
	}
}

func Err[T any](err error) Result[T] {
	return &result[T]{
		value: nil,
		err:   err,
	}
}

func Ok[T any](value T) Result[T] {
	return &result[T]{
		value: &value,
		err:   nil,
	}
}
