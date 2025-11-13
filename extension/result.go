package extension

import (
	"github.com/yaadata/optionsgo/core"
	"github.com/yaadata/optionsgo/internal"
)

func ResultFlatten[T any](result core.Result[core.Result[T]]) core.Result[T] {
	if result.IsError() {
		return internal.Err[T](result.UnwrapErr())
	}
	return result.Unwrap()
}

// ResultAnd returns other if result is Ok, otherwise returns the Err from result.
// This is useful for chaining results where you want to proceed with the second
// result only if the first one succeeded.
//
// Example:
//
//	result := internal.Ok(5)
//	other := internal.Ok("OTHER")
//	ResultAnd(result, other) // Returns Ok("OTHER")
//
//	result := internal.Err[int](errors.New("ERROR"))
//	other := internal.Ok("OTHER")
//	ResultAnd(result, other) // Returns Err("ERROR")
func ResultAnd[T, V any](result core.Result[T], other core.Result[V]) core.Result[V] {
	if result.IsOk() {
		return other
	}
	return internal.Err[V](result.UnwrapErr())
}

// ResultAndThen applies fn to the value inside result if it is Ok, otherwise returns the Err.
// This is useful for chaining operations where the next operation depends on the value
// of the first result.
//
// Example:
//
//	result := internal.Ok(5)
//	ResultAndThen(result, func(v int) core.Result[string] {
//	  return internal.Ok(strings.Repeat("A", v))
//	}) // Returns Ok("AAAAA")
//
//	result := internal.Err[int](errors.New("ERROR"))
//	ResultAndThen(result, func(v int) core.Result[string] {
//	  return internal.Ok(strings.Repeat("A", v))
//	}) // Returns Err("ERROR")
func ResultAndThen[T, V any](result core.Result[T], fn func(resultValue T) core.Result[V]) core.Result[V] {
	if result.IsOk() {
		return fn(result.Unwrap())
	}
	return internal.Err[V](result.UnwrapErr())
}

// ResultMap transforms a Result[T] to Result[V] by applying a function to the Ok value.
// If the result is Err, it returns Err[V] with the same error.
//
// Example:
//
//	result := Ok(3)
//	transformed := ResultMap(result, func(value int) string {
//	    return strings.Repeat("A", value)
//	})
//	transformed.Unwrap() // "AAA"
//
//	result := Err[int](errors.New("error"))
//	transformed := ResultMap(result, func(value int) string {
//	    return strings.Repeat("A", value)
//	})
//	transformed.IsError() // true
func ResultMap[T, V any](result core.Result[T], fn func(inner T) V) core.Result[V] {
	return internal.ResultMap(result, fn)
}

// ResultMapOr transforms a Result[T] to Result[V] by applying a function to the Ok value,
// or returning a default value if the result is Err.
// If the result is Ok, applies fn to the Ok value and returns Ok with the transformed value.
// If the result is Err, returns Ok with the provided default value 'or'.
//
// Example:
//
//	result := Ok(3)
//	transformed := ResultMapOr(result, func(value int) string {
//	    return strings.Repeat("A", value)
//	}, "DEFAULT")
//	transformed.Unwrap() // "AAA"
//
//	result := Err[int](errors.New("error"))
//	transformed := ResultMapOr(result, func(value int) string {
//	    return strings.Repeat("A", value)
//	}, "DEFAULT")
//	transformed.Unwrap() // "DEFAULT"
func ResultMapOr[T, V any](result core.Result[T], fn func(inner T) V, or V) core.Result[V] {
	return internal.ResultMapOr(result, fn, or)
}

// ResultMapOrElse transforms a Result[T] to Result[V] by applying a function to the Ok value,
// or using an alternative function if the result is Err.
// If the result is Ok, applies fn to the Ok value.
// If the result is Err, applies orElse to the error and returns Ok with that value.
//
// Example:
//
//	result := Ok(3)
//	transformed := ResultMapOrElse(result,
//	    func(value int) string {
//	        return strings.Repeat("A", value)
//	    },
//	    func(err error) string {
//	        return "OTHER"
//	    },
//	)
//	transformed.Unwrap() // "AAA"
//
//	result := Err[int](errors.New("error"))
//	transformed := ResultMapOrElse(result,
//	    func(value int) string {
//	        return strings.Repeat("A", value)
//	    },
//	    func(err error) string {
//	        return "EXPECTED"
//	    },
//	)
//	transformed.Unwrap() // "EXPECTED"
func ResultMapOrElse[T, V any](result core.Result[T], fn func(inner T) V, orElse func(error) V) core.Result[V] {
	return internal.ResultMapOrElse(result, fn, orElse)
}

func ResultTranspose[T any](result core.Result[core.Option[T]]) core.Option[core.Result[T]] {
	if result.IsError() {
		return internal.Some(internal.Err[T](result.UnwrapErr()))
	}
	option := result.Unwrap()
	if option.IsNone() {
		return internal.None[core.Result[T]]()
	}
	return internal.Some(internal.Ok(option.Unwrap()))
}
