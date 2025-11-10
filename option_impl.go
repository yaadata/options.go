package options

type option[T any] struct {
	value *T
}

// interface guard
var _ Option[string] = (*option[string])(nil)

func (o *option[T]) IsSome() bool {
	return o.value != nil
}

func (o *option[T]) IsSomeAnd(pred Predicate[T]) bool {
	if o.value != nil {
		return pred(*o.value)
	}
	return false
}

func (o *option[T]) IsNone() bool {
	return o.value == nil
}

func (o *option[T]) IsNoneOr(pred Predicate[T]) bool {
	if o.value != nil {
		return pred(*o.value)
	}
	return true
}

func (o *option[T]) Expect(msg string) T {
	if o.value == nil {
		panic(msg)
	}
	return *o.value
}

func (o *option[T]) Unwrap() T {
	return o.Expect("failed to unwrap None value")
}

func (o *option[T]) UnwrapOrElse(fn func() T) T {
	if o.value == nil {
		return fn()
	}
	return *o.value
}

func (o *option[T]) UnwrapOrDefault() T {
	if o.value == nil {
		return *new(T)
	}
	return *o.value
}

func (o *option[T]) Filter(pred Predicate[T]) Option[T] {
	if o.value != nil && pred(*o.value) {
		return o
	}
	return None[T]()
}

func (o *option[T]) OkOr(err error) Result[T] {
	if o.IsSome() {
		return Ok(*o.value)
	}
	return Err[T](err)
}

func (o *option[T]) OkOrElse(fn func() error) Result[T] {
	if o.IsSome() {
		return Ok(*o.value)
	}
	return Err[T](fn())
}

func (o *option[T]) Or(optb Option[T]) Option[T] {
	if o.IsNone() {
		return optb
	}
	return o
}

func (o *option[T]) OrElse(fn func() Option[T]) Option[T] {
	if o.IsNone() {
		return fn()
	}
	return o
}
