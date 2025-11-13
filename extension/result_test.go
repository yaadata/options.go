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

func TestResultMapFromReturn(t *testing.T) {
	t.Parallel()
	type Case struct {
		val string
	}
	t.Run("nil return with error", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		expected := errors.New("case a")
		fn := func() (*Case, error) {
			return nil, expected
		}
		// [A]ct
		actual := extension.ResultFromReturn(fn())
		// [A]ssert
		must.True(t, actual.IsError())
		must.Eq(t, expected, actual.UnwrapErr())
	})

	t.Run("value return with no error", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		expected := &Case{
			val: "EXPECTED",
		}
		fn := func() (*Case, error) {
			return expected, nil
		}
		// [A]ct
		actual := extension.ResultFromReturn(fn())
		// [A]ssert
		must.True(t, actual.IsOk())
		actualValue := actual.Unwrap()
		must.Eq(t, expected, actualValue)
	})

	t.Run("nil value and err", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		fn := func() (*Case, error) {
			return nil, nil
		}
		// [A]ct
		actual := extension.ResultFromReturn(fn())
		// [A]ssert
		must.True(t, actual.IsOk())
		actualValue := actual.Unwrap()
		must.Nil(t, actualValue)
	})
}

func TestResultMap(t *testing.T) {
	t.Parallel()
	t.Run("Original result is Ok", func(t *testing.T) {
		t.Parallel()
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
		t.Parallel()
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
	t.Parallel()
	t.Run("Original result is Ok", func(t *testing.T) {
		t.Parallel()
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
		t.Parallel()
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
	t.Parallel()
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
		t.Parallel()
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

func TestResultAnd(t *testing.T) {
	t.Parallel()
	t.Run("Ok returns other", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		result := internal.Ok(5)
		other := internal.Ok("OTHER")
		// [A]ct
		actual := extension.ResultAnd(result, other)
		// [A]ssert
		must.Eq(t, "OTHER", actual.Unwrap())
	})

	t.Run("Err returns Error", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		result := internal.Err[int](errors.New("ERROR"))
		other := internal.Ok("OTHER")
		// [A]ct
		actual := extension.ResultAnd(result, other)
		// [A]ssert
		must.Eq(t, "ERROR", actual.UnwrapErr().Error())
	})
}

func TestResultAndThen(t *testing.T) {
	t.Parallel()
	t.Run("Ok returns other", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		result := internal.Ok(5)
		// [A]ct
		actual := extension.ResultAndThen(result, func(resultValue int) core.Result[string] {
			return internal.Ok(strings.Repeat("A", resultValue))
		})
		// [A]ssert
		must.Eq(t, "AAAAA", actual.Unwrap())
	})

	t.Run("Err returns Error", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		result := internal.Err[int](errors.New("ERROR"))
		// [A]ct
		actual := extension.ResultAndThen(result, func(resultValue int) core.Result[string] {
			return internal.Ok(strings.Repeat("A", resultValue))
		})
		// [A]ssert
		must.Eq(t, "ERROR", actual.UnwrapErr().Error())
	})
}

func TestResultTranspose(t *testing.T) {
	t.Parallel()
	t.Run("Result Ok is Some(_)", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		result := internal.Ok(internal.Some(13))
		// [A]ct
		actual := extension.ResultTranspose(result)
		// [A]ssert
		must.True(t, actual.IsSome())
		must.Eq(t, 13, actual.Unwrap().Unwrap())
	})

	t.Run("Result Ok(None) is None", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		result := internal.Ok(internal.None[int]())
		// [A]ct
		actual := extension.ResultTranspose(result)
		// [A]ssert
		must.True(t, actual.IsNone())
	})

	t.Run("Result Ok(None) is None", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		result := internal.Ok(internal.None[int]())
		// [A]ct
		actual := extension.ResultTranspose(result)
		// [A]ssert
		must.True(t, actual.IsNone())
	})

	t.Run("Result Err() is Some(Err(_))", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		expected := errors.New("msg")
		result := internal.Err[core.Option[int]](expected)
		// [A]ct
		actual := extension.ResultTranspose(result)
		// [A]ssert
		must.True(t, actual.IsSome())
		must.Eq(t, expected.Error(), actual.Unwrap().UnwrapErr().Error())
	})
}

func TestResultFlatten(t *testing.T) {
	t.Parallel()
	t.Run("Ok returns in Ok Variant", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		result := internal.Ok(internal.Ok(5))
		// [A]ct
		actual := extension.ResultFlatten(result)
		// [A]ssert
		must.True(t, actual.IsOk())
		must.Eq(t, 5, actual.Unwrap())
	})

	t.Run("Ok returns in Ok Variant Only One Level Deep", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		result := internal.Ok(internal.Ok(internal.Ok(5)))
		// [A]ct
		actual := extension.ResultFlatten(extension.ResultFlatten(result))
		// [A]ssert
		must.True(t, actual.IsOk())
		must.Eq(t, 5, actual.Unwrap())
	})

	t.Run("Err returns in Err Variant", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		err := errors.New("ERROR")
		option := internal.Err[core.Result[int]](errors.New("ERROR"))
		// [A]ct
		actual := extension.ResultFlatten(option)
		// [A]ssert
		must.True(t, actual.IsError())
		must.Eq(t, err, actual.UnwrapErr())
	})
}
