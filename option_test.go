package options

import (
	"errors"
	"testing"

	"github.com/shoenig/test/must"
)

func TestOption_None(t *testing.T) {
	t.Run("IsSome is false", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		val := None[string]()
		// [A]ct
		actual := val.IsSome()
		// [A]ssert
		must.False(t, actual)
	})

	t.Run("IsSomeAnd is false", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		val := None[string]()
		pred := func(_ string) bool {
			return true
		}
		// [A]ct
		actual := val.IsSomeAnd(pred)
		// [A]ssert
		must.False(t, actual)
	})

	t.Run("IsNone is true", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		val := None[string]()
		// [A]ct
		actual := val.IsNone()
		// [A]ssert
		must.True(t, actual)
	})

	t.Run("IsNoneOr is true", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		val := None[string]()
		pred := func(_ string) bool {
			return false
		}
		// [A]ct
		actual := val.IsNoneOr(pred)
		// [A]ssert
		must.True(t, actual)
	})

	t.Run("Unwrap should panic", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		msg := "test panic"
		defer func() {
			if r := recover(); r != nil {
				must.Eq(t, msg, r.(string))
			} else {
				t.Error("expected a panic but none occurred")
			}
		}()
		val := None[string]()
		// [A]ct
		val.Expect(msg)
	})

	t.Run("UnwrapOrElse returns Else", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		expected := "ELSE"
		val := None[string]()
		// [A]ct
		actual := val.UnwrapOrElse(func() string {
			return expected
		})
		// [A]ssert
		must.Eq(t, expected, actual)
	})

	t.Run("UnwrapOrDefault returns Default", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		val := None[string]()
		// [A]ct
		actual := val.UnwrapOrDefault()
		// [A]ssert
		must.Eq(t, "", actual)
	})

	t.Run("OkOr returns Result Err", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		val := None[string]()
		originalErr := errors.New("OkOr")
		// [A]ct
		actual := val.OkOr(originalErr)
		// [A]ssert
		must.True(t, actual.IsError())
		err := actual.UnwrapErr()
		must.Eq(t, originalErr, err)
	})

	t.Run("OkOrElse returns Result Err", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		val := None[string]()
		originalErr := errors.New("OkOrElse")
		// [A]ct
		actual := val.OkOrElse(func() error {
			return originalErr
		})
		// [A]ssert
		must.True(t, actual.IsError())
		err := actual.UnwrapErr()
		must.Eq(t, originalErr, err)
	})

	t.Run("Or returns other option", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		optionA := None[string]()
		expected := Some("OptionB")
		// [A]ct
		result := optionA.Or(expected)
		// [A]ssert
		must.Eq(t, expected, result)
	})

	t.Run("OrElse returns other option", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		optionA := None[string]()
		expected := Some("OptionB")
		// [A]ct
		result := optionA.OrElse(func() Option[string] {
			return expected
		})
		// [A]ssert
		must.Eq(t, expected, result)
	})
}
