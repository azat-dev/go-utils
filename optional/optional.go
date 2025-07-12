package optional

import "reflect"

// Optional represents a value that may or may not be present (Some or None).
type Optional[T any] struct {
	value   T
	present bool
}

// isNil checks if a value is nil using reflection.
// This handles interface types and pointer types that could be nil.
func isNil(v any) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return rv.IsNil()
	default:
		return false
	}
}

// Some creates an Optional with a present value.
// Panics if the value is nil (for interface and pointer types).
func Some[T any](v T) Optional[T] {
	// Use reflection to check if the value is nil
	if isNil(v) {
		panic("Some() called with nil value")
	}
	return Optional[T]{value: v, present: true}
}

// None creates an empty Optional for the specified type.
func None[T any]() Optional[T] {
	var zero T
	return Optional[T]{present: false, value: zero}
}

// IsSome returns true if the value is present.
func (o Optional[T]) IsSome() bool {
	return o.present
}

// IsNone returns true if the value is absent.
func (o Optional[T]) IsNone() bool {
	return !o.present
}

// Get returns the value and true if it's present. Otherwise, it returns the zero-value and false.
// This is the idiomatic Go way to work with optional values.
func (o Optional[T]) Get() (T, bool) {
	if o.present {
		return o.value, true
	}
	var zero T
	return zero, false
}

// Unwrap returns the value if it's present. Otherwise, it panics.
// Use when you are certain the value exists.
func (o Optional[T]) Unwrap() T {
	if !o.present {
		panic("called Unwrap() on a None option")
	}
	return o.value
}

// UnwrapOr returns the value if it's present. Otherwise, it returns the default value.
func (o Optional[T]) UnwrapOr(defaultValue T) T {
	if o.present {
		return o.value
	}
	return defaultValue
}

// Map applies a function to the value inside the Optional, if it's present.
// It returns a new Optional with the result. If the original Optional was None, it returns None.
func Map[T, U any](o Optional[T], f func(T) U) Optional[U] {
	if o.present {
		return Some(f(o.value))
	}
	return None[U]()
}

// FlatMap (or AndThen) applies a function that itself returns an Optional.
// Use for chaining calls where each step might return an empty value.
func FlatMap[T, U any](o Optional[T], f func(T) Optional[U]) Optional[U] {
	if o.present {
		return f(o.value)
	}
	return None[U]()
}
