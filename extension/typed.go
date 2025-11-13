package extension

// MustCast attempts to cast the provided value to type T.
// If the value is of type T, it returns the value.
// If the value cannot be cast to type T, it panics.
//
// Example:
//
//	var value any = 42
//	result := MustCast[int](value) // 42
//
//	var value any = "string"
//	result := MustCast[int](value) // panics!
func MustCast[T any](original any) T {
	switch o := original.(type) {
	case T:
		return o
	default:
		panic("failed to coerce type")
	}
}

// CastOrZero attempts to cast the provided value to type V.
// If the value is of type V, it returns the value.
// If the value cannot be cast to type V, it returns the zero value of type V.
//
// Example:
//
//	var value any = 42
//	result := CastOrZero[int](value) // 42
//
//	var value any = "string"
//	result := CastOrZero[int](value) // 0
func CastOrZero[V any](original any) V {
	switch o := original.(type) {
	case V:
		return o
	default:
		return *new(V)
	}
}
