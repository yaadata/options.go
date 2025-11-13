# optionsgo

A Go implementation of Rust's `Option<T>` and `Result<T, E>` types.

## Overview

This library brings Rust's ergonomic error handling patterns to Go through two
core types:

- **`Option[T]`** - Represents an optional value that may or may not exist
- **`Result[T]`** - Represents the outcome of an operation that can either
  succeed or fail

## Dependencies

The following tools are used in this project

- [mise](https://github.com/jdx/mise)

  Version management of Go for local development.

- [just](https://github.com/casey/just)

  Command runner. Alternative to Make.

## Installation

```bash
go get github.com/yaadata/optionsgo
```

## Quick Start

```go
import (
    "errors"
    "fmt"

    "github.com/yaadata/optionsgo"
)

// Using Option
func findUser(id int) optionsgo.Option[string] {
    if id < 0 {
        return optionsgo.None[string]()
    }
    return optionsgo.Some("Alice")
}

result := findUser(123)
if result.IsSome() {
    fmt.Println(result.Unwrap()) // "Alice"
}

// Using Result
func divideNumbers(a, b int) optionsgo.Result[int] {
    if b == 0 {
        return optionsgo.Err[int](errors.New("division by zero"))
    }
    return optionsgo.Ok(a / b)
}

result := divideNumbers(10, 2)
value := result.UnwrapOr(0) // 5
```

## API Reference

### Option[T] Interface

The `Option[T]` type represents an optional value. Every Option is either `Some`
(contains a value) or `None` (does not contain a value).

| Method                                  | Description                                         |
| --------------------------------------- | --------------------------------------------------- |
| `IsSome() bool`                         | Returns true if the option contains a value         |
| `IsSomeAnd(pred Predicate[T]) bool`     | Returns true if Some and predicate matches          |
| `IsNone() bool`                         | Returns true if the option does not contain a value |
| `IsNoneOr(pred Predicate[T]) bool`      | Returns true if None or predicate matches           |
| `Unwrap() T`                            | Returns the contained value (panics if None)        |
| `Expect(msg string) T`                  | Returns the value with custom panic message if None |
| `UnwrapOrElse(fn func() T) T`           | Returns the value or computes a default             |
| `UnwrapOrDefault() T`                   | Returns the value or the zero value of type T       |
| `OkOr(err error) Result[T]`             | Converts Some(v) to Ok(v), None to Err(err)         |
| `OkOrElse(fn func() error) Result[T]`   | Converts to Result with computed error              |
| `Or(optb Option[T]) Option[T]`          | Returns self if Some, otherwise returns optb        |
| `OrElse(fn func() Option[T]) Option[T]` | Returns self if Some, otherwise calls fn            |
| `Equal(Option[T]) bool`                 | Returns true if both options are equal              |

#### Constructor Functions

```go
Some[T](val T) Option[T]     // Creates an Option containing a value
None[T]() Option[T]           // Creates an Option with no value
```

### Result[T] Interface

The `Result[T]` type represents an operation that can either succeed with a
value or fail with an error.

| Method                                   | Description                                  |
| ---------------------------------------- | -------------------------------------------- |
| `IsOk() bool`                            | Returns true if the result is Ok             |
| `IsOkAnd(pred Predicate[T]) bool`        | Returns true if Ok and predicate matches     |
| `IsError() bool`                         | Returns true if the result is Err            |
| `IsErrorAnd(pred Predicate[error]) bool` | Returns true if Err and predicate matches    |
| `Ok() Option[T]`                         | Returns Some(value) if Ok, otherwise None    |
| `Err() Option[error]`                    | Returns Some(error) if Err, otherwise None   |
| `Unwrap() T`                             | Returns the Ok value (panics if Err)         |
| `UnwrapErr() error`                      | Returns the Err value (panics if Ok)         |
| `UnwrapOr(value T) T`                    | Returns the Ok value or provided default     |
| `UnwrapOrElse(fn func() T) T`            | Returns the Ok value or computes default     |
| `UnwrapOrDefault() T`                    | Returns the Ok value or zero value of type T |

#### Constructor Functions

```go
Ok[T](value T) Result[T]                           // Creates a Result containing a success value
Err[T](err error) Result[T]                        // Creates a Result containing an error
ResultFromReturn[T](value T, err error) Result[T]  // Converts Go's (T, error) pattern to Result
```

## Usage Examples

### Working with Option[T]

#### Creating and Checking Options

```go
// Create Some option
opt := optionsgo.Some("hello")
opt.IsSome()  // true
opt.IsNone()  // false

// Create None option
empty := optionsgo.None[string]()
empty.IsSome()  // false
empty.IsNone()  // true
```

#### Conditional Checks with Predicates

```go
opt := optionsgo.Some("SOME")

// Check if Some and predicate passes
opt.IsSomeAnd(func(s string) bool {
    return len(s) == 4
}) // true

opt.IsSomeAnd(func(s string) bool {
    return len(s) == 3
}) // false

// Check if None or predicate passes
opt.IsNoneOr(func(s string) bool {
    return len(s) == 4
}) // true (predicate passes)

optionsgo.None[string]().IsNoneOr(func(s string) bool {
    return false
}) // true (is None)
```

#### More Examples ?

See the respective test files

## Roadmap

The following are some possible candidates to expand the interface surface of
Option and Result

**For Option[T]:**

- Transformation:
  - [ ] `Xor`
  - [ ] `Flatten`
  - [ ] `Replace`
- Inspection:
  - [ ] `Inspect`
- Combination:
  - [ ] `Zip`
  - [ ] `ZipWith`
  - [ ] `Unzip`
- Conversion:
  - [ ] `Transpose`

**For Result[T]:**

- Transformation:
  - [ ] `And`
  - [ ] `AndThen`
  - [ ] `Or`
  - [ ] `OrElse`
  - [ ] `Flatten`
- Inspection:
  - [ ] `Inspect`
  - [ ] `InspectErr`
- Extraction:
  - [ ] `Expect`
- Conversion:
  - [ ] `Transpose`

## Commands

### Test

Run the test suite:

```bash
just test
```

Each merge to main is validated on the last 2 major version of Go across
Windows, Mac and Linux.

### Docs

View dows

```bash
just doc
```

### Lint

Run the project linter

```bash
just lint
```

## Contributing

Contributions are welcome! If you'd like to implement any of the missing APIs or
have suggestions for Go-idiomatic alternatives, please open a discussion.

## References

- [Rust std::option::Option](https://doc.rust-lang.org/std/option/enum.Option.html)
- [Rust std::result::Result](https://doc.rust-lang.org/std/result/enum.Result.html)

## License

See LICENSE file for details.
