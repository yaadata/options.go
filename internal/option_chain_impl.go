package internal

import "github.com/yaadata/optionsgo/core"

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
