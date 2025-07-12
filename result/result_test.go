package result

import (
	"errors"
	"testing"

	go_utils "github.com/azat-dev/go-utils"
)

func TestOk(t *testing.T) {
	t.Run("with valid value", func(t *testing.T) {
		result := Ok("hello")
		if !result.IsOk() {
			t.Error("Expected Ok to return a successful result")
		}
		if result.IsErr() {
			t.Error("Expected Ok to not be an error")
		}
		value, err := result.Get()
		if err != nil {
			t.Error("Expected Get to return nil error for Ok")
		}
		if value != "hello" {
			t.Errorf("Expected value 'hello', got '%s'", value)
		}
	})

	t.Run("with zero value", func(t *testing.T) {
		// Zero values should be allowed
		result := Ok("")
		if !result.IsOk() {
			t.Error("Expected Ok to return a successful result for zero value")
		}
		value, err := result.Get()
		if err != nil {
			t.Error("Expected Get to return nil error for Ok with zero value")
		}
		if value != "" {
			t.Errorf("Expected empty string, got '%s'", value)
		}
	})

	t.Run("with int zero value", func(t *testing.T) {
		result := Ok(0)
		if !result.IsOk() {
			t.Error("Expected Ok to return a successful result for int zero")
		}
		value, err := result.Get()
		if err != nil {
			t.Error("Expected Get to return nil error for Ok with int zero")
		}
		if value != 0 {
			t.Errorf("Expected 0, got %d", value)
		}
	})

	t.Run("with nil interface - should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected Ok(nil) to panic")
			} else {
				expectedMsg := "Ok() called with nil value"
				if r != expectedMsg {
					t.Errorf("Expected panic message '%s', got '%s'", expectedMsg, r)
				}
			}
		}()
		Ok[any](nil)
	})

	t.Run("with nil pointer - should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected Ok(nil pointer) to panic")
			} else {
				expectedMsg := "Ok() called with nil value"
				if r != expectedMsg {
					t.Errorf("Expected panic message '%s', got '%s'", expectedMsg, r)
				}
			}
		}()
		var ptr *int
		Ok(ptr)
	})

	t.Run("with nil slice - should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected Ok(nil slice) to panic")
			} else {
				expectedMsg := "Ok() called with nil value"
				if r != expectedMsg {
					t.Errorf("Expected panic message '%s', got '%s'", expectedMsg, r)
				}
			}
		}()
		var slice []int
		Ok(slice)
	})

	t.Run("with nil map - should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected Ok(nil map) to panic")
			} else {
				expectedMsg := "Ok() called with nil value"
				if r != expectedMsg {
					t.Errorf("Expected panic message '%s', got '%s'", expectedMsg, r)
				}
			}
		}()
		var m map[string]int
		Ok(m)
	})

	t.Run("with nil function - should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected Ok(nil function) to panic")
			} else {
				expectedMsg := "Ok() called with nil value"
				if r != expectedMsg {
					t.Errorf("Expected panic message '%s', got '%s'", expectedMsg, r)
				}
			}
		}()
		var fn func()
		Ok(fn)
	})

	t.Run("with nil channel - should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected Ok(nil channel) to panic")
			} else {
				expectedMsg := "Ok() called with nil value"
				if r != expectedMsg {
					t.Errorf("Expected panic message '%s', got '%s'", expectedMsg, r)
				}
			}
		}()
		var ch chan int
		Ok(ch)
	})
}

func TestErr(t *testing.T) {
	testErr := errors.New("test error")
	result := Err[string](testErr)
	if result.IsOk() {
		t.Error("Expected Err to not be Ok")
	}
	if !result.IsErr() {
		t.Error("Expected Err to be an error")
	}
	value, err := result.Get()
	if err == nil {
		t.Error("Expected Get to return error for Err")
	}
	if err != testErr {
		t.Errorf("Expected error '%v', got '%v'", testErr, err)
	}
	if value != "" {
		t.Errorf("Expected empty string for Err, got '%s'", value)
	}
}

func TestErrorF(t *testing.T) {
	result := ErrorF[string]("test error: %s", "details")
	if result.IsOk() {
		t.Error("Expected ErrorF to not be Ok")
	}
	if !result.IsErr() {
		t.Error("Expected ErrorF to be an error")
	}
	value, err := result.Get()
	if err == nil {
		t.Error("Expected Get to return error for ErrorF")
	}
	expectedMsg := "test error: details"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
	if value != "" {
		t.Errorf("Expected empty string for ErrorF, got '%s'", value)
	}
}

func TestIsOk(t *testing.T) {
	t.Run("Ok result", func(t *testing.T) {
		result := Ok("test")
		if !result.IsOk() {
			t.Error("Expected IsOk to return true for Ok")
		}
	})

	t.Run("Err result", func(t *testing.T) {
		result := Err[string](errors.New("test"))
		if result.IsOk() {
			t.Error("Expected IsOk to return false for Err")
		}
	})
}

func TestIsErr(t *testing.T) {
	t.Run("Ok result", func(t *testing.T) {
		result := Ok("test")
		if result.IsErr() {
			t.Error("Expected IsErr to return false for Ok")
		}
	})

	t.Run("Err result", func(t *testing.T) {
		result := Err[string](errors.New("test"))
		if !result.IsErr() {
			t.Error("Expected IsErr to return true for Err")
		}
	})
}

func TestGet(t *testing.T) {
	t.Run("Ok result", func(t *testing.T) {
		result := Ok("hello")
		value, err := result.Get()
		if err != nil {
			t.Error("Expected Get to return nil error for Ok")
		}
		if value != "hello" {
			t.Errorf("Expected 'hello', got '%s'", value)
		}
	})

	t.Run("Err result", func(t *testing.T) {
		testErr := errors.New("test error")
		result := Err[string](testErr)
		value, err := result.Get()
		if err == nil {
			t.Error("Expected Get to return error for Err")
		}
		if err != testErr {
			t.Errorf("Expected error '%v', got '%v'", testErr, err)
		}
		if value != "" {
			t.Errorf("Expected empty string, got '%s'", value)
		}
	})
}

func TestUnwrap(t *testing.T) {
	t.Run("Ok result", func(t *testing.T) {
		result := Ok("hello")
		value := result.Unwrap()
		if value != "hello" {
			t.Errorf("Expected 'hello', got '%s'", value)
		}
	})

	t.Run("Err result - should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected Unwrap on Err to panic")
			} else {
				expectedMsg := "called Unwrap() on an Err Result: test error"
				if r != expectedMsg {
					t.Errorf("Expected panic message '%s', got '%s'", expectedMsg, r)
				}
			}
		}()
		result := Err[string](errors.New("test error"))
		result.Unwrap()
	})
}

func TestUnwrapOr(t *testing.T) {
	t.Run("Ok result", func(t *testing.T) {
		result := Ok("hello")
		value := result.UnwrapOr("default")
		if value != "hello" {
			t.Errorf("Expected 'hello', got '%s'", value)
		}
	})

	t.Run("Err result", func(t *testing.T) {
		result := Err[string](errors.New("test error"))
		value := result.UnwrapOr("default")
		if value != "default" {
			t.Errorf("Expected 'default', got '%s'", value)
		}
	})
}

func TestUnwrapErr(t *testing.T) {
	t.Run("Err result", func(t *testing.T) {
		testErr := errors.New("test error")
		result := Err[string](testErr)
		err := result.UnwrapErr()
		if err != testErr {
			t.Errorf("Expected error '%v', got '%v'", testErr, err)
		}
	})

	t.Run("Ok result - should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected UnwrapErr on Ok to panic")
			} else {
				expectedMsg := "called UnwrapErr() on an Ok Result"
				if r != expectedMsg {
					t.Errorf("Expected panic message '%s', got '%s'", expectedMsg, r)
				}
			}
		}()
		result := Ok("hello")
		_ = result.UnwrapErr()
	})
}

func TestMustGet(t *testing.T) {
	t.Run("Ok result", func(t *testing.T) {
		result := Ok("hello")
		value := result.MustGet()
		if value != "hello" {
			t.Errorf("Expected 'hello', got '%s'", value)
		}
	})

	t.Run("Err result - should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected MustGet on Err to panic")
			} else {
				expectedMsg := "called MustGet() on an Err Result: test error"
				if r != expectedMsg {
					t.Errorf("Expected panic message '%s', got '%s'", expectedMsg, r)
				}
			}
		}()
		result := Err[string](errors.New("test error"))
		result.MustGet()
	})
}

func TestMapResult(t *testing.T) {
	t.Run("Ok result", func(t *testing.T) {
		result := Ok(5)
		mapped := MapResult(result, func(x int) string {
			return "value: " + string(rune(x+'0'))
		})
		if !mapped.IsOk() {
			t.Error("Expected mapped result to be Ok")
		}
		value, err := mapped.Get()
		if err != nil {
			t.Error("Expected Get to return nil error for mapped Ok")
		}
		expected := "value: 5"
		if value != expected {
			t.Errorf("Expected '%s', got '%s'", expected, value)
		}
	})

	t.Run("Err result", func(t *testing.T) {
		testErr := errors.New("test error")
		result := Err[int](testErr)
		mapped := MapResult(result, func(x int) string {
			return "value: " + string(rune(x+'0'))
		})
		if !mapped.IsErr() {
			t.Error("Expected mapped result to be Err")
		}
		_, err := mapped.Get()
		if err != testErr {
			t.Errorf("Expected error '%v', got '%v'", testErr, err)
		}
	})
}

func TestFlatMapResult(t *testing.T) {
	t.Run("Ok result with Ok result", func(t *testing.T) {
		result := Ok(5)
		flatMapped := FlatMapResult(result, func(x int) Result[string] {
			return Ok("value: " + string(rune(x+'0')))
		})
		if !flatMapped.IsOk() {
			t.Error("Expected flatMapped result to be Ok")
		}
		value, err := flatMapped.Get()
		if err != nil {
			t.Error("Expected Get to return nil error for flatMapped Ok")
		}
		expected := "value: 5"
		if value != expected {
			t.Errorf("Expected '%s', got '%s'", expected, value)
		}
	})

	t.Run("Ok result with Err result", func(t *testing.T) {
		testErr := errors.New("test error")
		result := Ok(5)
		flatMapped := FlatMapResult(result, func(x int) Result[string] {
			return Err[string](testErr)
		})
		if !flatMapped.IsErr() {
			t.Error("Expected flatMapped result to be Err")
		}
		_, err := flatMapped.Get()
		if err != testErr {
			t.Errorf("Expected error '%v', got '%v'", testErr, err)
		}
	})

	t.Run("Err result", func(t *testing.T) {
		testErr := errors.New("test error")
		result := Err[int](testErr)
		flatMapped := FlatMapResult(result, func(x int) Result[string] {
			return Ok("value: " + string(rune(x+'0')))
		})
		if !flatMapped.IsErr() {
			t.Error("Expected flatMapped result to be Err")
		}
		_, err := flatMapped.Get()
		if err != testErr {
			t.Errorf("Expected error '%v', got '%v'", testErr, err)
		}
	})
}

func TestMapErr(t *testing.T) {
	t.Run("Err result", func(t *testing.T) {
		testErr := errors.New("test error")
		result := Err[string](testErr)
		mapped := MapErr(result, func(e error) error {
			return errors.New("wrapped: " + e.Error())
		})
		if !mapped.IsErr() {
			t.Error("Expected mapped result to be Err")
		}
		_, err := mapped.Get()
		expectedMsg := "wrapped: test error"
		if err.Error() != expectedMsg {
			t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
		}
	})

	t.Run("Ok result", func(t *testing.T) {
		result := Ok("hello")
		mapped := MapErr(result, func(e error) error {
			return errors.New("wrapped: " + e.Error())
		})
		if !mapped.IsOk() {
			t.Error("Expected mapped result to be Ok")
		}
		value, err := mapped.Get()
		if err != nil {
			t.Error("Expected Get to return nil error for mapped Ok")
		}
		if value != "hello" {
			t.Errorf("Expected 'hello', got '%s'", value)
		}
	})
}

func TestInspect(t *testing.T) {
	t.Run("Ok result", func(t *testing.T) {
		var inspected string
		result := Ok("hello")
		inspectedResult := result.Inspect(func(v string) {
			inspected = v
		})
		if inspected != "hello" {
			t.Errorf("Expected inspected value 'hello', got '%s'", inspected)
		}
		if !inspectedResult.IsOk() {
			t.Error("Expected inspected result to be Ok")
		}
		value, err := inspectedResult.Get()
		if err != nil {
			t.Error("Expected Get to return nil error for inspected Ok")
		}
		if value != "hello" {
			t.Errorf("Expected 'hello', got '%s'", value)
		}
	})

	t.Run("Err result", func(t *testing.T) {
		var inspected string
		result := Err[string](errors.New("test error"))
		inspectedResult := result.Inspect(func(v string) {
			inspected = v
		})
		if inspected != "" {
			t.Errorf("Expected no inspection for Err, got '%s'", inspected)
		}
		if !inspectedResult.IsErr() {
			t.Error("Expected inspected result to be Err")
		}
	})
}

func TestInspectErr(t *testing.T) {
	t.Run("Err result", func(t *testing.T) {
		var inspected error
		testErr := errors.New("test error")
		result := Err[string](testErr)
		inspectedResult := result.InspectErr(func(e error) {
			inspected = e
		})
		if inspected != testErr {
			t.Errorf("Expected inspected error '%v', got '%v'", testErr, inspected)
		}
		if !inspectedResult.IsErr() {
			t.Error("Expected inspected result to be Err")
		}
		_, err := inspectedResult.Get()
		if err != testErr {
			t.Errorf("Expected error '%v', got '%v'", testErr, err)
		}
	})

	t.Run("Ok result", func(t *testing.T) {
		var inspected error
		result := Ok("hello")
		inspectedResult := result.InspectErr(func(e error) {
			inspected = e
		})
		if inspected != nil {
			t.Errorf("Expected no error inspection for Ok, got '%v'", inspected)
		}
		if !inspectedResult.IsOk() {
			t.Error("Expected inspected result to be Ok")
		}
	})
}

func TestOrElse(t *testing.T) {
	t.Run("Ok result", func(t *testing.T) {
		result := Ok("hello")
		alternative := Ok("alternative")
		orElse := result.OrElse(alternative)
		if !orElse.IsOk() {
			t.Error("Expected OrElse to be Ok")
		}
		value, err := orElse.Get()
		if err != nil {
			t.Error("Expected Get to return nil error for OrElse Ok")
		}
		if value != "hello" {
			t.Errorf("Expected 'hello', got '%s'", value)
		}
	})

	t.Run("Err result", func(t *testing.T) {
		result := Err[string](errors.New("test error"))
		alternative := Ok("alternative")
		orElse := result.OrElse(alternative)
		if !orElse.IsOk() {
			t.Error("Expected OrElse to be Ok")
		}
		value, err := orElse.Get()
		if err != nil {
			t.Error("Expected Get to return nil error for OrElse alternative")
		}
		if value != "alternative" {
			t.Errorf("Expected 'alternative', got '%s'", value)
		}
	})
}

func TestOrElseDo(t *testing.T) {
	t.Run("Ok result", func(t *testing.T) {
		result := Ok("hello")
		orElse := result.OrElseDo(func(e error) Result[string] {
			return Ok("alternative")
		})
		if !orElse.IsOk() {
			t.Error("Expected OrElseDo to be Ok")
		}
		value, err := orElse.Get()
		if err != nil {
			t.Error("Expected Get to return nil error for OrElseDo Ok")
		}
		if value != "hello" {
			t.Errorf("Expected 'hello', got '%s'", value)
		}
	})

	t.Run("Err result", func(t *testing.T) {
		result := Err[string](errors.New("test error"))
		orElse := result.OrElseDo(func(e error) Result[string] {
			return Ok("alternative")
		})
		if !orElse.IsOk() {
			t.Error("Expected OrElseDo to be Ok")
		}
		value, err := orElse.Get()
		if err != nil {
			t.Error("Expected Get to return nil error for OrElseDo alternative")
		}
		if value != "alternative" {
			t.Errorf("Expected 'alternative', got '%s'", value)
		}
	})
}

func TestToOptional(t *testing.T) {
	t.Run("Ok result", func(t *testing.T) {
		result := Ok("hello")
		opt := result.ToOptional()
		if !opt.IsSome() {
			t.Error("Expected ToOptional to return Some for Ok")
		}
		value, ok := opt.Get()
		if !ok {
			t.Error("Expected Get to return true for Some from Ok")
		}
		if value != "hello" {
			t.Errorf("Expected 'hello', got '%s'", value)
		}
	})

	t.Run("Err result", func(t *testing.T) {
		result := Err[string](errors.New("test error"))
		opt := result.ToOptional()
		if !opt.IsNone() {
			t.Error("Expected ToOptional to return None for Err")
		}
		_, ok := opt.Get()
		if ok {
			t.Error("Expected Get to return false for None from Err")
		}
	})
}

func TestIsNil(t *testing.T) {
	t.Run("nil interface", func(t *testing.T) {
		if !go_utils.IsNil(nil) {
			t.Error("Expected isNil to return true for nil interface")
		}
	})

	t.Run("nil pointer", func(t *testing.T) {
		var ptr *int
		if !go_utils.IsNil(ptr) {
			t.Error("Expected isNil to return true for nil pointer")
		}
	})

	t.Run("nil slice", func(t *testing.T) {
		var slice []int
		if !go_utils.IsNil(slice) {
			t.Error("Expected isNil to return true for nil slice")
		}
	})

	t.Run("nil map", func(t *testing.T) {
		var m map[string]int
		if !go_utils.IsNil(m) {
			t.Error("Expected isNil to return true for nil map")
		}
	})

	t.Run("nil function", func(t *testing.T) {
		var fn func()
		if !go_utils.IsNil(fn) {
			t.Error("Expected isNil to return true for nil function")
		}
	})

	t.Run("nil channel", func(t *testing.T) {
		var ch chan int
		if !go_utils.IsNil(ch) {
			t.Error("Expected isNil to return true for nil channel")
		}
	})

	t.Run("non-nil values", func(t *testing.T) {
		if go_utils.IsNil("hello") {
			t.Error("Expected isNil to return false for string")
		}
		if go_utils.IsNil(42) {
			t.Error("Expected isNil to return false for int")
		}
		if go_utils.IsNil(true) {
			t.Error("Expected isNil to return false for bool")
		}
	})

	t.Run("non-nil pointer", func(t *testing.T) {
		value := 42
		ptr := &value
		if go_utils.IsNil(ptr) {
			t.Error("Expected isNil to return false for non-nil pointer")
		}
	})

	t.Run("non-nil slice", func(t *testing.T) {
		slice := []int{1, 2, 3}
		if go_utils.IsNil(slice) {
			t.Error("Expected isNil to return false for non-nil slice")
		}
	})
}
