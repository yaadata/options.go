package core

// Result represents the outcome of an operation that can either succeed with a value
// or fail with an error. This is a Go implementation of Rust's std::result::Result type.
//
// A Result is either:
//   - Ok: contains a successful value of type T
//   - Err: contains an error
type Result[T any] interface {
	ResultChain[T]
	ResultToOption[T]

	Expect(msg string) T

	ExpectErr(msg string) error

	// IsOk returns true if the result is Ok.
	//
	// Example:
	//
	//  result := Ok("value")
	//  result.IsOk() // true
	//
	//  result := Err[string](errors.New("error"))
	//  result.IsOk() // false
	IsOk() bool

	// IsOkAnd returns true if the result is Ok and the predicate returns true for the value.
	//
	// Example:
	//
	//  result := Ok("value")
	//  result.IsOkAnd(func(s string) bool { return len(s) > 0 }) // true
	//  result.IsOkAnd(func(s string) bool { return len(s) == 0 }) // false
	//
	//  result := Err[string](errors.New("error"))
	//  result.IsOkAnd(func(s string) bool { return true }) // false
	IsOkAnd(pred Predicate[T]) bool

	// IsError returns true if the result is Err.
	//
	// Example:
	//
	//  result := Err[string](errors.New("error"))
	//  result.IsError() // true
	//
	//  result := Ok("value")
	//  result.IsError() // false
	IsError() bool

	// IsErrorAnd returns true if the result is Err and the predicate returns true for the error.
	//
	// Example:
	//
	//  result := Err[string](errors.New("error_message"))
	//  result.IsErrorAnd(func(e error) bool { return e.Error() == "error_message" }) // true
	//  result.IsErrorAnd(func(e error) bool { return e.Error() == "other" }) // false
	//
	//  result := Ok("value")
	//  result.IsErrorAnd(func(e error) bool { return true }) // false
	IsErrorAnd(pred Predicate[error]) bool

	// Unwrap returns the contained Ok value.
	// Panics if the result is Err.
	//
	// Example:
	//
	//  result := Ok("EXPECTED")
	//  value := result.Unwrap() // "EXPECTED"
	//
	//  result := Err[string](errors.New("error"))
	//  value := result.Unwrap() // panics!
	Unwrap() T

	// UnwrapErr returns the contained Err value.
	// Panics if the result is Ok.
	//
	// Example:
	//
	//  err := errors.New("error_message")
	//  result := Err[string](err)
	//  e := result.UnwrapErr() // err
	//
	//  result := Ok("value")
	//  e := result.UnwrapErr() // panics!
	UnwrapErr() error

	// UnwrapOrDefault returns the contained Ok value or the zero value of type T.
	//
	// Example:
	//  result := Ok("EXPECTED")
	//  value := result.UnwrapOrDefault() // "EXPECTED"
	//
	//  result := Err[string](errors.New("error"))
	//  value := result.UnwrapOrDefault() // ""
	UnwrapOrDefault() T

	// UnwrapOr returns the contained Ok value or the provided default value.
	//
	// Example:
	//
	//  result := Ok("EXPECTED")
	//  value := result.UnwrapOr("default") // "EXPECTED"
	//
	//  result := Err[string](errors.New("error"))
	//  value := result.UnwrapOr("default") // "default"
	UnwrapOr(value T) T

	// UnwrapOrElse returns the contained Ok value or computes it from the provided function.
	//
	// Example:
	//
	//  result := Ok("EXPECTED")
	//  value := result.UnwrapOrElse(func() string { return "default" }) // "EXPECTED"
	//
	//  result := Err[string](errors.New("error"))
	//  value := result.UnwrapOrElse(func() string { return "default" }) // "default"
	UnwrapOrElse(fn func() T) T
}

type ResultChain[T any] interface {
	Inspect(fn func(value T)) Result[T]
	InspectErr(fn func(err error)) Result[T]
	// Map transforms a Result[T] to Result[any] by applying a function to the Ok value.
	// If the result is Err, it returns Err with the same error unchanged.
	//
	// Example:
	//
	//  result := Ok("parallel")
	//  transformed := result.Map(func(value string) any {
	//      return len(value)
	//  })
	//  transformed.Unwrap() // 8
	//
	//  result := Err[string](errors.New("error"))
	//  transformed := result.Map(func(value string) any {
	//      return len(value)
	//  })
	//  transformed.IsError() // true
	Map(fn func(value T) any) Result[any]
	// MapErr applies a transformation function to the error if the Result is Err.
	// If the Result is Ok, it returns the Result unchanged.
	//
	// Example:
	//
	//  result := Err[string](errors.New("A"))
	//  transformed := result.MapErr(func(err error) error {
	//      return fmt.Errorf("%s - B", err.Error())
	//  })
	//  transformed.UnwrapErr().Error() // "A - B"
	//
	//  result := Ok(15)
	//  transformed := result.MapErr(func(err error) error {
	//      return fmt.Errorf("transformed: %v", err)
	//  })
	//  transformed.Unwrap() // 15 (unchanged)
	MapErr(func(error) error) Result[T]
	// MapOr transforms the Ok value using fn, or returns Ok(or) if the Result is Err.
	// Unlike Map which propagates errors, MapOr always returns an Ok result.
	//
	// Example:
	//
	//  result := Ok(3)
	//  transformed := result.MapOr(func(v int) any { return v * 2 }, 0)
	//  transformed.Unwrap() // 6
	//
	//  result := Err[int](errors.New("error"))
	//  transformed := result.MapOr(func(v int) any { return v * 2 }, 0)
	//  transformed.Unwrap() // 0
	MapOr(fn func(value T) any, or any) Result[any]
	// MapOrElse transforms the Ok value using fn, or computes a value from the error using orElse.
	// Always returns an Ok result, either from transforming the success value or handling the error.
	//
	// Example:
	//
	//  result := Ok(3)
	//  transformed := result.MapOrElse(
	//      func(v int) any { return v * 2 },
	//      func(e error) any { return -1 },
	//  )
	//  transformed.Unwrap() // 6
	//
	//  result := Err[int](errors.New("error"))
	//  transformed := result.MapOrElse(
	//      func(v int) any { return v * 2 },
	//      func(e error) any { return -1 },
	//  )
	//  transformed.Unwrap() // -1
	MapOrElse(fn func(value T) any, orElse func(err error) any) Result[any]
	// Or returns this Result if it is Ok, otherwise returns the provided alternative Result.
	//
	// Example:
	//
	//  result := Ok("value")
	//  other := Ok("other")
	//  result.Or(other).Unwrap() // "value"
	//
	//  result := Err[string](errors.New("error"))
	//  other := Ok("other")
	//  result.Or(other).Unwrap() // "other"
	Or(res Result[T]) Result[T]
	// OrElse returns this Result if it is Ok, otherwise calls fn with the error to produce an alternative Result.
	//
	// Example:
	//
	//  result := Ok("value")
	//  result.OrElse(func(e error) Result[string] {
	//      return Ok("fallback")
	//  }).Unwrap() // "value"
	//
	//  result := Err[string](errors.New("error"))
	//  result.OrElse(func(e error) Result[string] {
	//      return Ok("fallback")
	//  }).Unwrap() // "fallback"
	OrElse(fn func(err error) Result[T]) Result[T]
}

type ResultToOption[T any] interface {
	// Ok returns Some(value) if the result is Ok, otherwise returns None.
	//
	// Example:
	//
	//  result := Ok("value")
	//  opt := result.Ok() // Some("value")
	//
	//  result := Err[string](errors.New("error"))
	//  opt := result.Ok() // None
	Ok() Option[T]

	// Err returns Some(error) if the result is Err, otherwise returns None.
	//
	// Example:
	//
	//  result := Err[string](errors.New("error_message"))
	//  opt := result.Err() // Some(error)
	//
	//  result := Ok("value")
	//  opt := result.Err() // None
	Err() Option[error]
}
