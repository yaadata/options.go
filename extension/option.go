package extension

import (
	"github.com/yaadata/optionsgo/core"
	"github.com/yaadata/optionsgo/internal"
)

// OptionFromPointer converts a pointer to an Option.
// The Option retains the original pointer.
//
// Example:
//
//	var ptr *string // nil pointer
//	opt := OptionFromPointer(ptr) // returns None
//	opt.IsNone() // true
//
//	value := "hello"
//	ptr = &value
//	opt = OptionFromPointer(ptr) // returns Some("hello")
//	opt.IsSome() // true
func OptionFromPointer[T any](ptr *T) core.Option[T] {
	return internal.OptionFromPointer(ptr)
}

// OptionFlatten removes one level of nesting from a nested Option.
// It converts Option[Option[T]] into Option[T].
//
// If the outer Option is None, it returns None.
// If the outer Option is Some(inner), it returns the inner Option.
//
// Note: OptionFlatten only flattens one level at a time. For deeply nested
// Options, multiple calls are required.
//
// Examples:
//
//	// Some(Some(5)) -> Some(5)
//	option := Some(Some(5))
//	result := OptionFlatten(option) // Some(5)
//
//	// Some(None) -> None
//	option := Some(None[int]())
//	result := OptionFlatten(option) // None
//
//	// None -> None
//	option := None[Option[int]]()
//	result := OptionFlatten(option) // None
//
//	// Flattens only one level deep
//	option := Some(Some(Some(5)))
//	result := OptionFlatten(OptionFlatten(option)) // Some(5)
func OptionFlatten[T any](option core.Option[core.Option[T]]) core.Option[T] {
	if option.IsNone() {
		return internal.None[T]()
	}
	return option.Unwrap()
}

// OptionAndThen transforms an Option[T] to Option[V] by applying a function to the contained value.
// If the option is Some, applies fn to the value and returns Some with the transformed value.
// If the option is None, returns None[V].
//
// This function is similar to OptionMap and is useful for chaining operations that transform values.
//
// Example:
//
//	option := Some(3)
//	transformed := OptionAndThen(option, func(value int) string {
//	    return strings.Repeat("A", value)
//	})
//	transformed.Unwrap() // "AAA"
//
//	option := None[int]()
//	transformed := OptionAndThen(option, func(value int) string {
//	    return strings.Repeat("A", value)
//	})
//	transformed.IsNone() // true
func OptionAndThen[T, V any](option core.Option[T], fn func(T) core.Option[V]) core.Option[V] {
	return internal.OptionAndThen(option, fn)
}

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
	return internal.OptionMap(option, fn)
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
func OptionMapOr[T, V any](option core.Option[T], fn func(value T) V, or V) V {
	return internal.OptionMapOr(option, fn, or)
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
func OptionMapOrElse[T, V any](option core.Option[T], fn func(value T) V, orElse func() V) V {
	if option.IsSome() {
		return fn(option.Unwrap())
	}
	return orElse()
}

// OptionTranspose converts an Option[Result[T]] into a Result[Option[T]].
// It "transposes" the nested types, swapping the order of Option and Result.
//
// This is useful when you have an optional result and want to convert it into
// a result containing an optional value, making error handling the outer concern.
//
// Behavior:
//   - None -> Ok(None)
//   - Some(Ok(value)) -> Ok(Some(value))
//   - Some(Err(error)) -> Err(error)
//
// Examples:
//
//	// None => Ok(None)
//	option := None[Result[int]]()
//	result := OptionTranspose(option) // Ok(None)
//	result.IsOk() // true
//
//	// Some(Ok(5)) => Ok(Some(5))
//	option := Some(Ok(5))
//	result := OptionTranspose(option) // Ok(Some(5))
//	result.Unwrap().Unwrap() // 5
//
//	// Some(Err) => Err
//	option := Some(Err[int](errors.New("ERROR")))
//	result := OptionTranspose(option) // Err("ERROR")
//	result.IsError() // true
func OptionTranspose[T any](option core.Option[core.Result[T]]) core.Result[core.Option[T]] {
	if option.IsNone() {
		return internal.Ok(internal.None[T]())
	}
	result := option.Unwrap()
	if result.IsError() {
		return internal.Err[core.Option[T]](result.UnwrapErr())
	}
	return internal.Ok(internal.Some(result.Unwrap()))
}
