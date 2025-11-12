package extension_test

import (
	"strings"
	"testing"

	"github.com/shoenig/test/must"
	"github.com/yaadata/optionsgo/core"
	"github.com/yaadata/optionsgo/extension"
	"github.com/yaadata/optionsgo/internal"
)

func TestOptionAndThen(t *testing.T) {
	t.Run("Some transforms to new type", func(t *testing.T) {
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
	t.Run("Some maps to new type", func(t *testing.T) {
		// [A]rrange
		option := internal.Some(3)
		fn := func(value int) string {
			return strings.Repeat("A", value)
		}
		// [A]ct
		actual := extension.OptionMapOr(option, fn, "DEFAULT")
		// [A]ssert
		must.True(t, actual.IsSome())
		must.Eq(t, "AAA", actual.Unwrap())
	})

	t.Run("None returns Some of default", func(t *testing.T) {
		// [A]rrange
		option := internal.None[int]()
		fn := func(value int) string {
			return strings.Repeat("A", value)
		}
		expected := "DEFAULT"
		// [A]ct
		actual := extension.OptionMapOr(option, fn, expected)
		// [A]ssert
		must.True(t, actual.IsSome())
		must.Eq(t, expected, actual.Unwrap())
	})
}

func TestOptionMapOrElse(t *testing.T) {
	t.Run("Some maps to new type", func(t *testing.T) {
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
		must.True(t, actual.IsSome())
		must.Eq(t, "AAA", actual.Unwrap())
	})

	t.Run("None returns Some of default", func(t *testing.T) {
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
		must.True(t, actual.IsSome())
		must.Eq(t, expected, actual.Unwrap())
	})
}
