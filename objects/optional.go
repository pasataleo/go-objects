package objects

import (
	"encoding/json"
	"fmt"

	"github.com/pasataleo/go-errors/errors"
)

type Optional[V any] interface {
	Object

	Get() V
	GetSafe() (V, error)

	IsEmpty() bool
}

func EmptyOptional[V Object]() Optional[V] {
	return &optional[V]{
		Empty: true,
	}
}

func OptionalOf[V Object](value V) Optional[V] {
	return &optional[V]{
		Value: value,
	}
}

type optional[V Object] struct {
	Empty bool `json:"empty"`
	Value V    `json:"value,omitempty"`
}

func (optional *optional[V]) MarshalJSON() ([]byte, error) {
	return json.Marshal(optional)
}

func (optional *optional[V]) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, optional)
}

func (optional *optional[V]) Equals(other any) bool {
	if oOptional, ok := other.(Optional[V]); ok {
		if oOptional.IsEmpty() != optional.IsEmpty() {
			return false
		}

		if optional.Empty {
			// If neither optional has values, then they are equal.
			return true
		}

		return optional.Get().Equals(oOptional.Get())
	}
	return false
}

func (optional *optional[V]) HashCode() uint64 {
	if optional.Empty {
		return 0
	}

	return optional.Value.HashCode()
}

func (optional *optional[V]) String() string {
	if optional.Empty {
		return "optional(empty)"
	}

	return fmt.Sprintf("optional(%s)", optional.Value.String())
}

func (optional *optional[V]) Get() V {
	if optional.Empty {
		panic("called Get() on empty Optional")
	}

	return optional.Value
}

func (optional *optional[V]) GetSafe() (V, error) {
	if optional.Empty {
		return optional.Value, errors.New(nil, ErrorCodeNotPresent, "empty")
	}

	return optional.Value, nil
}

func (optional *optional[V]) IsEmpty() bool {
	return optional.Empty
}
