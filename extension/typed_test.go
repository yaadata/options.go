package extension

import (
	"testing"

	"github.com/shoenig/test/must"
)

func TestMustCast(t *testing.T) {
	t.Parallel()
	t.Run("Cast int to int succeeds", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		var value any = 42
		// [A]ct
		actual := MustCast[int](value)
		// [A]ssert
		must.Eq(t, 42, actual)
	})

	t.Run("Cast string to int panics", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		var value any = "not an int"
		// [A]ct
		fn := func() {
			MustCast[int](value)
		}
		// [A]ssert
		must.Panic(t, fn)
	})
}

func TestCastOrZero(t *testing.T) {
	t.Parallel()
	t.Run("Cast int to int succeeds", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		var value any = 42
		// [A]ct
		actual := CastOrZero[int](value)
		// [A]ssert
		must.Eq(t, 42, actual)
	})

	t.Run("Cast string to int returns zero", func(t *testing.T) {
		t.Parallel()
		// [A]rrange
		var value any = "not an int"
		// [A]ct
		actual := CastOrZero[int](value)
		// [A]ssert
		must.Eq(t, 0, actual)
	})
}
