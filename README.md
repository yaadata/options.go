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
go get github.com/yaadata/optionsgo@v0.2
```

## Quick Start

```go
import (
    "errors"
    "fmt"

    "github.com/yaadata/optionsgo/core"
    . "github.com/yaadata/optionsgo"
)

// Using Option
func findUser(id int) core.Option[string] {
    if id < 0 {
        return None[string]()
    }
    return Some("Alice")
}

result := findUser(123)
if result.IsSome() {
    fmt.Println(result.Unwrap()) // "Alice"
}

// Using Result
func divideNumbers(a, b int) core.Result[int] {
    if b == 0 {
        return Err[int](errors.New("division by zero"))
    }
    return Ok(a / b)
}

result := divideNumbers(10, 2)
value := result.UnwrapOr(0) // 5
```

## API Reference

### Option[T] Interface

The interface methods can be found in [here](./core/option.go)

View the interface definition via html by running `just doc`

#### Constructor Functions

```go
Some[T](val T) Option[T]     // Creates an Option containing a value
None[T]() Option[T]           // Creates an Option with no value
```

### Result[T] Interface

The interface methods can be found in [here](./core/result.go)

View the interface definition via html by running `just doc`

#### Constructor Functions

```go
Ok[T](value T) Result[T]                           // Creates a Result containing a success value
Err[T](err error) Result[T]                        // Creates a Result containing an error
```

### Extension Package

The `extension` package provides additional utilities and advanced operations
for `Option[T]` and `Result[T]` types. A brief example is shown below. Read
ahead to see each functions capability.

```go
import ( 
    "github.com/yaadata/optionsgo/extension"
    . "github.com/yaadata/optionsgo"
)

// ...

opt := Some(5)
extension.OptionAndThen(opt(3), toOption)
```

| Function                                    | Description                                           | Example                                     |
| ------------------------------------------- | ----------------------------------------------------- | ------------------------------------------- |
| `ResultFromReturn[T](value T, err error)`   | Converts Go's `(value, error)` pattern to `Result[T]` | `extension.ResultFromReturn(fetchUser(id))` |
| `ResultFlatten[T](result)`                  | Removes one level of nesting from `Result[Result[T]]` | `ResultFlatten(Ok(Ok(42))) // Ok(42)`       |
| `ResultAnd[T, V](result, other)`            | Returns `other` if `result` is Ok, otherwise Err      | `ResultAnd(Ok(5), Ok("hi")) // Ok("hi")`    |
| `ResultAndThen[T, V](result, fn)`           | Chains operations that return Results                 | `ResultAndThen(Ok(5), toResult)`            |
| `ResultMap[T, V](result, fn)`               | Transforms an Ok value, preserves errors              | `ResultMap(Ok(3), toString) // Ok("3")`     |
| `ResultMapOr[T, V](result, fn, or)`         | Transforms Ok value or returns default                | `ResultMapOr(Err(e), fn, "default")`        |
| `ResultMapOrElse[T, V](result, fn, orElse)` | Transforms Ok or computes from error                  | `ResultMapOrElse(r, fn, errHandler)`        |
| `ResultTranspose[T](result)`                | Converts `Result[Option[T]]` to `Option[Result[T]]`   | `ResultTranspose(Ok(Some(42)))`             |
| `OptionFlatten[T](option)`                  | Removes one level of nesting from `Option[Option[T]]` | `OptionFlatten(Some(Some(42))) // Some(42)` |
| `OptionAndThen[T, V](option, fn)`           | Chains operations that return Options                 | `OptionAndThen(Some(3), toOption)`          |
| `OptionMap[T, V](option, fn)`               | Transforms a Some value, preserves None               | `OptionMap(Some(3), toString) // Some("3")` |
| `OptionMapOr[T, V](option, fn, or)`         | Transforms Some value or returns default              | `OptionMapOr(None(), fn, "default")`        |
| `OptionMapOrElse[T, V](option, fn, orElse)` | Transforms Some or computes alternative               | `OptionMapOrElse(opt, fn, compute)`         |
| `OptionTranspose[T](option)`                | Converts `Option[Result[T]]` to `Result[Option[T]]`   | `OptionTranspose(Some(Ok(42)))`             |
| `MustCast[T](original any)`                 | Casts value to type T, panics on failure              | `MustCast[int](value) // 42 or panic`       |
| `CastOrZero[V](original any)`               | Casts value to type V, returns zero value on failure  | `CastOrZero[int]("text") // 0`              |

## Usage Examples

### Working with Option[T]

#### Creating and Checking Options

```go
import (
    . "github.com/yaadata/optionsgo"
)
// Create Some option
opt := Some("hello")
opt.IsSome()  // true
opt.IsNone()  // false

// Create None option
empty := None[string]()
empty.IsSome()  // false
empty.IsNone()  // true
```

#### Conditional Checks with Predicates

```go
import (
    . "github.com/yaadata/optionsgo"
)

//... 

opt := Some("SOME")

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

None[string]().IsNoneOr(func(s string) bool {
    return false
}) // true (is None)
```

#### More Examples ?

See the respective test files

## Commands

### Test

Run the test suite:

```bash
just test
```

Each merge to main is validated on the last 2 major version of Go across
Windows, Mac and Linux.

### Docs

View docs

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
