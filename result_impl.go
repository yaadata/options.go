package options

type result[T any] struct {
	value *T
	err   error
}

// interface guard
var _ Result[string] = (*result[string])(nil)

func (r *result[T]) Ok() Option[T] {
	if r.value == nil {
		return None[T]()
	}
	return Some(*r.value)
}

func (r *result[T]) Err() Option[error] {
	if r.err == nil {
		return None[error]()
	}
	return Some(r.err)
}

func (r *result[T]) IsOk() bool {
	return r.value != nil
}

func (r *result[T]) IsOkAnd(pred Predicate[T]) bool {
	if r.IsOk() {
		return pred(*r.value)
	}
	return false
}

func (r *result[T]) IsError() bool {
	return r.err != nil
}

func (r *result[T]) IsErrorAnd(pred Predicate[error]) bool {
	if r.IsError() {
		return pred(r.err)
	}
	return false
}

func (r *result[T]) Unwrap() T {
	if r.value == nil {
		panic("cannot unwrap Err result to value")
	}
	return *r.value
}

func (r *result[T]) UnwrapErr() error {
	if r.err == nil {
		panic("cannot unwrap Ok result to error")
	}
	return r.err
}

func (r *result[T]) UnwrapOr(val T) T {
	if r.value == nil {
		return val
	}
	return *r.value
}

func (r *result[T]) UnwrapOrElse(fn func() T) T {
	if r.value == nil {
		return fn()
	}
	return *r.value
}

func (r *result[T]) UnwrapOrDefault() T {
	if r.value == nil {
		return *new(T)
	}
	return *r.value
}
