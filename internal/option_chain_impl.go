package internal

import "github.com/yaadata/optionsgo/core"

type optionChain[T any] struct {
	option core.Option[T]
}

// interface guard
var _ core.OptionChain[uint8] = (*optionChain[uint8])(nil)

func (o *optionChain[T]) Option() core.Option[T] {
	return o.option
}

func (o *optionChain[T]) AndThen(fn func(T) core.Option[any]) core.OptionChain[any] {
	return &optionChain[any]{
		option: OptionAndThen(o.option, fn),
	}
}

func (o *optionChain[T]) Map(fn func(T) any) core.OptionChain[any] {
	return &optionChain[any]{
		option: OptionMap(o.option, fn),
	}
}

func (o *optionChain[T]) MapOr(fn func(T) any, or any) core.OptionChain[any] {
	return &optionChain[any]{
		option: OptionMapOr(o.option, fn, or),
	}
}

func (o *optionChain[T]) MapOrElse(fn func(T) any, orElse func() any) core.OptionChain[any] {
	return &optionChain[any]{
		option: OptionMapOrElse(o.option, fn, orElse),
	}
}

func OptionAndThen[T, V any](option core.Option[T], fn func(T) core.Option[V]) core.Option[V] {
	if option.IsNone() {
		return None[V]()
	}
	return fn(option.Unwrap())
}

func OptionMap[T, V any](option core.Option[T], fn func(value T) V) core.Option[V] {
	if option.IsSome() {
		return Some(fn(option.Unwrap()))
	}
	return None[V]()
}

func OptionMapOr[T, V any](option core.Option[T], fn func(value T) V, or V) core.Option[V] {
	if option.IsSome() {
		return Some(fn(option.Unwrap()))
	}
	return Some(or)
}

func OptionMapOrElse[T, V any](option core.Option[T], fn func(value T) V, orElse func() V) core.Option[V] {
	if option.IsSome() {
		return Some(fn(option.Unwrap()))
	}
	return Some(orElse())
}
