package optional

import go_utils "github.com/azat-dev/go-utils"

// Optional represents a value that may or may not be present (Some or None).
type Optional[T any] struct {
	value   T
	present bool
}

// Some creates an Optional with a present value.
// Panics if the value is nil (for interface and pointer types).
func Some[T any](v T) Optional[T] {
	// Use reflection to check if the value is nil
	if go_utils.IsNil(v) {
		panic("Some() called with nil value")
	}
	return Optional[T]{
		value:   v,
		present: true,
	}
}

// None creates an empty Optional for the specified type.
func None[T any]() Optional[T] {
	var zero T
	return Optional[T]{
		present: false,
		value:   zero,
	}
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
func (o Optional[T]) Get() (
	T,
	bool,
) {
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
func Map[T, U any](
	o Optional[T],
	f func(T) U,
) Optional[U] {
	if o.present {
		return Some(f(o.value))
	}
	return None[U]()
}

// FlatMap (or AndThen) applies a function that itself returns an Optional.
// Use for chaining calls where each step might return an empty value.
func FlatMap[T, U any](
	o Optional[T],
	f func(T) Optional[U],
) Optional[U] {
	if o.present {
		return f(o.value)
	}
	return None[U]()
}

// NewFromNullable creates an Optional from a nullable value.
// If the value is nil (for pointer or interface types), it returns None.
// Otherwise, it returns Some.
func NewFromNullable[T any](v T) Optional[T] {
	if go_utils.IsNil(v) {
		return None[T]()
	}
	return Some(v)
}

// NewFromNullablePointer creates an Optional from a pointer.
// If the pointer is nil, it returns None.
// Otherwise, it dereferences the pointer and returns Some(*ptr).
func NewFromNullablePointer[T any](ptr *T) Optional[T] {
	if ptr == nil {
		return None[T]()
	}
	return Some(*ptr)
}

func Equal[T any](
	a, b Optional[T],
	eq func(
		T,
		T,
	) bool,
) bool {
	if a.present != b.present {
		return false
	}
	if !a.present {
		return true // оба None
	}
	return eq(a.value, b.value)
}
