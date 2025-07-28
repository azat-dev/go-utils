package optional

import (
	"slices"
	"testing"

	go_utils "github.com/azat-dev/go-utils"
)

func TestSome(t *testing.T) {
	t.Run(
		"with valid value", func(t *testing.T) {
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
		},
	)

	t.Run(
		"with zero value", func(t *testing.T) {
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
		},
	)

	t.Run(
		"with int zero value", func(t *testing.T) {
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
		},
	)

	t.Run(
		"with nil interface - should panic", func(t *testing.T) {
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
		},
	)

	t.Run(
		"with nil pointer - should panic", func(t *testing.T) {
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
		},
	)

	t.Run(
		"with nil slice - should panic", func(t *testing.T) {
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
		},
	)

	t.Run(
		"with nil map - should panic", func(t *testing.T) {
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
		},
	)

	t.Run(
		"with nil function - should panic", func(t *testing.T) {
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
		},
	)

	t.Run(
		"with nil channel - should panic", func(t *testing.T) {
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
		},
	)
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
	t.Run(
		"Some value", func(t *testing.T) {
			opt := Some("test")
			if !opt.IsSome() {
				t.Error("Expected IsSome to return true for Some")
			}
		},
	)

	t.Run(
		"None value", func(t *testing.T) {
			opt := None[string]()
			if opt.IsSome() {
				t.Error("Expected IsSome to return false for None")
			}
		},
	)
}

func TestIsNone(t *testing.T) {
	t.Run(
		"Some value", func(t *testing.T) {
			opt := Some("test")
			if opt.IsNone() {
				t.Error("Expected IsNone to return false for Some")
			}
		},
	)

	t.Run(
		"None value", func(t *testing.T) {
			opt := None[string]()
			if !opt.IsNone() {
				t.Error("Expected IsNone to return true for None")
			}
		},
	)
}

func TestGet(t *testing.T) {
	t.Run(
		"Some value", func(t *testing.T) {
			opt := Some("hello")
			value, ok := opt.Get()
			if !ok {
				t.Error("Expected Get to return true for Some")
			}
			if value != "hello" {
				t.Errorf("Expected 'hello', got '%s'", value)
			}
		},
	)

	t.Run(
		"None value", func(t *testing.T) {
			opt := None[string]()
			value, ok := opt.Get()
			if ok {
				t.Error("Expected Get to return false for None")
			}
			if value != "" {
				t.Errorf("Expected empty string, got '%s'", value)
			}
		},
	)
}

func TestUnwrap(t *testing.T) {
	t.Run(
		"Some value", func(t *testing.T) {
			opt := Some("hello")
			value := opt.Unwrap()
			if value != "hello" {
				t.Errorf("Expected 'hello', got '%s'", value)
			}
		},
	)

	t.Run(
		"None value - should panic", func(t *testing.T) {
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
		},
	)
}

func TestUnwrapOr(t *testing.T) {
	t.Run(
		"Some value", func(t *testing.T) {
			opt := Some("hello")
			value := opt.UnwrapOr("default")
			if value != "hello" {
				t.Errorf("Expected 'hello', got '%s'", value)
			}
		},
	)

	t.Run(
		"None value", func(t *testing.T) {
			opt := None[string]()
			value := opt.UnwrapOr("default")
			if value != "default" {
				t.Errorf("Expected 'default', got '%s'", value)
			}
		},
	)
}

func TestMap(t *testing.T) {
	t.Run(
		"Some value", func(t *testing.T) {
			opt := Some(5)
			mapped := Map(
				opt, func(x int) string {
					return "value: " + string(rune(x+'0'))
				},
			)
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
		},
	)

	t.Run(
		"None value", func(t *testing.T) {
			opt := None[int]()
			mapped := Map(
				opt, func(x int) string {
					return "value: " + string(rune(x+'0'))
				},
			)
			if !mapped.IsNone() {
				t.Error("Expected mapped result to be None")
			}
		},
	)
}

func TestFlatMap(t *testing.T) {
	t.Run(
		"Some value with Some result", func(t *testing.T) {
			opt := Some(5)
			flatMapped := FlatMap(
				opt, func(x int) Optional[string] {
					return Some("value: " + string(rune(x+'0')))
				},
			)
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
		},
	)

	t.Run(
		"Some value with None result", func(t *testing.T) {
			opt := Some(5)
			flatMapped := FlatMap(
				opt, func(x int) Optional[string] {
					return None[string]()
				},
			)
			if !flatMapped.IsNone() {
				t.Error("Expected flatMapped result to be None")
			}
		},
	)

	t.Run(
		"None value", func(t *testing.T) {
			opt := None[int]()
			flatMapped := FlatMap(
				opt, func(x int) Optional[string] {
					return Some("value: " + string(rune(x+'0')))
				},
			)
			if !flatMapped.IsNone() {
				t.Error("Expected flatMapped result to be None")
			}
		},
	)
}

func TestIsNil(t *testing.T) {
	t.Run(
		"nil interface", func(t *testing.T) {
			if !go_utils.IsNil(nil) {
				t.Error("Expected isNil to return true for nil interface")
			}
		},
	)

	t.Run(
		"nil pointer", func(t *testing.T) {
			var ptr *int
			if !go_utils.IsNil(ptr) {
				t.Error("Expected isNil to return true for nil pointer")
			}
		},
	)

	t.Run(
		"nil slice", func(t *testing.T) {
			var slice []int
			if !go_utils.IsNil(slice) {
				t.Error("Expected isNil to return true for nil slice")
			}
		},
	)

	t.Run(
		"nil map", func(t *testing.T) {
			var m map[string]int
			if !go_utils.IsNil(m) {
				t.Error("Expected isNil to return true for nil map")
			}
		},
	)

	t.Run(
		"nil function", func(t *testing.T) {
			var fn func()
			if !go_utils.IsNil(fn) {
				t.Error("Expected isNil to return true for nil function")
			}
		},
	)

	t.Run(
		"nil channel", func(t *testing.T) {
			var ch chan int
			if !go_utils.IsNil(ch) {
				t.Error("Expected isNil to return true for nil channel")
			}
		},
	)

	t.Run(
		"non-nil values", func(t *testing.T) {
			if go_utils.IsNil("hello") {
				t.Error("Expected isNil to return false for string")
			}
			if go_utils.IsNil(42) {
				t.Error("Expected isNil to return false for int")
			}
			if go_utils.IsNil(true) {
				t.Error("Expected isNil to return false for bool")
			}
		},
	)

	t.Run(
		"non-nil pointer", func(t *testing.T) {
			value := 42
			ptr := &value
			if go_utils.IsNil(ptr) {
				t.Error("Expected isNil to return false for non-nil pointer")
			}
		},
	)

	t.Run(
		"non-nil slice", func(t *testing.T) {
			slice := []int{
				1,
				2,
				3,
			}
			if go_utils.IsNil(slice) {
				t.Error("Expected isNil to return false for non-nil slice")
			}
		},
	)
}

func TestNewFromNullable(t *testing.T) {
	t.Run(
		"non-nil pointer", func(t *testing.T) {
			val := 42
			opt := NewFromNullable(&val)
			if opt.IsNone() {
				t.Error("Expected NewFromNullable to return Some for non-nil pointer")
			}
			ptr, ok := opt.Get()
			if !ok {
				t.Error("Expected Get to return true for non-nil pointer")
			}
			if *ptr != 42 {
				t.Errorf("Expected 42, got %d", *ptr)
			}
		},
	)

	t.Run(
		"nil pointer", func(t *testing.T) {
			var ptr *int
			opt := NewFromNullable(ptr)
			if opt.IsSome() {
				t.Error("Expected NewFromNullable to return None for nil pointer")
			}
		},
	)

	t.Run(
		"nil interface", func(t *testing.T) {
			var i any = nil
			opt := NewFromNullable(i)
			if opt.IsSome() {
				t.Error("Expected NewFromNullable to return None for nil interface")
			}
		},
	)

	t.Run(
		"non-nil interface", func(t *testing.T) {
			var i any = "test"
			opt := NewFromNullable(i)
			if opt.IsNone() {
				t.Error("Expected NewFromNullable to return Some for non-nil interface")
			}
			val, ok := opt.Get()
			if !ok {
				t.Error("Expected Get to return true for non-nil interface")
			}
			if val != "test" {
				t.Errorf("Expected 'test', got '%v'", val)
			}
		},
	)

	t.Run(
		"nil slice", func(t *testing.T) {
			var s []int
			opt := NewFromNullable(s)
			if opt.IsSome() {
				t.Error("Expected NewFromNullable to return None for nil slice")
			}
		},
	)

	t.Run(
		"non-nil slice", func(t *testing.T) {
			s := []int{
				1,
				2,
				3,
			}
			opt := NewFromNullable(s)
			if opt.IsNone() {
				t.Error("Expected NewFromNullable to return Some for non-nil slice")
			}
			val, ok := opt.Get()
			if !ok || len(val) != 3 {
				t.Error("Expected Get to return valid slice")
			}
		},
	)

	t.Run(
		"non-pointer struct", func(t *testing.T) {
			type Person struct {
				Name string
			}
			p := Person{Name: "Alice"}
			opt := NewFromNullable(p)
			if opt.IsNone() {
				t.Error("Expected NewFromNullable to return Some for non-pointer struct")
			}
		},
	)
}

func TestNewFromNullablePointer(t *testing.T) {
	t.Run(
		"nil pointer", func(t *testing.T) {
			var p *int
			opt := NewFromNullablePointer(p)
			if opt.IsSome() {
				t.Error("Expected None for nil pointer")
			}
		},
	)

	t.Run(
		"non-nil pointer", func(t *testing.T) {
			val := 123
			opt := NewFromNullablePointer(&val)
			if opt.IsNone() {
				t.Error("Expected Some for non-nil pointer")
			}
			v, ok := opt.Get()
			if !ok || v != 123 {
				t.Errorf("Expected 123, got %v", v)
			}
		},
	)
}

func TestOptionalEquality_Primitives(t *testing.T) {
	type testCase[T comparable] struct {
		name     string
		a, b     Optional[T]
		expected bool
	}

	intTests := []testCase[int]{
		{
			"Equal Some(42)",
			Some(42),
			Some(42),
			true,
		},
		{
			"Different Some(42) vs Some(100)",
			Some(42),
			Some(100),
			false,
		},
		{
			"Some vs None",
			Some(42),
			None[int](),
			false,
		},
		{
			"None vs None",
			None[int](),
			None[int](),
			true,
		},
	}

	for _, tt := range intTests {
		t.Run(
			"int/"+tt.name, func(t *testing.T) {
				if (tt.a == tt.b) != tt.expected {
					t.Errorf("expected %v == %v to be %v", tt.a, tt.b, tt.expected)
				}
			},
		)
	}

	stringTests := []testCase[string]{
		{
			"Equal Some(\"hello\")",
			Some("hello"),
			Some("hello"),
			true,
		},
		{
			"Different Some(\"hello\") vs Some(\"world\")",
			Some("hello"),
			Some("world"),
			false,
		},
	}

	for _, tt := range stringTests {
		t.Run(
			"string/"+tt.name, func(t *testing.T) {
				if (tt.a == tt.b) != tt.expected {
					t.Errorf("expected %v == %v to be %v", tt.a, tt.b, tt.expected)
				}
			},
		)
	}
}

func TestOptionalEquality_Struct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	tests := []struct {
		name     string
		a, b     Optional[Person]
		expected bool
	}{
		{
			"Equal Persons",
			Some(
				Person{
					"Alice",
					30,
				},
			),
			Some(
				Person{
					"Alice",
					30,
				},
			),
			true,
		},
		{
			"Different Persons",
			Some(
				Person{
					"Alice",
					30,
				},
			),
			Some(
				Person{
					"Bob",
					25,
				},
			),
			false,
		},
		{
			"Some vs None",
			Some(
				Person{
					"Alice",
					30,
				},
			),
			None[Person](),
			false,
		},
		{
			"None vs None",
			None[Person](),
			None[Person](),
			true,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if (tt.a == tt.b) != tt.expected {
					t.Errorf("expected %v == %v to be %v", tt.a, tt.b, tt.expected)
				}
			},
		)
	}
}

func TestOptionalEquality_SlicesWithCustomEqual(t *testing.T) {
	tests := []struct {
		name     string
		a, b     Optional[[]int]
		expected bool
	}{
		{
			"Equal slices",
			Some(
				[]int{
					1,
					2,
					3,
				},
			),
			Some(
				[]int{
					1,
					2,
					3,
				},
			),
			true,
		},
		{
			"Different slices",
			Some(
				[]int{
					1,
					2,
					3,
				},
			),
			Some(
				[]int{
					4,
					5,
					6,
				},
			),
			false,
		},
		{
			"Some vs None",
			Some([]int{1}),
			None[[]int](),
			false,
		},
		{
			"None vs None",
			None[[]int](),
			None[[]int](),
			true,
		},
	}

	eq := func(x, y []int) bool {
		return slices.Equal(x, y)
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				result := Equal(tt.a, tt.b, eq)
				if result != tt.expected {
					t.Errorf("Equal(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
				}
			},
		)
	}
}
