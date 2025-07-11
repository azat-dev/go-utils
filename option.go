package go_utils

// Option представляет собой значение, которое может существовать (Some) или отсутствовать (None).
type Option[T any] struct {
	value   T
	present bool
}

// Some создает Option с присутствующим значением.
func Some[T any](v T) Option[T] {
	return Option[T]{value: v, present: true}
}

// None создает пустой Option для указанного типа.
func None[T any]() Option[T] {
	return Option[T]{present: false}
}

// IsSome возвращает true, если значение присутствует.
func (o Option[T]) IsSome() bool {
	return o.present
}

// IsNone возвращает true, если значение отсутствует.
func (o Option[T]) IsNone() bool {
	return !o.present
}

// Get возвращает значение и true, если оно есть. Иначе возвращает zero-value и false.
// Это идиоматический Go-способ работы с опциональными значениями.
func (o Option[T]) Get() (T, bool) {
	if o.present {
		return o.value, true
	}
	var zero T
	return zero, false
}

// Unwrap возвращает значение, если оно есть. Иначе паникует.
// Используйте, когда вы уверены, что значение существует.
func (o Option[T]) Unwrap() T {
	if !o.present {
		panic("called Unwrap() on a None option")
	}
	return o.value
}

// UnwrapOr возвращает значение, если оно есть. Иначе возвращает значение по умолчанию.
func (o Option[T]) UnwrapOr(defaultValue T) T {
	if o.present {
		return o.value
	}
	return defaultValue
}

// Map применяет функцию к значению внутри Option, если оно есть.
// Возвращает новый Option с результатом. Если исходный Option был None, возвращает None.
func Map[T, U any](o Option[T], f func(T) U) Option[U] {
	if o.present {
		return Some(f(o.value))
	}
	return None[U]()
}

// FlatMap (или AndThen) применяет функцию, которая сама возвращает Option.
// Используется для цепочек вызовов, каждый из которых может вернуть пустое значение.
func FlatMap[T, U any](o Option[T], f func(T) Option[U]) Option[U] {
	if o.present {
		return f(o.value)
	}
	return None[U]()
}
