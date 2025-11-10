// Package options provides Rust-style Option and Result types for Go.
//
// This package replicates the semantics of Rust's std::option::Option and
// std::result::Result types
//
// # Options Type
//
// Options[T] represents an optional value: every Options is either Some containing
// a value of type T, or None indicating the absence of a value.
//
// # Result Type
//
// Result[T, E] represents a value that could be successful (Ok) or an error (Err).
// This is similar to Go's (value, error).
package options
