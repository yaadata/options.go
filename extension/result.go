package extension

import (
	"github.com/yaadata/optionsgo/core"
	"github.com/yaadata/optionsgo/internal"
)

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
	if result.IsOk() {
		return internal.Ok(fn(result.Unwrap()))
	}
	return internal.Err[V](result.UnwrapErr())
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
	if result.IsOk() {
		return internal.Ok(fn(result.Unwrap()))
	}
	return internal.Ok(or)
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
	if result.IsOk() {
		return internal.Ok(fn(result.Unwrap()))
	}
	return internal.Ok(orElse(result.UnwrapErr()))
}
