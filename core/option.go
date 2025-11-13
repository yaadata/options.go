package core

// Option is a Go implementation of Rust's Option<T> type.
// It represents an optional value: every Option is either Some and contains a value,
// or None and does not contain a value.
//
// Example:
//
//	// Create an Option with a value
//	opt := Some("hello")
//	if opt.IsSome() {
//	    fmt.Println(opt.Unwrap()) // prints: hello
//	}
//
//	// Create an Option without a value
//	opt := None[string]()
//	if opt.IsNone() {
//	    fmt.Println("no value")
//	}
type Option[T any] interface {
	OptionChain[T]
	OptionToResult[T]
	// Equal returns true if both options are equal.
	//
	// Equality rules:
	//  - Two None options are always equal
	//  - A Some and None option are never equal
	//  - Two Some options with comparable types are equal if their contained values are equal
	//  - For non-comparable types, only the exact same Option instance equals itself
	//  - Different Option instances with non-comparable types are never equal
	//
	// Example:
	//	opt1 := Some("hello")
	//	opt2 := Some("hello")
	//	opt1.Equal(opt2) // returns true (comparable type)
	Equal(Option[T]) bool

	// Expect returns the contained Some value.
	// Panics with the provided message if the value is None.
	//
	// Use this method when you want to extract the value and provide a custom
	// panic message if the value is None. This is useful for cases where you
	// expect a value and want to provide context about why the None case is invalid.
	//
	// Example:
	//	opt := Some("SOME")
	//	val := opt.Expect("oops") // returns "SOME", no panic
	//
	//	opt := None[string]()
	//	val := opt.Expect("test panic") // panics with message "test panic"
	Expect(msg string) T

	// IsNone returns true if the option does not contain a value (is None).
	//
	// Example:
	//	opt := None[string]()
	//	opt.IsNone() // returns true
	//
	//	opt := Some("SOME")
	//	opt.IsNone() // returns false
	IsNone() bool

	// IsNoneOr returns true if the option is None or the value matches the predicate.
	// Returns false only if the option is Some and the predicate returns false.
	//
	// Example:
	//	opt := None[string]()
	//	opt.IsNoneOr(func(v string) bool {
	//	    return false
	//	}) // returns true
	//
	//	opt := Some("SOME")
	//	opt.IsNoneOr(func(v string) bool {
	//	    return len(v) == 4
	//	}) // returns true
	//
	//	opt.IsNoneOr(func(v string) bool {
	//	    return len(v) == 3
	//	}) // returns false
	IsNoneOr(pred Predicate[T]) bool

	// IsSome returns true if the option contains a value (is Some).
	//
	// Example:
	//	opt := Some("SOME")
	//	opt.IsSome() // returns true
	//
	//	opt := None[string]()
	//	opt.IsSome() // returns false
	IsSome() bool

	// IsSomeAnd returns true if the option is Some and the value matches the predicate.
	// Returns false if the option is None or the predicate returns false.
	//
	// Example:
	//	opt := Some("SOME")
	//	opt.IsSomeAnd(func(v string) bool {
	//	    return len(v) == 4
	//	}) // returns true
	//
	//	opt.IsSomeAnd(func(v string) bool {
	//	    return len(v) == 3
	//	}) // returns false
	//
	//	opt := None[string]()
	//	opt.IsSomeAnd(func(v string) bool {
	//	    return true
	//	}) // returns false
	IsSomeAnd(pred Predicate[T]) bool

	// MapOr transforms the value or returns a default value, terminating the chain.
	// If the current chain represents Some, applies fn to the value and returns the transformed result.
	// If the current chain represents None, returns the provided default value 'or' without calling fn.
	//
	// This method terminates the chain and returns the final value directly (not an Option).
	//
	// Example:
	//  result := None[string]().
	//      Map(func(a string) any { return len(a) }).
	//      MapOr(func(value any) any {
	//          if val, ok := value.(int); ok {
	//              return fmt.Sprintf("VALUE=%d", val)
	//          }
	//          return ""
	//      }, "OTHER")
	//  result // "OTHER"
	MapOr(fn func(T) any, or any) any

	// MapOrElse transforms the value or computes a default value, terminating the chain.
	// If the current chain represents Some, applies fn to the value and returns the transformed result.
	// If the current chain represents None, calls orElse to compute a default value and returns that without calling fn.
	//
	// This method terminates the chain and returns the final value directly (not an Option).
	// Use this when computing the default value is expensive and should only happen when needed.
	//
	// Example:
	//  result := None[string]().
	//      Map(func(a string) any { return len(a) }).
	//      MapOrElse(func(value any) any {
	//          if val, ok := value.(int); ok {
	//              return fmt.Sprintf("VALUE=%d", val)
	//          }
	//          return ""
	//      }, func() any {
	//          return "OTHER"
	//      })
	//  result // "OTHER"
	MapOrElse(fn func(T) any, orElse func() any) any

	// Unwrap returns the contained Some value.
	// Panics if the value is None.
	//
	// Use this method when you are certain the option contains a value.
	// For more control over panic messages, use Expect instead.
	//
	// Example:
	//	opt := Some("SOME")
	//	val := opt.Unwrap() // returns "SOME"
	//
	//	opt := None[string]()
	//	val := opt.Unwrap() // panics
	Unwrap() T

	// UnwrapOrDefault returns the contained Some value or the zero value of type T.
	// If the option is Some, returns the contained value.
	// If the option is None, returns the zero value (e.g., "" for string, 0 for int).
	//
	// Example:
	//	opt := Some("SOME")
	//	val := opt.UnwrapOrDefault() // returns "SOME"
	//
	//	opt := None[string]()
	//	val := opt.UnwrapOrDefault() // returns ""
	UnwrapOrDefault() T

	// UnwrapOrElse returns the contained Some value or computes it from the provided function.
	// If the option is Some, returns the contained value.
	// If the option is None, calls the provided function and returns its result.
	//
	// Example:
	//	opt := Some("SOME")
	//	val := opt.UnwrapOrElse(func() string {
	//	    return "ELSE"
	//	}) // returns "SOME"
	//
	//	opt := None[string]()
	//	val := opt.UnwrapOrElse(func() string {
	//	    return "ELSE"
	//	}) // returns "ELSE"
	UnwrapOrElse(fn func() T) T
}

// OptionChain provides a chainable interface for transforming Option values.
// It allows multiple operations to be composed fluently before converting back to an Option.
type OptionChain[T any] interface {
	// And returns the other option if this option is None, otherwise returns this option.
	// This is the opposite of Or, which returns this option if Some, otherwise returns other.
	//
	// Example:
	//	opt := Some("SOME")
	//	other := Some("OTHER")
	//	result := opt.And(other)
	//	result.Equal(opt) // returns true, returns the original Some
	//
	//	opt := None[string]()
	//	other := Some("OTHER")
	//	result := opt.And(other)
	//	result.Equal(other) // returns true, returns the other option
	//
	//	opt := None[string]()
	//	other := None[string]()
	//	result := opt.And(other)
	//	result.IsNone() // returns true, returns the other None
	And(other Option[T]) Option[T]

	// AndThen chains another option-returning operation.
	// If the current chain represents Some, applies fn to the value and returns a new OptionChain with the result.
	// If the current chain represents None, returns an OptionChain representing None without calling fn.
	//
	// This is useful for chaining operations that might fail and return None.
	//
	// Example:
	//  result := Some(3).
	//      Map(func(v int) any { return v * 2 }).
	//      AndThen(func(v any) Option[any] {
	//          if num, ok := v.(int); ok && num > 5 {
	//              return Some(num)
	//          }
	//          return None[any]()
	//      }).
	//      Option()
	AndThen(fn func(T) Option[any]) Option[any]

	// Filter returns None if the option is None, otherwise calls the predicate
	// with the wrapped value and returns:
	//  - Some(t) if the predicate returns true
	//  - None if the predicate returns false
	//
	// Example:
	//	opt := Some(5)
	//	result := opt.Filter(func(value int) bool {
	//	    return value < 10
	//	})
	//	result.IsSome() // returns true
	//
	//	opt := Some(15)
	//	result := opt.Filter(func(value int) bool {
	//	    return value < 10
	//	})
	//	result.IsNone() // returns true
	//
	//	opt := None[int]()
	//	result := opt.Filter(func(value int) bool {
	//	    return value < 10
	//	})
	//	result.IsNone() // returns true
	Filter(pred Predicate[T]) Option[T]

	// Map transforms the value in the chain by applying a function.
	// If the current chain represents Some, applies fn to the value and returns a new OptionChain with the transformed value.
	// If the current chain represents None, returns an OptionChain representing None without calling fn.
	//
	// This enables fluent transformation of values while maintaining the Option context.
	//
	// Example:
	//  result := Some(15).
	//      Filter(func(val int) bool { return val > 10 }).
	//      Map(func(value int) any { return fmt.Sprintf("Value=%d", value) })

	//  result.Unwrap() // "Value=15"
	//
	//  none := None[int]().
	//      Map(func(value int) any { return fmt.Sprintf("Value=%d", value) })
	//  none.IsNone() // true
	Map(fn func(T) any) Option[any]

	// Or returns the option if it contains a value, otherwise returns optb.
	//
	// Example:
	//	opt := Some("SOME")
	//	other := Some("OTHER")
	//	result := opt.Or(other)
	//	result.Equal(opt) // returns true
	//
	//	opt := None[string]()
	//	other := Some("OptionB")
	//	result := opt.Or(other)
	//	result.Equal(other) // returns true
	Or(optb Option[T]) Option[T]

	// OrElse returns the option if it contains a value,
	// otherwise calls fn and returns the result.
	//
	// Example:
	//	opt := Some("SOME")
	//	result := opt.OrElse(func() Option[string] {
	//	    return Some("OTHER")
	//	})
	//	result.Equal(opt) // returns true
	//
	//	opt := None[string]()
	//	other := Some("OptionB")
	//	result := opt.OrElse(func() Option[string] {
	//	    return other
	//	})
	//	result.Equal(other) // returns true
	OrElse(fn func() Option[T]) Option[T]

	// XOr returns Some if exactly one of self or optb is Some, otherwise returns None.
	// This implements exclusive OR logic for Options.
	//
	// Truth table:
	//  - Some XOr Some = None
	//  - Some XOr None = Some(self)
	//  - None XOr Some = Some(optb)
	//  - None XOr None = None
	//
	// Example:
	//	opt := Some("ORIGINAL")
	//	other := None[string]()
	//	result := opt.XOr(other)
	//	result.Unwrap() // returns "ORIGINAL"
	//
	//	opt := Some("ORIGINAL")
	//	other := Some("OTHER")
	//	result := opt.XOr(other)
	//	result.IsNone() // returns true
	//
	//	opt := None[string]()
	//	other := Some("OTHER")
	//	result := opt.XOr(other)
	//	result.Unwrap() // returns "OTHER"
	//
	//	opt := None[string]()
	//	other := None[string]()
	//	result := opt.XOr(other)
	//	result.IsNone() // returns true
	XOr(optb Option[T]) Option[T]
}

type OptionToResult[T any] interface {
	// OkOr transforms the Option into a Result, mapping Some(v) to Ok(v) and None to Err(err).
	//
	// Example:
	//	opt := Some("SOME")
	//	result := opt.OkOr(errors.New("error"))
	//	result.IsOk() // returns true
	//	result.Unwrap() // returns "SOME"
	//
	//	opt := None[string]()
	//	result := opt.OkOr(errors.New("OkOr"))
	//	result.IsError() // returns true
	//	result.UnwrapErr() // returns the error
	OkOr(err error) Result[T]

	// OkOrElse transforms the Option into a Result, mapping Some(v) to Ok(v) and
	// None to Err(fn()), where fn is a function that produces an error.
	//
	// Example:
	//	opt := Some("SOME")
	//	result := opt.OkOrElse(func() error {
	//	    return errors.New("error")
	//	})
	//	result.IsOk() // returns true
	//
	//	opt := None[string]()
	//	result := opt.OkOrElse(func() error {
	//	    return errors.New("OkOrElse")
	//	})
	//	result.IsError() // returns true
	OkOrElse(fn func() error) Result[T]
}
