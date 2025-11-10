package options

type Predicate[T any] func(val T) bool
