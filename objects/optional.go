package objects

import (
	"fmt"

	"github.com/pasataleo/go-errors/errors"
)

type Optional[V any] interface {
	Object

	Get() V
	GetSafe() (V, error)

	Empty() bool
}

func EmptyOptional[V Object]() Optional[V] {
	return EmptyOptionalT(ObjectIdentityConverter[V]())
}

func EmptyOptionalT[V any](converter ObjectConverter[V]) Optional[V] {
	return optional[V]{
		empty:     true,
		converter: converter,
	}
}

func OptionalOf[V Object](value V) Optional[V] {
	return OptionalOfT(value, ObjectIdentityConverter[V]())
}

func OptionalOfT[V Object](value V, converter ObjectConverter[V]) Optional[V] {
	return optional[V]{
		empty:     false,
		value:     value,
		converter: converter,
	}
}

type optional[V any] struct {
	empty bool
	value V

	converter ObjectConverter[V]
}

func (optional optional[V]) Equals(other any) bool {
	if oOptional, ok := other.(Optional[V]); ok {
		if oOptional.Empty() != optional.Empty() {
			return false
		}

		if optional.empty {
			// If neither optional has values, then they are equal.
			return true
		}

		return optional.converter.Equals(optional.Get(), oOptional.Get())
	}
	return false
}

func (optional optional[V]) HashCode() uint64 {
	if optional.empty {
		return 0
	}

	return optional.converter.HashCode(optional.value)
}

func (optional optional[V]) ToString() string {
	if optional.empty {
		return "optional(empty)"
	}

	return fmt.Sprintf("optional(%s)", optional.converter.ToString(optional.value))
}

func (optional optional[V]) Get() V {
	if optional.empty {
		panic("called Get() on empty Optional")
	}

	return optional.value
}

func (optional optional[V]) GetSafe() (V, error) {
	if optional.empty {
		return optional.value, errors.New(nil, ErrorCodeNotPresent, "empty")
	}

	return optional.value, nil
}

func (optional optional[V]) Empty() bool {
	return optional.empty
}
