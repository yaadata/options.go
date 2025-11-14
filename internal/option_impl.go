package internal

import (
	"reflect"

	"github.com/yaadata/optionsgo/core"
	"github.com/yaadata/optionsgo/shared"
)

const (
	_FAILED_UNWRAP = "failed to unwrap None value"
)

type option[T any] struct {
	value *T
}

func OptionFromPointer[T any](ptr *T) core.Option[T] {
	return &option[T]{value: ptr}
}

func None[T any]() core.Option[T] {
	return &option[T]{value: nil}
}

func Some[T any](val T) core.Option[T] {
	return &option[T]{value: &val}
}

func (o *option[T]) Option() core.Option[T] {
	return o
}

func (o *option[T]) And(other core.Option[T]) core.Option[T] {
	if o.IsNone() {
		return other
	}
	return o
}

func (o *option[T]) IsSome() bool {
	return o.value != nil
}

func (o *option[T]) IsSomeAnd(pred shared.Predicate[T]) bool {
	if o.value != nil {
		return pred(*o.value)
	}
	return false
}

func (o *option[T]) IsNone() bool {
	return o.value == nil
}

func (o *option[T]) IsNoneOr(pred shared.Predicate[T]) bool {
	if o.value != nil {
		return pred(*o.value)
	}
	return true
}

func (o *option[T]) Equal(other core.Option[T]) bool {
	if o.IsNone() && other.IsNone() || (o == other) {
		return true
	}
	if o.IsSome() && other.IsSome() {
		oValue := *o.value
		otherValue := other.Unwrap()
		if reflect.TypeOf(oValue).Comparable() {
			return reflect.DeepEqual(oValue, otherValue)
		}
	}
	return false
}

func (o *option[T]) Expect(msg string) T {
	if o.value == nil {
		panic(msg)
	}
	return *o.value
}

func (o *option[T]) Unwrap() T {
	return o.Expect(_FAILED_UNWRAP)
}

func (o *option[T]) UnwrapOrElse(fn func() T) T {
	if o.value == nil {
		return fn()
	}
	return *o.value
}

func (o *option[T]) UnwrapOrDefault() T {
	if o.value == nil {
		return *new(T)
	}
	return *o.value
}

func (o *option[T]) AndThen(fn func(T) core.Option[any]) core.Option[any] {
	return OptionAndThen(o, fn)
}

func (o *option[T]) Map(fn func(T) any) core.Option[any] {
	return OptionMap(o, fn)
}

func (o *option[T]) MapOr(fn func(T) any, or any) any {
	return OptionMapOr(o, fn, or)
}

func (o *option[T]) MapOrElse(fn func(T) any, orElse func() any) any {
	return OptionMapOrElse(o, fn, orElse)
}

func (o *option[T]) Filter(pred shared.Predicate[T]) core.Option[T] {
	if o.value != nil && pred(*o.value) {
		return o
	}
	return None[T]()
}

func (o *option[T]) Inspect(fn func(value T)) core.Option[T] {
	if o.IsSome() {
		fn(o.Unwrap())
	}
	return o
}

func (o *option[T]) OkOr(err error) core.Result[T] {
	if o.IsSome() {
		return Ok(*o.value)
	}
	return Err[T](err)
}

func (o *option[T]) OkOrElse(fn func() error) core.Result[T] {
	if o.IsSome() {
		return Ok(*o.value)
	}
	return Err[T](fn())
}

func (o *option[T]) Or(optb core.Option[T]) core.Option[T] {
	if o.IsNone() {
		return optb
	}
	return o
}

func (o *option[T]) OrElse(fn func() core.Option[T]) core.Option[T] {
	if o.IsNone() {
		return fn()
	}
	return o
}

func (o *option[T]) Reduce(optb core.Option[T], fn func(a, b T) T) core.Option[T] {
	if o.IsNone() {
		return optb
	}
	if optb.IsSome() {
		return Some(fn(o.Unwrap(), optb.Unwrap()))
	}
	return o
}

func (o *option[T]) Replace(value T) core.Option[T] {
	o.value = &value
	return o
}

func (o *option[T]) XOr(optb core.Option[T]) core.Option[T] {
	if o.IsSome() {
		if optb.IsSome() {
			return None[T]()
		}
		return o
	}
	if optb.IsSome() {
		return optb
	}
	return None[T]()
}
