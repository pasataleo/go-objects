package objects

import (
	"encoding/json"
	"fmt"
)

type Object interface {
	json.Marshaler
	json.Unmarshaler

	Equals(other any) bool
	HashCode() uint64
	String() string
}

var _ Object = (*AbstractObject)(nil)

type AbstractObject struct{}

func (obj *AbstractObject) Equals(other any) bool {
	return memaddr(obj) == memaddr(other)
}

func (obj *AbstractObject) HashCode() uint64 {
	return uint64(memaddr(obj))
}

func (obj *AbstractObject) String() string {
	return fmt.Sprintf("%d", memaddr(obj))
}

func (obj *AbstractObject) MarshalJSON() ([]byte, error) {
	// This must be implemented by the concrete type.
	panic("not implemented")
}

func (obj *AbstractObject) UnmarshalJSON([]byte) error {
	// This must be implemented by the concrete type.
	panic("not implemented")
}
