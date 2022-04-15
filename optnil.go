package option

import (
	"fmt"
)

// Optnil represents an optional value:
// every [`Optnil`] is either [`NonNil`](which is nonnull *T), or [`Nil`](which is nil).
type Optnil[T any] struct {
	value *T
}

// String returns the string representation.
func (o Optnil[T]) String() string {
	if o.IsNil() {
		return "Nil"
	}
	return fmt.Sprintf("NonNil(%v)", o.value)
}

// Ptr wraps a value pointer.
func Ptr[T any](value *T) Optnil[T] {
	return Optnil[T]{value: value}
}

// Nil returns a none.
func Nil[T any]() Optnil[T] {
	return Optnil[T]{value: nil}
}

// NotNil returns `true` if the option has value.
func (o Optnil[T]) NotNil() bool {
	return !o.IsNil()
}

// NotNilAnd returns `true` if the option has value and the value inside of it matches a predicate.
func (o Optnil[T]) NotNilAnd(f func(*T) bool) bool {
	if o.NotNil() {
		return f(o.value)
	}
	return false
}

// IsNil returns `true` if the option is none.
func (o Optnil[T]) IsNil() bool {
	return o.value == nil
}

// Expect returns the contained [`NonNil`] value.
// Panics if the value is null with a custom panic message provided by `msg`.
func (o Optnil[T]) Expect(msg string) *T {
	if o.IsNil() {
		panic(fmt.Errorf("%s", msg))
	}
	return o.value
}

// Unwrap returns the contained value.
// Panics if the value is null.
func (o Optnil[T]) Unwrap() *T {
	if o.NotNil() {
		return o.value
	}
	var t T
	panic(fmt.Sprintf("call Optnil[%T].Unwrap() on nonnull", t))
}

// UnwrapOr returns the contained value or a provided default.
func (o Optnil[T]) UnwrapOr(defaultPtr *T) *T {
	if o.NotNil() {
		return o.value
	}
	return defaultPtr
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (o Optnil[T]) UnwrapOrElse(defaultPtr func() *T) *T {
	if o.NotNil() {
		return o.value
	}
	return defaultPtr()
}

// UnwrapUnchecked returns the contained value.
func (o Optnil[T]) UnwrapUnchecked() *T {
	return o.value
}

// Map maps an `Optnil[T]` to `Optnil[T]` by applying a function to a contained value.
func (o Optnil[T]) Map(f func(*T) *T) Optnil[T] {
	if o.NotNil() {
		return Ptr[T](f(o.value))
	}
	return Nil[T]()
}

// OptnilMap maps an `Optnil[T]` to `Optnil[U]` by applying a function to a contained value.
func OptnilMap[T any, U any](o Optnil[T], f func(*T) *U) Optnil[U] {
	if o.NotNil() {
		return Ptr[U](f(o.value))
	}
	return Nil[U]()
}

// Inspect calls the provided closure with a reference to the contained value (if it has value).
func (o Optnil[T]) Inspect(f func(*T)) Optnil[T] {
	if o.NotNil() {
		f(o.value)
	}
	return o
}

// MapOr returns the provided default value (if none),
// or applies a function to the contained value (if any).
func (o Optnil[T]) MapOr(defaultPtr *T, f func(*T) *T) *T {
	if o.NotNil() {
		return f(o.value)
	}
	return defaultPtr
}

// OptnilMapOr returns the provided default value (if none),
// or applies a function to the contained value (if any).
func OptnilMapOr[T any, U any](o Optnil[T], defaultPtr *U, f func(*T) *U) *U {
	if o.NotNil() {
		return f(o.value)
	}
	return defaultPtr
}

// MapOrElse computes a default function value (if none), or
// applies a different function to the contained value (if any).
func (o Optnil[T]) MapOrElse(defaultFn func() *T, f func(*T) *T) *T {
	if o.NotNil() {
		return f(o.value)
	}
	return defaultFn()
}

// OptnilMapOrElse computes a default function value (if none), or
// applies a different function to the contained value (if any).
func OptnilMapOrElse[T any, U any](o Optnil[T], defaultFn func() *U, f func(*T) *U) *U {
	if o.NotNil() {
		return f(o.value)
	}
	return defaultFn()
}

// And returns [`Nil`] if the option is [`Nil`], otherwise returns `optb`.
func (o Optnil[T]) And(optb Optnil[T]) Optnil[T] {
	if o.NotNil() {
		return optb
	}
	return o
}

// OptnilAnd returns [`Nil`] if the option is [`Nil`], otherwise returns `optb`.
func OptnilAnd[T any, U any](o Optnil[T], optb Optnil[U]) Optnil[U] {
	if o.NotNil() {
		return optb
	}
	return Nil[U]()
}

// AndThen returns [`Nil`] if the option is [`Nil`], otherwise calls `f` with the
func (o Optnil[T]) AndThen(f func(*T) Optnil[T]) Optnil[T] {
	if o.IsNil() {
		return o
	}
	return f(o.value)
}

// OptnilAndThen returns [`Nil`] if the option is [`Nil`], otherwise calls `f` with the
func OptnilAndThen[T any, U any](o Optnil[T], f func(*T) Optnil[U]) Optnil[U] {
	if o.IsNil() {
		return Nil[U]()
	}
	return f(o.value)
}

// Filter returns [`Nil`] if the option is [`Nil`], otherwise calls `predicate`
// with the wrapped value and returns.
func (o Optnil[T]) Filter(predicate func(*T) bool) Optnil[T] {
	if o.NotNil() {
		if predicate(o.value) {
			return o
		}
	}
	return Nil[T]()
}

// Or returns the option if it contains a value, otherwise returns `optb`.
func (o Optnil[T]) Or(optb Optnil[T]) Optnil[T] {
	if o.IsNil() {
		return optb
	}
	return o
}

// OrElse returns [`Nil`] if the option is [`Nil`], otherwise calls `f` with the returns the result.
func (o Optnil[T]) OrElse(f func() Optnil[T]) Optnil[T] {
	if o.IsNil() {
		return f()
	}
	return o
}

// XorElse [`NonNil`] if exactly one of `self`, `optb` is [`NonNil`], otherwise returns [`Nil`].
func (o Optnil[T]) XorElse(optb Optnil[T]) Optnil[T] {
	if o.NotNil() && optb.IsNil() {
		return o
	}
	if o.IsNil() && optb.NotNil() {
		return optb
	}
	return Nil[T]()
}

// Insert inserts `value` into the option, then returns a reference to it.
func (o *Optnil[T]) Insert(some *T) *T {
	o.value = some
	return o.value
}

// GetOrInsert inserts `value` into the option if it is [`Nil`], then
// returns a reference to the contained value.
func (o *Optnil[T]) GetOrInsert(some *T) *T {
	if o.IsNil() {
		o.value = some
	}
	return o.value
}

// GetOrInsertWith inserts a value computed from `f` into the option if it is [`Nil`],
// then returns a mutable reference to the contained value.
func (o *Optnil[T]) GetOrInsertWith(f func() *T) *T {
	if o.IsNil() {
		o.value = f()
	}
	return o.value
}

// Replace replaces the actual value in the option by the value given in parameter,
// returning the old value if present,
// leaving a [`NonNil`] in its place without deinitializing either one.
func (o *Optnil[T]) Replace(some *T) *Optnil[T] {
	o.value = some
	return o
}

// OptnilContains returns `true` if the option is a [`NonNil`] value containing the given value.
func OptnilContains[T comparable](o Optnil[T], x *T) bool {
	return o.value == x
}

// OptnilZipWith zips `value` and another `Optnil` with function `f`.
//
// If `value` is `Ptr(s)` and `other` is `Ptr(o)`, this method returns `Ptr(f(s, o))`.
// Otherwise, `Nil` is returned.
func OptnilZipWith[T any, U any, R any](some Optnil[T], other Optnil[U], f func(*T, *U) *R) Optnil[R] {
	if some.NotNil() && other.NotNil() {
		return Ptr(f(some.value, other.value))
	}
	return Nil[R]()
}
