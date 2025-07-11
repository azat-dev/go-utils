package go_utils

// Optional represents a value that may or may not be present (Some or None).
type Optional[T any] struct {
	value   T
	present bool
}

// Some creates an Optional with a present value.
func Some[T any](v T) Optional[T] {
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

// MapOptional applies a function to the value inside the Optional, if it's present.
// It returns a new Optional with the result. If the original Optional was None, it returns None.
func MapOptional[T, U any](o Optional[T], f func(T) U) Optional[U] {
	if o.present {
		return Some(f(o.value))
	}
	return None[U]()
}

// FlatMapOptional (or AndThen) applies a function that itself returns an Optional.
// Use for chaining calls where each step might return an empty value.
func FlatMapOptional[T, U any](o Optional[T], f func(T) Optional[U]) Optional[U] {
	if o.present {
		return f(o.value)
	}
	return None[U]()
}

// OptionalFromResult converts a Result into an Optional.
// If the Result contains a successful value (IsOk), it returns a Some Optional with that value.
// If the Result contains an error (IsErr), it returns a None Optional.
func OptionalFromResult[T any](r Result[T]) Optional[T] {
	if r.IsOk() {
		// We use .Get() to extract the value when we are sure it's present (IsOk).
		// In Go, it's idiomatic to return (value, error), so we just take the value.
		val, _ := r.Get()
		return Some(val)
	}
	return None[T]()
}
