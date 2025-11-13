package optionsgo_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/shoenig/test/must"

	. "github.com/yaadata/optionsgo"
	"github.com/yaadata/optionsgo/extension"
)

func TestResult_Error(t *testing.T) {
	t.Parallel()
	t.Run("Expect panics with the expected message", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		expected := "TEST"
		defer func() {
			if rec := recover(); rec != nil {
				r := extension.MustCast[string](rec)
				must.Eq(t, expected, r)
			} else {
				t.Fail()
			}
		}()
		result := Err[string](errors.New("err"))
		// [A]ct
		result.Expect(expected)
	})

	t.Run("ExpectErr does not panic", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		expected := "TEST"
		result := Err[string](errors.New("err"))
		// [A]ct
		var actual error
		fn := func() {
			actual = result.ExpectErr(expected)
		}
		// [A]ssert
		must.NotPanic(t, fn)
		must.Eq(t, "err", actual.Error())
	})

	t.Run("Inspect runs fn", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		expected := "EXPECTED"
		result := Ok(expected)
		// [A]ct
		var actual bool
		fn := func(_ string) {
			actual = true
		}
		result.Inspect(fn)
		// [A]ssert
		must.True(t, actual)
	})

	t.Run("InspectErr does not run fn", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		result := Ok("msg")
		var actual bool
		// [A]ct
		fn := func(_ error) {
			actual = true
		}
		result.InspectErr(fn)
		// [A]ssert
		must.False(t, actual)
	})

	t.Run("IsOk returns false", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		result := Err[string](errors.New("err"))
		// [A]ct
		actual := result.IsOk()
		// [A]ssert
		must.False(t, actual)
	})

	t.Run("IsOkAnd returns false", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		result := Err[string](errors.New("err"))
		pred := func(_ string) bool {
			return true
		}
		// [A]ct
		actual := result.IsOkAnd(pred)
		// [A]ssert
		must.False(t, actual)
	})

	t.Run("IsError returns true", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		result := Err[string](errors.New("err"))
		// [A]ct
		actual := result.IsError()
		// [A]ssert
		must.True(t, actual)
	})

	t.Run("IsErrorAnd returns true", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		msg := "error_message"
		result := Err[string](errors.New(msg))
		pred := func(err error) bool {
			return err.Error() == msg
		}
		// [A]ct
		actual := result.IsErrorAnd(pred)
		// [A]ssert
		must.True(t, actual)
	})

	t.Run("IsErrorAnd returns false", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		msg := "error_message"
		result := Err[string](errors.New(msg))
		pred := func(err error) bool {
			return err.Error() != msg
		}
		// [A]ct
		actual := result.IsErrorAnd(pred)
		// [A]ssert
		must.False(t, actual)
	})

	t.Run("MapErr returns transformed error", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		result := Err[string](errors.New("A"))
		// [A]ct
		actual := result.MapErr(func(err error) error {
			return fmt.Errorf("%s - B", err.Error())
		})
		// [A]ssert
		must.True(t, actual.IsError())
		must.Eq(t, "A - B", actual.UnwrapErr().Error())
	})

	t.Run("MapErr returns transformed error", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		result := Ok(15)
		// [A]ct
		actual := result.MapErr(func(err error) error {
			return fmt.Errorf("%s - B", err.Error())
		})
		// [A]ssert
		must.True(t, actual.IsOk())
		must.Eq(t, 15, result.Unwrap())
		must.False(t, actual.IsError())
	})

	t.Run("Unwrap panics", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		msg := "error_message"
		result := Err[string](errors.New(msg))
		// [A]ct
		fn := func() {
			result.Unwrap()
		}
		// [A]ssert
		must.Panic(t, fn)
	})

	t.Run("UnwrapErr returns back the err", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		expected := errors.New("error_message")
		result := Err[string](expected)
		// [A]ct
		actual := result.UnwrapErr()
		// [A]ssert
		must.Eq(t, expected, actual)
	})

	t.Run("UnwrapOr returns Other option", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		result := Err[string](errors.New("error_message"))
		expected := "EXPECTED"
		// [A]ct
		actual := result.UnwrapOr(expected)
		// [A]ssert
		must.Eq(t, expected, actual)
	})

	t.Run("UnwrapOrElse returns Other option", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		result := Err[string](errors.New("error_message"))
		expected := "EXPECTED"
		// [A]ct
		actual := result.UnwrapOrElse(func() string {
			return expected
		})
		// [A]ssert
		must.Eq(t, expected, actual)
	})

	t.Run("UnwrapOrDefault returns Default", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		result := Err[string](errors.New("error_message"))
		// [A]ct
		actual := result.UnwrapOrDefault()
		// [A]ssert
		must.Eq(t, "", actual)
	})

	t.Run("Ok returns None", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		result := Err[string](errors.New("error_message"))
		// [A]ct
		actual := result.Ok()
		// [A]ssert
		must.True(t, actual.IsNone())
	})

	t.Run("Err returns Some", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		expected := errors.New("error_message")
		result := Err[string](expected)
		// [A]ct
		actual := result.Err()
		// [A]ssert
		must.True(t, actual.IsSome())
		must.Eq(t, expected, actual.Unwrap())
	})
}

func TestResult_Value(t *testing.T) {
	t.Parallel()
	t.Run("Expect returns the inner value", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		expected := 13
		result := Ok(expected)
		// [A]ct
		actual := result.Expect("ERROR_MESSAGE")
		// [A]ssert
		must.Eq(t, expected, actual)
	})

	t.Run("ExpectErr panics with the expected message", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		expected := "TEST"
		defer func() {
			if rec := recover(); rec != nil {
				r := extension.MustCast[string](rec)
				must.Eq(t, expected, r)
			} else {
				t.Fail()
			}
		}()
		result := Ok(13)
		// [A]ct
		_ = result.ExpectErr(expected)
	})

	t.Run("Inspect runs fn", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		result := Ok("msg")
		// [A]ct
		var actual bool
		fn := func(value string) {
			actual = true
		}
		result.Inspect(fn)
		// [A]ssert
		must.True(t, actual)
	})

	t.Run("InspectErr does not run fn", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		result := Ok("msg")
		var actual bool
		// [A]ct
		fn := func(_ error) {
			actual = true
		}
		result.InspectErr(fn)
		// [A]ssert
		must.False(t, actual)
	})

	t.Run("IsOk returns true", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		result := Ok("value")
		// [A]ct
		actual := result.IsOk()
		// [A]ssert
		must.True(t, actual)
	})

	t.Run("IsOkAnd returns false", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		result := Ok("value")
		pred := func(_ string) bool {
			return false
		}
		// [A]ct
		actual := result.IsOkAnd(pred)
		// [A]ssert
		must.False(t, actual)
	})

	t.Run("IsOkAnd returns true", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		result := Ok("value")
		pred := func(_ string) bool {
			return true
		}
		// [A]ct
		actual := result.IsOkAnd(pred)
		// [A]ssert
		must.True(t, actual)
	})

	t.Run("IsError returns false", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		result := Ok("value")
		// [A]ct
		actual := result.IsError()
		// [A]ssert
		must.False(t, actual)
	})

	t.Run("IsErrorAnd returns false", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		result := Ok("value")
		pred := func(_ error) bool {
			return true
		}
		// [A]ct
		actual := result.IsErrorAnd(pred)
		// [A]ssert
		must.False(t, actual)
	})

	t.Run("Unwrap does not panic", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		expected := "EXPECTED"
		result := Ok(expected)
		// [A]ct
		var actual string
		fn := func() {
			actual = result.Unwrap()
		}
		// [A]ssert
		must.NotPanic(t, fn)
		must.Eq(t, expected, actual)
	})

	t.Run("UnwrapErr panics", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		result := Ok("value")
		// [A]ct
		fn := func() {
			_ = result.UnwrapErr()
		}
		// [A]ssert
		must.Panic(t, fn)
	})

	t.Run("UnwrapOr returns original value", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		expected := "EXPECTED"
		result := Ok(expected)
		// [A]ct
		actual := result.UnwrapOr("OTHER_OPTION")
		// [A]ssert
		must.Eq(t, expected, actual)
	})

	t.Run("UnwrapOrElse returns original value", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		expected := "EXPECTED"
		result := Ok(expected)
		// [A]ct
		actual := result.UnwrapOrElse(func() string {
			return "OTHER_OPTION"
		})
		// [A]ssert
		must.Eq(t, expected, actual)
	})

	t.Run("UnwrapOrDefault returns original value", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		expected := "EXPECTED"
		result := Ok(expected)
		// [A]ct
		actual := result.UnwrapOrDefault()
		// [A]ssert
		must.Eq(t, expected, actual)
	})

	t.Run("Ok returns Some", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		expected := "EXPECTED"
		result := Ok(expected)
		// [A]ct
		actual := result.Ok()
		// [A]ssert
		must.True(t, actual.IsSome())
		must.Eq(t, expected, actual.Unwrap())
	})

	t.Run("Err returns Some", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		result := Ok("EXPECTED")
		// [A]ct
		actual := result.Err()
		// [A]ssert
		must.True(t, actual.IsNone())
	})
}

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
		actual := ResultFromReturn(fn())
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
		actual := ResultFromReturn(fn())
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
		actual := ResultFromReturn(fn())
		// [A]ssert
		must.True(t, actual.IsOk())
		actualValue := actual.Unwrap()
		must.Nil(t, actualValue)
	})
}

func TestResultChaining(t *testing.T) {
	t.Parallel()
	t.Run("Chaining Ok leads to Result", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		result := Ok("parallel")
		// [A]ct
		actual := result.
			Map(func(value string) any {
				return len(value)
			}).
			Map(func(value any) any {
				return extension.MustCast[int](value) * 10
			})
		// [A]ssert
		must.True(t, actual.IsOk())
		must.Eq(t, 80, actual.Ok().Unwrap())
	})

	t.Run("Chaining Err leads to Err(_)", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		result := Err[string](errors.New("parallel"))
		// [A]ct
		actual :=
			result.
				Map(func(value string) any {
					return len(value)
				}).
				Map(func(value any) any {
					return extension.MustCast[int](value) * 10
				})
		// [A]ssert
		must.True(t, actual.IsError())
	})
}
