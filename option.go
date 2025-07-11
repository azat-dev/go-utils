package go_utils

// Optional представляет собой значение, которое может существовать (Some) или отсутствовать (None).
type Optional[T any] struct {
	value   T
	present bool
}

// Some создает Optional с присутствующим значением.
func Some[T any](v T) Optional[T] {
	return Optional[T]{value: v, present: true}
}

// None создает пустой Optional для указанного типа.
func None[T any]() Optional[T] {
	return Optional[T]{present: false}
}

// IsSome возвращает true, если значение присутствует.
func (o Optional[T]) IsSome() bool {
	return o.present
}

// IsNone возвращает true, если значение отсутствует.
func (o Optional[T]) IsNone() bool {
	return !o.present
}

// Get возвращает значение и true, если оно есть. Иначе возвращает zero-value и false.
// Это идиоматический Go-способ работы с опциональными значениями.
func (o Optional[T]) Get() (T, bool) {
	if o.present {
		return o.value, true
	}
	var zero T
	return zero, false
}

// Unwrap возвращает значение, если оно есть. Иначе паникует.
// Используйте, когда вы уверены, что значение существует.
func (o Optional[T]) Unwrap() T {
	if !o.present {
		panic("called Unwrap() on a None option")
	}
	return o.value
}

// UnwrapOr возвращает значение, если оно есть. Иначе возвращает значение по умолчанию.
func (o Optional[T]) UnwrapOr(defaultValue T) T {
	if o.present {
		return o.value
	}
	return defaultValue
}

// Map применяет функцию к значению внутри Optional, если оно есть.
// Возвращает новый Optional с результатом. Если исходный Optional был None, возвращает None.
func Map[T, U any](o Optional[T], f func(T) U) Optional[U] {
	if o.present {
		return Some(f(o.value))
	}
	return None[U]()
}

// FlatMap (или AndThen) применяет функцию, которая сама возвращает Optional.
// Используется для цепочек вызовов, каждый из которых может вернуть пустое значение.
func FlatMap[T, U any](o Optional[T], f func(T) Optional[U]) Optional[U] {
	if o.present {
		return f(o.value)
	}
	return None[U]()
}
