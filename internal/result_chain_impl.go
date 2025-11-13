package internal

import "github.com/yaadata/optionsgo/core"

func ResultMap[T, V any](result core.Result[T], fn func(inner T) V) core.Result[V] {
	if result.IsOk() {
		return Ok(fn(result.Unwrap()))
	}
	return Err[V](result.UnwrapErr())
}

func ResultMapOr[T, V any](result core.Result[T], fn func(inner T) V, or V) core.Result[V] {
	if result.IsOk() {
		return Ok(fn(result.Unwrap()))
	}
	return Ok(or)
}

func ResultMapOrElse[T, V any](result core.Result[T], fn func(inner T) V, orElse func(error) V) core.Result[V] {
	if result.IsOk() {
		return Ok(fn(result.Unwrap()))
	}
	return Ok(orElse(result.UnwrapErr()))
}
