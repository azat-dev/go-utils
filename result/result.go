package result

import (
	"fmt"

	go_utils "github.com/azat-dev/go-utils"

	"github.com/azat-dev/go-utils/optional"
)

// Result represents the outcome of an operation, which can either be a successful value (Ok) or an error (Err).
// T - the type of the successful value.
type Result[T any] struct {
	value T
	err   error
}

// Ok creates a Result with a successful value.
// Panics if the value is nil (for interface and pointer types).
// Use this to indicate a successful operation.
func Ok[T any](v T) Result[T] {
	// Use reflection to check if the value is nil
	if go_utils.IsNil(v) {
		panic("Ok() called with nil value")
	}
	return Result[T]{value: v, err: nil}
}

// Err creates a Result with an error.
// Use this to indicate a failed operation.
func Err[T any](e error) Result[T] {
	var zero T // Zero value for type T
	return Result[T]{value: zero, err: e}
}

// ErrorF creates a new Result with a formatted error and the zero value for the type.
// It uses fmt.Errorf for formatting the error message.
func ErrorF[T any](format string, a ...any) Result[T] {
	var zero T // Zero value for type T
	return Result[T]{value: zero, err: fmt.Errorf(format, a...)}
}

// IsOk returns true if the Result contains a successful value.
func (r Result[T]) IsOk() bool {
	return r.err == nil
}

// IsErr returns true if the Result contains an error.
func (r Result[T]) IsErr() bool {
	return r.err != nil
}

// Get returns the successful value and nil if it's present. Otherwise, it returns the zero-value and the error.
// This is the idiomatic Go way to work with Results.
func (r Result[T]) Get() (T, error) {
	return r.value, r.err
}

// Unwrap returns the successful value if it's present. Otherwise, it panics with the error.
// Use when you are certain the operation was successful.
func (r Result[T]) Unwrap() T {
	if r.err != nil {
		panic("called Unwrap() on an Err Result: " + r.err.Error())
	}
	return r.value
}

// UnwrapOr returns the successful value if it's present. Otherwise, it returns the default value.
func (r Result[T]) UnwrapOr(defaultValue T) T {
	if r.err == nil {
		return r.value
	}
	return defaultValue
}

// UnwrapErr returns the error if it's present. Otherwise, it panics.
// Use when you are certain the Result contains an error.
func (r Result[T]) UnwrapErr() error {
	if r.err == nil {
		panic("called UnwrapErr() on an Ok Result")
	}
	return r.err
}

// MustGet returns the successful value. If there's an error, it panics.
// Convenient for initialization or test code where errors are not expected.
func (r Result[T]) MustGet() T {
	if r.err != nil {
		panic("called MustGet() on an Err Result: " + r.err.Error())
	}
	return r.value
}

// MapResult applies a function to the successful value inside the Result, if it's present.
// It returns a new Result with the outcome. If the original Result was Err, it returns the same Err.
// f - a function that transforms the successful value.
func MapResult[T, U any](r Result[T], f func(T) U) Result[U] {
	if r.IsOk() {
		return Ok(f(r.value))
	}
	return Err[U](r.err) // Propagate the error
}

// FlatMapResult (or AndThen) applies a function that itself returns a Result.
// Use for chaining calls where each step might return an error.
// f - a function that takes a successful value and returns a new Result.
func FlatMapResult[T, U any](r Result[T], f func(T) Result[U]) Result[U] {
	if r.IsOk() {
		return f(r.value)
	}
	return Err[U](r.err) // Propagate the error
}

// MapErr applies a function to the error inside the Result, if it's present.
// Allows transforming or wrapping errors.
func MapErr[T any](r Result[T], f func(error) error) Result[T] {
	if r.IsErr() {
		return Err[T](f(r.err))
	}
	return r // Return the Ok Result unchanged
}

// Inspect executes an action (function) with the successful value, if present,
// without altering the Result itself. Useful for logging or debugging.
func (r Result[T]) Inspect(f func(T)) Result[T] {
	if r.IsOk() {
		f(r.value)
	}
	return r
}

// InspectErr executes an action (function) with the error, if present,
// without altering the Result itself. Useful for logging errors.
func (r Result[T]) InspectErr(f func(error)) Result[T] {
	if r.IsErr() {
		f(r.err)
	}
	return r
}

// OrElse returns the current Result if it's Ok. Otherwise, it returns another Result.
// Useful for providing a fallback option in case of an error.
func (r Result[T]) OrElse(alternative Result[T]) Result[T] {
	if r.IsOk() {
		return r
	}
	return alternative
}

// OrElseDo returns the current Result if it's Ok. Otherwise, it calls a supplier function to get another Result.
func (r Result[T]) OrElseDo(f func(error) Result[T]) Result[T] {
	if r.IsOk() {
		return r
	}
	return f(r.err)
}

// ToOptional converts a Result into an Optional.
// If the Result is Ok, the Optional will be Some. If the Result is Err, the Optional will be None.
func (r Result[T]) ToOptional() optional.Optional[T] {
	if r.IsOk() {
		return optional.Some(r.value)
	}
	return optional.None[T]()
}
