package extension

import (
	"github.com/yaadata/optionsgo/core"
	"github.com/yaadata/optionsgo/internal"
)

// OptionMap transforms an Option[T] to Option[V] by applying a function to the contained value.
// If the option is Some, applies fn to the value and returns Some with the transformed value.
// If the option is None, returns None[V].
//
// Example:
//
//	option := Some(3)
//	transformed := OptionMap(option, func(value int) string {
//	    return strings.Repeat("A", value)
//	})
//	transformed.Unwrap() // "AAA"
//
//	option := None[int]()
//	transformed := OptionMap(option, func(value int) string {
//	    return strings.Repeat("A", value)
//	})
//	transformed.IsNone() // true
func OptionMap[T, V any](option core.Option[T], fn func(value T) V) core.Option[V] {
	if option.IsSome() {
		return internal.Some(fn(option.Unwrap()))
	}
	return internal.None[V]()
}

// OptionMapOr transforms an Option[T] to Option[V] by applying a function to the contained value,
// or returning a default value if the option is None.
// If the option is Some, applies fn to the value and returns Some with the transformed value.
// If the option is None, returns Some with the provided default value 'or'.
//
// Example:
//
//	option := Some(3)
//	transformed := OptionMapOr(option, func(value int) string {
//	    return strings.Repeat("A", value)
//	}, "DEFAULT")
//	transformed.Unwrap() // "AAA"
//
//	option := None[int]()
//	transformed := OptionMapOr(option, func(value int) string {
//	    return strings.Repeat("A", value)
//	}, "DEFAULT")
//	transformed.Unwrap() // "DEFAULT"
func OptionMapOr[T, V any](option core.Option[T], fn func(value T) V, or V) core.Option[V] {
	if option.IsSome() {
		return internal.Some(fn(option.Unwrap()))
	}
	return internal.Some(or)
}

// OptionMapOrElse transforms an Option[T] to Option[V] by applying a function to the contained value,
// or computing a default value if the option is None.
// If the option is Some, applies fn to the value and returns Some with the transformed value.
// If the option is None, calls orElse to compute a default value and returns Some with that value.
//
// Example:
//
//	option := Some(3)
//	transformed := OptionMapOrElse(option, func(value int) string {
//	    return strings.Repeat("A", value)
//	}, func() string {
//	    return "DEFAULT"
//	})
//	transformed.Unwrap() // "AAA"
//
//	option := None[int]()
//	transformed := OptionMapOrElse(option, func(value int) string {
//	    return strings.Repeat("A", value)
//	}, func() string {
//	    return "DEFAULT"
//	})
//	transformed.Unwrap() // "DEFAULT"
func OptionMapOrElse[T, V any](option core.Option[T], fn func(value T) V, orElse func() V) core.Option[V] {
	if option.IsSome() {
		return internal.Some(fn(option.Unwrap()))
	}
	return internal.Some(orElse())
}
