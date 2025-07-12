package optional

import (
	"testing"
)

func TestSome(t *testing.T) {
	t.Run("with valid value", func(t *testing.T) {
		opt := Some("hello")
		if !opt.IsSome() {
			t.Error("Expected Some to return a present value")
		}
		if opt.IsNone() {
			t.Error("Expected Some to not be None")
		}
		value, ok := opt.Get()
		if !ok {
			t.Error("Expected Get to return true for Some")
		}
		if value != "hello" {
			t.Errorf("Expected value 'hello', got '%s'", value)
		}
	})

	t.Run("with zero value", func(t *testing.T) {
		// Zero values should be allowed
		opt := Some("")
		if !opt.IsSome() {
			t.Error("Expected Some to return a present value for zero value")
		}
		value, ok := opt.Get()
		if !ok {
			t.Error("Expected Get to return true for Some with zero value")
		}
		if value != "" {
			t.Errorf("Expected empty string, got '%s'", value)
		}
	})

	t.Run("with int zero value", func(t *testing.T) {
		opt := Some(0)
		if !opt.IsSome() {
			t.Error("Expected Some to return a present value for int zero")
		}
		value, ok := opt.Get()
		if !ok {
			t.Error("Expected Get to return true for Some with int zero")
		}
		if value != 0 {
			t.Errorf("Expected 0, got %d", value)
		}
	})

	t.Run("with nil interface - should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected Some(nil) to panic")
			} else {
				expectedMsg := "Some() called with nil value"
				if r != expectedMsg {
					t.Errorf("Expected panic message '%s', got '%s'", expectedMsg, r)
				}
			}
		}()
		Some[any](nil)
	})

	t.Run("with nil pointer - should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected Some(nil pointer) to panic")
			} else {
				expectedMsg := "Some() called with nil value"
				if r != expectedMsg {
					t.Errorf("Expected panic message '%s', got '%s'", expectedMsg, r)
				}
			}
		}()
		var ptr *int
		Some(ptr)
	})

	t.Run("with nil slice - should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected Some(nil slice) to panic")
			} else {
				expectedMsg := "Some() called with nil value"
				if r != expectedMsg {
					t.Errorf("Expected panic message '%s', got '%s'", expectedMsg, r)
				}
			}
		}()
		var slice []int
		Some(slice)
	})

	t.Run("with nil map - should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected Some(nil map) to panic")
			} else {
				expectedMsg := "Some() called with nil value"
				if r != expectedMsg {
					t.Errorf("Expected panic message '%s', got '%s'", expectedMsg, r)
				}
			}
		}()
		var m map[string]int
		Some(m)
	})

	t.Run("with nil function - should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected Some(nil function) to panic")
			} else {
				expectedMsg := "Some() called with nil value"
				if r != expectedMsg {
					t.Errorf("Expected panic message '%s', got '%s'", expectedMsg, r)
				}
			}
		}()
		var fn func()
		Some(fn)
	})

	t.Run("with nil channel - should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected Some(nil channel) to panic")
			} else {
				expectedMsg := "Some() called with nil value"
				if r != expectedMsg {
					t.Errorf("Expected panic message '%s', got '%s'", expectedMsg, r)
				}
			}
		}()
		var ch chan int
		Some(ch)
	})
}

func TestNone(t *testing.T) {
	opt := None[string]()
	if opt.IsSome() {
		t.Error("Expected None to not be Some")
	}
	if !opt.IsNone() {
		t.Error("Expected None to be None")
	}
	value, ok := opt.Get()
	if ok {
		t.Error("Expected Get to return false for None")
	}
	if value != "" {
		t.Errorf("Expected empty string for None, got '%s'", value)
	}
}

func TestIsSome(t *testing.T) {
	t.Run("Some value", func(t *testing.T) {
		opt := Some("test")
		if !opt.IsSome() {
			t.Error("Expected IsSome to return true for Some")
		}
	})

	t.Run("None value", func(t *testing.T) {
		opt := None[string]()
		if opt.IsSome() {
			t.Error("Expected IsSome to return false for None")
		}
	})
}

func TestIsNone(t *testing.T) {
	t.Run("Some value", func(t *testing.T) {
		opt := Some("test")
		if opt.IsNone() {
			t.Error("Expected IsNone to return false for Some")
		}
	})

	t.Run("None value", func(t *testing.T) {
		opt := None[string]()
		if !opt.IsNone() {
			t.Error("Expected IsNone to return true for None")
		}
	})
}

func TestGet(t *testing.T) {
	t.Run("Some value", func(t *testing.T) {
		opt := Some("hello")
		value, ok := opt.Get()
		if !ok {
			t.Error("Expected Get to return true for Some")
		}
		if value != "hello" {
			t.Errorf("Expected 'hello', got '%s'", value)
		}
	})

	t.Run("None value", func(t *testing.T) {
		opt := None[string]()
		value, ok := opt.Get()
		if ok {
			t.Error("Expected Get to return false for None")
		}
		if value != "" {
			t.Errorf("Expected empty string, got '%s'", value)
		}
	})
}

func TestUnwrap(t *testing.T) {
	t.Run("Some value", func(t *testing.T) {
		opt := Some("hello")
		value := opt.Unwrap()
		if value != "hello" {
			t.Errorf("Expected 'hello', got '%s'", value)
		}
	})

	t.Run("None value - should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected Unwrap on None to panic")
			} else {
				expectedMsg := "called Unwrap() on a None option"
				if r != expectedMsg {
					t.Errorf("Expected panic message '%s', got '%s'", expectedMsg, r)
				}
			}
		}()
		opt := None[string]()
		opt.Unwrap()
	})
}

func TestUnwrapOr(t *testing.T) {
	t.Run("Some value", func(t *testing.T) {
		opt := Some("hello")
		value := opt.UnwrapOr("default")
		if value != "hello" {
			t.Errorf("Expected 'hello', got '%s'", value)
		}
	})

	t.Run("None value", func(t *testing.T) {
		opt := None[string]()
		value := opt.UnwrapOr("default")
		if value != "default" {
			t.Errorf("Expected 'default', got '%s'", value)
		}
	})
}

func TestMap(t *testing.T) {
	t.Run("Some value", func(t *testing.T) {
		opt := Some(5)
		mapped := Map(opt, func(x int) string {
			return "value: " + string(rune(x+'0'))
		})
		if !mapped.IsSome() {
			t.Error("Expected mapped result to be Some")
		}
		value, ok := mapped.Get()
		if !ok {
			t.Error("Expected Get to return true for mapped Some")
		}
		expected := "value: 5"
		if value != expected {
			t.Errorf("Expected '%s', got '%s'", expected, value)
		}
	})

	t.Run("None value", func(t *testing.T) {
		opt := None[int]()
		mapped := Map(opt, func(x int) string {
			return "value: " + string(rune(x+'0'))
		})
		if !mapped.IsNone() {
			t.Error("Expected mapped result to be None")
		}
	})
}

func TestFlatMap(t *testing.T) {
	t.Run("Some value with Some result", func(t *testing.T) {
		opt := Some(5)
		flatMapped := FlatMap(opt, func(x int) Optional[string] {
			return Some("value: " + string(rune(x+'0')))
		})
		if !flatMapped.IsSome() {
			t.Error("Expected flatMapped result to be Some")
		}
		value, ok := flatMapped.Get()
		if !ok {
			t.Error("Expected Get to return true for flatMapped Some")
		}
		expected := "value: 5"
		if value != expected {
			t.Errorf("Expected '%s', got '%s'", expected, value)
		}
	})

	t.Run("Some value with None result", func(t *testing.T) {
		opt := Some(5)
		flatMapped := FlatMap(opt, func(x int) Optional[string] {
			return None[string]()
		})
		if !flatMapped.IsNone() {
			t.Error("Expected flatMapped result to be None")
		}
	})

	t.Run("None value", func(t *testing.T) {
		opt := None[int]()
		flatMapped := FlatMap(opt, func(x int) Optional[string] {
			return Some("value: " + string(rune(x+'0')))
		})
		if !flatMapped.IsNone() {
			t.Error("Expected flatMapped result to be None")
		}
	})
}

func TestIsNil(t *testing.T) {
	t.Run("nil interface", func(t *testing.T) {
		if !isNil(nil) {
			t.Error("Expected isNil to return true for nil interface")
		}
	})

	t.Run("nil pointer", func(t *testing.T) {
		var ptr *int
		if !isNil(ptr) {
			t.Error("Expected isNil to return true for nil pointer")
		}
	})

	t.Run("nil slice", func(t *testing.T) {
		var slice []int
		if !isNil(slice) {
			t.Error("Expected isNil to return true for nil slice")
		}
	})

	t.Run("nil map", func(t *testing.T) {
		var m map[string]int
		if !isNil(m) {
			t.Error("Expected isNil to return true for nil map")
		}
	})

	t.Run("nil function", func(t *testing.T) {
		var fn func()
		if !isNil(fn) {
			t.Error("Expected isNil to return true for nil function")
		}
	})

	t.Run("nil channel", func(t *testing.T) {
		var ch chan int
		if !isNil(ch) {
			t.Error("Expected isNil to return true for nil channel")
		}
	})

	t.Run("non-nil values", func(t *testing.T) {
		if isNil("hello") {
			t.Error("Expected isNil to return false for string")
		}
		if isNil(42) {
			t.Error("Expected isNil to return false for int")
		}
		if isNil(true) {
			t.Error("Expected isNil to return false for bool")
		}
	})

	t.Run("non-nil pointer", func(t *testing.T) {
		value := 42
		ptr := &value
		if isNil(ptr) {
			t.Error("Expected isNil to return false for non-nil pointer")
		}
	})

	t.Run("non-nil slice", func(t *testing.T) {
		slice := []int{1, 2, 3}
		if isNil(slice) {
			t.Error("Expected isNil to return false for non-nil slice")
		}
	})
}
