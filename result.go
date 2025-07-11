package go_utils

// Result представляет собой результат операции, который может быть либо успешным значением (Ok), либо ошибкой (Err).
// T - тип успешного значения.
type Result[T any] struct {
	value T
	err   error
}

// Ok создает Result с успешным значением.
// Используйте для обозначения успешного выполнения операции.
func Ok[T any](v T) Result[T] {
	return Result[T]{value: v, err: nil}
}

// Err создает Result с ошибкой.
// Используйте для обозначения неудачного выполнения операции.
func Err[T any](e error) Result[T] {
	var zero T // Нулевое значение для типа T
	return Result[T]{value: zero, err: e}
}

// IsOk возвращает true, если Result содержит успешное значение.
func (r Result[T]) IsOk() bool {
	return r.err == nil
}

// IsErr возвращает true, если Result содержит ошибку.
func (r Result[T]) IsErr() bool {
	return r.err != nil
}

// Get возвращает успешное значение и nil, если оно есть. Иначе возвращает zero-value и ошибку.
// Это идиоматический Go-способ работы с Result.
func (r Result[T]) Get() (T, error) {
	return r.value, r.err
}

// Unwrap возвращает успешное значение, если оно есть. Иначе паникует с ошибкой.
// Используйте, когда вы уверены, что операция была успешной.
func (r Result[T]) Unwrap() T {
	if r.err != nil {
		panic("called Unwrap() on an Err Result: " + r.err.Error())
	}
	return r.value
}

// UnwrapOr возвращает успешное значение, если оно есть. Иначе возвращает значение по умолчанию.
func (r Result[T]) UnwrapOr(defaultValue T) T {
	if r.err == nil {
		return r.value
	}
	return defaultValue
}

// UnwrapErr возвращает ошибку, если она есть. Иначе паникует.
// Используйте, когда вы уверены, что Result содержит ошибку.
func (r Result[T]) UnwrapErr() error {
	if r.err == nil {
		panic("called UnwrapErr() on an Ok Result")
	}
	return r.err
}

// MustGet возвращает успешное значение. Если есть ошибка, паникует.
// Удобно для инициализации или тестового кода, где ошибки не ожидаются.
func (r Result[T]) MustGet() T {
	if r.err != nil {
		panic("called MustGet() on an Err Result: " + r.err.Error())
	}
	return r.value
}

// MapResult применяет функцию к успешному значению внутри Result, если оно есть.
// Возвращает новый Result с результатом. Если исходный Result был Err, возвращает тот же Err.
// F - функция, которая преобразует успешное значение.
func MapResult[T, U any](r Result[T], f func(T) U) Result[U] {
	if r.IsOk() {
		return Ok(f(r.value))
	}
	return Err[U](r.err) // Прокидываем ошибку дальше
}

// FlatMapResult (или AndThen) применяет функцию, которая сама возвращает Result.
// Используется для цепочек вызовов, каждый из которых может вернуть ошибку.
// F - функция, которая принимает успешное значение и возвращает новый Result.
func FlatMapResult[T, U any](r Result[T], f func(T) Result[U]) Result[U] {
	if r.IsOk() {
		return f(r.value)
	}
	return Err[U](r.err) // Прокидываем ошибку дальше
}

// MapErr применяет функцию к ошибке внутри Result, если она есть.
// Позволяет преобразовывать или оборачивать ошибки.
func MapErr[T any](r Result[T], f func(error) error) Result[T] {
	if r.IsErr() {
		return Err[T](f(r.err))
	}
	return r // Возвращаем Ok Result без изменений
}

// Inspect выполняет действие (функцию) с успешным значением, если оно присутствует,
// без изменения самого Result. Полезно для логирования или отладки.
func (r Result[T]) Inspect(f func(T)) Result[T] {
	if r.IsOk() {
		f(r.value)
	}
	return r
}

// InspectErr выполняет действие (функцию) с ошибкой, если она присутствует,
// без изменения самого Result. Полезно для логирования ошибок.
func (r Result[T]) InspectErr(f func(error)) Result[T] {
	if r.IsErr() {
		f(r.err)
	}
	return r
}

// OrElse возвращает текущий Result, если он Ok. Иначе возвращает другой Result.
// Полезно для предоставления запасного варианта в случае ошибки.
func (r Result[T]) OrElse(alternative Result[T]) Result[T] {
	if r.IsOk() {
		return r
	}
	return alternative
}

// OrElseDo возвращает текущий Result, если он Ok. Иначе вызывает функцию-поставщик для получения другого Result.
func (r Result[T]) OrElseDo(f func(error) Result[T]) Result[T] {
	if r.IsOk() {
		return r
	}
	return f(r.err)
}

// ToOptional преобразует Result в Optional.
// Если Result Ok, Optional будет Some. Если Result Err, Optional будет None.
func (r Result[T]) ToOptional() Optional[T] {
	if r.IsOk() {
		return Some(r.value)
	}
	return None[T]()
}

// FromOptional преобразует Optional в Result.
// Если Optional Some, Result будет Ok. Если Optional None, Result будет Err с указанной ошибкой.
func FromOptional[T any](o Optional[T], errIfNone error) Result[T] {
	if o.IsSome() {
		return Ok(o.Unwrap())
	}
	return Err[T](errIfNone)
}
