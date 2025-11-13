package optionsgo_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/shoenig/test/must"
	. "github.com/yaadata/optionsgo"
	"github.com/yaadata/optionsgo/core"
)

func TestOption_None(t *testing.T) {
	t.Parallel()
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

	t.Run("Equal is true between two None containers", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		optionA := None[string]()
		optionB := None[string]()
		// [A]ct
		actual := optionA.Equal(optionB)
		// [A]ssert
		must.True(t, actual)
	})

	t.Run("Equal is false between None and Some containers", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		optionA := None[string]()
		optionB := Some("other")
		// [A]ct
		actual := optionA.Equal(optionB)
		// [A]ssert
		must.False(t, actual)
	})

	t.Run("Expect should panic", func(t *testing.T) {
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

	t.Run("Unwrap should panic", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		defer func() {
			if r := recover(); r != nil {
				must.Eq(t, "failed to unwrap None value", r.(string))
			} else {
				t.Error("expected a panic but none occurred")
			}
		}()
		val := None[string]()
		// [A]ct
		val.Unwrap()
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
		opt := None[string]()
		expected := Some("OptionB")
		// [A]ct
		actual := opt.Or(expected)
		// [A]ssert
		must.True(t, actual.Equal(expected))
	})

	t.Run("OrElse returns other option", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		opt := None[string]()
		expected := Some("OptionB")
		// [A]ct
		result := opt.OrElse(func() core.Option[string] {
			return expected
		})
		// [A]ssert
		must.Eq(t, expected, result)
	})

	t.Run("Filter on None returns None", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		opt := None[int]()
		pred := func(value int) bool {
			return value < 10
		}
		// [A]ct
		actual := opt.Filter(pred)
		// [A]ssert
		must.True(t, actual.IsNone())
	})

	t.Run("And returns other option when called on None with None", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		opt := None[string]()
		other := None[string]()
		// [A]ct
		actual := opt.And(other)
		// [A]ssert
		must.True(t, actual.IsNone())
		must.True(t, actual.Equal(other))
	})

	t.Run("And returns other option when called on None with Some", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		opt := None[string]()
		expected := Some("other")
		// [A]ct
		actual := opt.And(expected)
		// [A]ssert
		must.True(t, actual.IsSome())
		must.True(t, actual.Equal(expected))
	})
}

func TestOption_Some(t *testing.T) {
	t.Parallel()
	t.Run("And returns current Some when called on None", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		opt := Some("current")
		other := None[string]()
		// [A]ct
		actual := opt.And(other)
		// [A]ssert
		must.True(t, opt.IsSome())
		must.True(t, actual.Equal(opt))
	})

	t.Run("And returns current Some when called on Some", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		opt := Some("current")
		other := Some("other")
		// [A]ct
		actual := opt.And(other)
		// [A]ssert
		must.True(t, opt.IsSome())
		must.True(t, actual.Equal(opt))
	})

	t.Run("Filter on Some where predicate is true", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		opt := Some(5)
		pred := func(value int) bool {
			return value < 10
		}
		// [A]ct
		actual := opt.Filter(pred)
		// [A]ssert
		must.True(t, actual.IsSome())
	})

	t.Run("Filter on Some where predicate is true", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		opt := Some(15)
		pred := func(value int) bool {
			return value < 10
		}
		// [A]ct
		actual := opt.Filter(pred)
		// [A]ssert
		must.True(t, actual.IsNone())
	})

	t.Run("IsSome is true", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		opt := Some("SOME")
		// [A]ct
		actual := opt.IsSome()
		// [A]ssert
		must.True(t, actual)
	})

	t.Run("IsSomeAnd with predicate leading to true", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		opt := Some("SOME")
		pred := func(value string) bool {
			return len(value) == 4
		}
		// [A]ct
		actual := opt.IsSomeAnd(pred)
		// [A]ssert
		must.True(t, actual)
	})

	t.Run("IsSomeAnd with predicate leading to false", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		opt := Some("SOME")
		pred := func(value string) bool {
			return len(value) == 3
		}
		// [A]ct
		actual := opt.IsSomeAnd(pred)
		// [A]ssert
		must.False(t, actual)
	})

	t.Run("IsNone is false", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		opt := Some("SOME")
		// [A]ct
		actual := opt.IsNone()
		// [A]ssert
		must.False(t, actual)
	})

	t.Run("IsNoneOr where the predicate leads to false", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		opt := Some("SOME")
		pred := func(value string) bool {
			return len(value) == 3
		}
		// [A]ct
		actual := opt.IsNoneOr(pred)
		// [A]ssert
		must.False(t, actual)
	})

	t.Run("IsNoneOr where the predicate leads to true", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		opt := Some("SOME")
		pred := func(value string) bool {
			return len(value) == 4
		}
		// [A]ct
		actual := opt.IsNoneOr(pred)
		// [A]ssert
		must.True(t, actual)
	})

	t.Run("Equal is true for two containers with the same value", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		optionA := Some(5)
		optionB := Some(5)
		// [A]ct
		actual := optionA.Equal(optionB)
		// [A]ssert
		must.True(t, actual)
	})

	t.Run("Equal is true for exact same container with Non-Comparable Type", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		value := map[string]string{
			"key": "value",
		}
		must.False(t, reflect.TypeOf(value).Comparable())
		opt := Some(value)
		// [A]ct
		actual := opt.Equal(opt)
		// [A]ssert
		must.True(t, actual)
	})

	t.Run("Equal is false for different containers with the same Non-Comparable Type", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		value := map[string]string{
			"key": "value",
		}
		must.False(t, reflect.TypeOf(value).Comparable())
		optionA := Some(value)
		optionB := Some(value)
		// [A]ct
		actual := optionA.Equal(optionB)
		// [A]ssert
		must.False(t, actual)
	})

	t.Run("Expect does not panic and returns the inner value", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		const EXPECTED = "SOME"
		opt := Some(EXPECTED)
		// [A]ct
		var actual string
		evaluate := func() {
			actual = opt.Expect("oops")
		}
		// [A]ssert
		must.NotPanic(t, evaluate)
		must.Eq(t, EXPECTED, actual)
	})

	t.Run("Unwrap does not panic and returns the inner value", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		const EXPECTED = "SOME"
		opt := Some(EXPECTED)
		// [A]ct
		var actual string
		evaluate := func() {
			actual = opt.Unwrap()
		}
		// [A]ssert
		must.NotPanic(t, evaluate)
		must.Eq(t, EXPECTED, actual)
	})

	t.Run("UnwrapOrElse returns some value", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		const EXPECTED = "SOME"
		opt := Some(EXPECTED)
		// [A]ct
		actual := opt.UnwrapOrElse(func() string {
			return "ELSE"
		})
		// [A]ssert
		must.Eq(t, EXPECTED, actual)
	})

	t.Run("UnwrapOrDefault returns some value", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		const EXPECTED = "SOME"
		opt := Some(EXPECTED)
		// [A]ct
		actual := opt.UnwrapOrDefault()
		// [A]ssert
		must.Eq(t, EXPECTED, actual)
	})

	t.Run("OkOr returns Result Ok", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		const EXPECTED = "SOME"
		opt := Some(EXPECTED)
		// [A]ct
		actual := opt.OkOr(errors.New("OkOr"))
		// [A]ssert
		must.True(t, actual.IsOk())
		must.Eq(t, EXPECTED, actual.Unwrap())
	})

	t.Run("OkOrElse returns Result Ok", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		const EXPECTED = "SOME"
		opt := Some(EXPECTED)
		// [A]ct
		actual := opt.OkOrElse(func() error {
			return errors.New("OkOrElse")
		})
		// [A]ssert
		must.True(t, actual.IsOk())
		must.Eq(t, EXPECTED, actual.Unwrap())
	})

	t.Run("OkOrElse returns Result Ok", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		const EXPECTED = "SOME"
		opt := Some(EXPECTED)
		// [A]ct
		actual := opt.OkOrElse(func() error {
			return errors.New("OkOrElse")
		})
		// [A]ssert
		must.True(t, actual.IsOk())
		must.Eq(t, EXPECTED, actual.Unwrap())
	})

	t.Run("Or returns Result original", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		const EXPECTED = "SOME"
		opt := Some(EXPECTED)
		other := Some("OTHER")
		// [A]ct
		actual := opt.Or(other)
		// [A]ssert
		must.False(t, actual.Equal(other))
		must.True(t, actual.Equal(opt))
		must.Eq(t, EXPECTED, actual.Unwrap())
	})

	t.Run("OrElse returns Result original", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		const EXPECTED = "SOME"
		opt := Some(EXPECTED)
		other := Some("OTHER")
		// [A]ct
		actual := opt.OrElse(func() core.Option[string] {
			return other
		})
		// [A]ssert
		must.False(t, actual.Equal(other))
		must.True(t, actual.Equal(opt))
		must.Eq(t, EXPECTED, actual.Unwrap())
	})
}
