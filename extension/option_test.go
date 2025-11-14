package extension_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/shoenig/test/must"
	"github.com/yaadata/optionsgo/core"
	"github.com/yaadata/optionsgo/extension"
	"github.com/yaadata/optionsgo/internal"
)

func TestOptionAndThen(t *testing.T) {
	t.Parallel()
	t.Run("Some transforms to new type", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		option := internal.Some(3)
		fn := func(value int) core.Option[string] {
			return internal.Some(strings.Repeat("A", value))
		}
		// [A]ct
		actual := extension.OptionAndThen(option, fn)
		// [A]ssert
		must.True(t, actual.IsSome())
		must.Eq(t, "AAA", actual.Unwrap())
	})

	t.Run("None returns None", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		option := internal.None[int]()
		fn := func(value int) core.Option[string] {
			return internal.Some(strings.Repeat("A", value))
		}
		// [A]ct
		actual := extension.OptionAndThen(option, fn)
		// [A]ssert
		must.True(t, actual.IsNone())
	})
}

func TestOptionMap(t *testing.T) {
	t.Run("Some maps to new type", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		option := internal.Some(3)
		fn := func(value int) string {
			return strings.Repeat("A", value)
		}
		// [A]ct
		actual := extension.OptionMap(option, fn)
		// [A]ssert
		must.True(t, actual.IsSome())
		must.Eq(t, "AAA", actual.Unwrap())
	})

	t.Run("None returns None", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		option := internal.None[int]()
		fn := func(value int) string {
			return strings.Repeat("A", value)
		}
		// [A]ct
		actual := extension.OptionMap(option, fn)
		// [A]ssert
		must.True(t, actual.IsNone())
	})
}

func TestOptionMapOr(t *testing.T) {
	t.Parallel()
	t.Run("Some maps to new type", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		option := internal.Some(3)
		fn := func(value int) string {
			return strings.Repeat("A", value)
		}
		// [A]ct
		actual := extension.OptionMapOr(option, fn, "DEFAULT")
		// [A]ssert
		must.Eq(t, "AAA", actual)
	})

	t.Run("None returns Some of default", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		option := internal.None[int]()
		fn := func(value int) string {
			return strings.Repeat("A", value)
		}
		expected := "DEFAULT"
		// [A]ct
		actual := extension.OptionMapOr(option, fn, expected)
		// [A]ssert
		must.Eq(t, expected, actual)
	})
}

func TestOptionMapOrElse(t *testing.T) {
	t.Parallel()
	t.Run("Some maps to new type", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		option := internal.Some(3)
		fn := func(value int) string {
			return strings.Repeat("A", value)
		}
		// [A]ct
		actual := extension.OptionMapOrElse(option, fn, func() string {
			return "DEFAULT"
		})
		// [A]ssert
		must.Eq(t, "AAA", actual)
	})

	t.Run("None returns Some of default", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		option := internal.None[int]()
		fn := func(value int) string {
			return strings.Repeat("A", value)
		}
		expected := "DEFAULT"
		// [A]ct
		actual := extension.OptionMapOrElse(option, fn, func() string {
			return expected
		})
		// [A]ssert
		must.Eq(t, expected, actual)
	})
}

func TestOptionFlatten(t *testing.T) {
	t.Parallel()
	t.Run("Some returns in Some Variant", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		option := internal.Some(internal.Some(5))
		// [A]ct
		actual := extension.OptionFlatten(option)
		// [A]ssert
		must.True(t, actual.IsSome())
		must.Eq(t, 5, actual.Unwrap())
	})

	t.Run("Some returns in Some Variant Only One Level Deep", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		option := internal.Some(internal.Some(internal.Some(5)))
		// [A]ct
		actual := extension.OptionFlatten(extension.OptionFlatten(option))
		// [A]ssert
		must.True(t, actual.IsSome())
		must.Eq(t, 5, actual.Unwrap())
	})

	t.Run("None returns in None Variant", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		option := internal.Some(internal.None[int]())
		// [A]ct
		actual := extension.OptionFlatten(option)
		// [A]ssert
		must.True(t, actual.IsNone())
	})
}

func TestOptionTranspose(t *testing.T) {
	t.Parallel()
	t.Run("Option is None => Ok(None)", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		option := internal.None[core.Result[int]]()
		// [A]ct
		actual := extension.OptionTranspose(option)
		// [A]ssert
		must.True(t, actual.IsOk())
	})

	t.Run("Option is Some Value => Ok(Value)", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		option := internal.Some(internal.Ok(5))
		// [A]ct
		actual := extension.OptionTranspose(option)
		// [A]ssert
		must.True(t, actual.IsOk())
		must.True(t, actual.Unwrap().IsSome())
		must.Eq(t, 5, actual.Unwrap().Unwrap())
	})

	t.Run("Option is Some Err => Ok(Err)", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		option := internal.Some(internal.Err[int](errors.New("ERROR")))
		// [A]ct
		actual := extension.OptionTranspose(option)
		// [A]ssert
		must.True(t, actual.IsError())
		must.Eq(t, "ERROR", actual.UnwrapErr().Error())
	})

	t.Run("Transpose back and forth works", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		option := internal.Some(internal.Ok(5))
		result := extension.OptionTranspose(option)
		// [A]ct
		actual := extension.ResultTranspose(result)
		// [A]ssert
		must.True(t, option.Equal(actual))
	})
}

func TestOptionFromPointer(t *testing.T) {
	t.Parallel()
	t.Run("Nil pointer returns None", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		var ptr *string
		// [A]ct
		actual := extension.OptionFromPointer(ptr)
		// [A]ssert
		must.True(t, actual.IsNone())
	})

	t.Run("Nil pointer returns None", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		value := "value"
		ptr := &value
		// [A]ct
		actual := extension.OptionFromPointer(ptr)
		// [A]ssert
		must.True(t, actual.IsSome())
		must.Eq(t, value, actual.Unwrap())
	})
}
