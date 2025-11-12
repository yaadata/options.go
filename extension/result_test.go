package extension_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/shoenig/test/must"
	"github.com/yaadata/optionsgo/extension"
	"github.com/yaadata/optionsgo/internal"
)

func TestResultMap(t *testing.T) {
	t.Run("Original result is Ok", func(t *testing.T) {
		// [A]rrange
		result := internal.Ok(3)
		fn := func(value int) string {
			return strings.Repeat("A", value)
		}
		// [A]ct
		actual := extension.ResultMap(result, fn)
		// [A]ssert
		must.True(t, actual.IsOk())
		must.Eq(t, "AAA", actual.Unwrap())
	})

	t.Run("Original result is Err", func(t *testing.T) {
		// [A]rrange
		result := internal.Err[int](errors.New("error"))
		fn := func(value int) string {
			return strings.Repeat("A", value)
		}
		// [A]ct
		actual := extension.ResultMap(result, fn)
		// [A]ssert
		must.True(t, actual.IsError())
	})
}

func TestResultMapOr(t *testing.T) {
	t.Run("Original result is Ok", func(t *testing.T) {
		// [A]rrange
		result := internal.Ok(3)
		fn := func(value int) string {
			return strings.Repeat("A", value)
		}
		// [A]ct
		actual := extension.ResultMapOr(result, fn, "DEFAULT")
		// [A]ssert
		must.True(t, actual.IsOk())
		must.Eq(t, "AAA", actual.Unwrap())
	})

	t.Run("Original result is Err", func(t *testing.T) {
		// [A]rrange
		result := internal.Err[int](errors.New("error"))
		fn := func(value int) string {
			return strings.Repeat("A", value)
		}
		expected := "DEFAULT"
		// [A]ct
		actual := extension.ResultMapOr(result, fn, expected)
		// [A]ssert
		must.True(t, actual.IsOk())
		must.Eq(t, expected, actual.Unwrap())
	})
}

func TestResultMapOrElse(t *testing.T) {
	t.Run("Original result is Ok", func(t *testing.T) {
		// [A]rrange
		result := internal.Ok(3)
		fn := func(value int) string {
			return strings.Repeat("A", value)
		}
		orElse := func(_ error) string {
			return "OTHER"
		}
		// [A]ct
		actual := extension.ResultMapOrElse(result, fn, orElse)
		// [A]ssert
		must.True(t, actual.IsOk())
		must.Eq(t, "AAA", actual.Unwrap())
	})

	t.Run("Original result is Err", func(t *testing.T) {
		// [A]rrange
		result := internal.Err[int](errors.New("error"))
		expected := "EXPECTED"
		fn := func(value int) string {
			return strings.Repeat("A", value)
		}
		orElse := func(_ error) string {
			return expected
		}
		// [A]ct
		actual := extension.ResultMapOrElse(result, fn, orElse)
		// [A]ssert
		must.True(t, actual.IsOk())
		must.Eq(t, expected, actual.Unwrap())
	})
}
