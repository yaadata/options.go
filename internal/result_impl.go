package internal

import "github.com/yaadata/optionsgo/core"

type result[T any] struct {
	value *T
	err   error
}

// interface guard
var _ core.Result[string] = (*result[string])(nil)

func ResultFromReturn[T any](value T, err error) core.Result[T] {
	if err != nil {
		return Err[T](err)
	}
	return &result[T]{
		value: &value,
		err:   nil,
	}
}

func Err[T any](err error) core.Result[T] {
	return &result[T]{
		value: nil,
		err:   err,
	}
}

func Ok[T any](value T) core.Result[T] {
	return &result[T]{
		value: &value,
		err:   nil,
	}
}

func (r *result[T]) Ok() core.Option[T] {
	if r.value == nil {
		return None[T]()
	}
	return Some(*r.value)
}

func (r *result[T]) Expect(msg string) T {
	if r.IsError() {
		panic(msg)
	}
	return *r.value
}

func (r *result[T]) ExpectErr(msg string) error {
	if r.IsOk() {
		panic(msg)
	}
	return r.err
}

func (r *result[T]) Err() core.Option[error] {
	if r.err == nil {
		return None[error]()
	}
	return Some(r.err)
}

func (r *result[T]) IsOk() bool {
	return r.value != nil
}

func (r *result[T]) IsOkAnd(pred core.Predicate[T]) bool {
	if r.IsOk() {
		return pred(*r.value)
	}
	return false
}

func (r *result[T]) IsError() bool {
	return r.err != nil
}

func (r *result[T]) IsErrorAnd(pred core.Predicate[error]) bool {
	if r.IsError() {
		return pred(r.err)
	}
	return false
}

func (r *result[T]) Inspect(fn func(value T)) core.Result[T] {
	if r.IsOk() {
		fn(r.Unwrap())
	}
	return r
}

func (r *result[T]) InspectErr(fn func(err error)) core.Result[T] {
	if r.IsError() {
		fn(r.UnwrapErr())
	}
	return r
}

func (r *result[T]) Map(fn func(value T) any) core.Result[any] {
	return ResultMap(r, fn)
}

func (r *result[T]) MapOr(fn func(value T) any, or any) core.Result[any] {
	return ResultMapOr(r, fn, or)
}

func (r *result[T]) MapOrElse(fn func(value T) any, orElse func(err error) any) core.Result[any] {
	return ResultMapOrElse(r, fn, orElse)
}

func (r *result[T]) MapErr(fn func(inner error) error) core.Result[T] {
	if r.IsError() {
		return Err[T](fn(r.UnwrapErr()))
	}
	return r
}

func (r *result[T]) Or(other core.Result[T]) core.Result[T] {
	if r.IsError() {
		return other
	}
	return r
}

func (r *result[T]) OrElse(fn func(err error) core.Result[T]) core.Result[T] {
	if r.IsError() {
		return fn(r.UnwrapErr())
	}
	return r
}

func (r *result[T]) Unwrap() T {
	return r.Expect("cannot unwrap Err result to value")
}

func (r *result[T]) UnwrapErr() error {
	return r.ExpectErr("cannot unwrap Ok result to error")
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
