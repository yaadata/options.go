package optionsgo

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
}

// None creates an Option that contains no value.
//
// Use None when you want to represent the absence of a value//
// Example:
//
//	// Function that may not return a value
//	func findUser(id int) Option[User] {
//	    if id < 0 {
//	        return None[User]()
//	    }
//	    // ... find user logic
//	    return Some(user)
//	}
//
//	result := findUser(-1)
//	if result.IsNone() {
//	    fmt.Println("User not found")
//	}
func None[T any]() Option[T] {
	return &option[T]{value: nil}
}

// Some creates an Option that contains the provided value.
//
// Use Some when you have a valid value to wrap in an Option type.
// Example:
//
//	// Function that returns a value when found
//	func findUser(id int) Option[User] {
//	    if id < 0 {
//	        return None[User]()
//	    }
//	    // ... find user logic
//	    return Some(user)
//	}
//
//	result := findUser(123)
//	if result.IsSome() {
//	    user := result.Unwrap()
//	    fmt.Printf("Found user: %v\n", user)
//	}
func Some[T any](val T) Option[T] {
	return &option[T]{value: &val}
}
