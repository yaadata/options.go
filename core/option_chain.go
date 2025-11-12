package core

type OptionChain[T any] interface {
	Option() Option[T]
	AndThen(fn func(T) Option[any]) OptionChain[any]
	Map(fn func(T) any) OptionChain[any]
	MapOr(fn func(T) any, or any) OptionChain[any]
	MapOrElse(fn func(T) any, orElse func() any) OptionChain[any]
}
