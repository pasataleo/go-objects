package objects

import "fmt"

type EqualsFn[O any] func(value O, other any) bool
type HashCodeFn[O any] func(value O) uint64
type ToStringFn[O any] func(value O) string

type Object interface {
	Equals(other any) bool
	HashCode() uint64
	ToString() string
}

var _ Object = (*AbstractObject)(nil)

type AbstractObject struct{}

func (obj *AbstractObject) Equals(other any) bool {
	return memaddr(obj) == memaddr(other)
}

func (obj *AbstractObject) HashCode() uint64 {
	return uint64(memaddr(obj))
}

func (obj *AbstractObject) ToString() string {
	return fmt.Sprintf("%d", memaddr(obj))
}

type ObjectConverter[O any] struct {
	Equals   EqualsFn[O]
	HashCode HashCodeFn[O]
	ToString ToStringFn[O]
}

func ObjectIdentityConverter[O Object]() ObjectConverter[O] {
	return ObjectConverter[O]{
		Equals: func(value O, other any) bool {
			return value.Equals(other)
		},
		HashCode: func(value O) uint64 {
			return value.HashCode()
		},
		ToString: func(value O) string {
			return value.ToString()
		},
	}
}

func DefaultObjectConverter[O any]() ObjectConverter[O] {
	return ObjectConverter[O]{
		Equals: func(value O, other any) bool {
			return memaddr(value) == memaddr(other)
		},
		HashCode: func(value O) uint64 {
			return uint64(memaddr(value))
		},
		ToString: func(value O) string {
			return fmt.Sprintf("%d", memaddr(value))
		},
	}
}

func (c ObjectConverter[O]) SetEquals(equals EqualsFn[O]) ObjectConverter[O] {
	c.Equals = equals
	return c
}

func (c ObjectConverter[O]) SetHashCode(hashCode HashCodeFn[O]) ObjectConverter[O] {
	c.HashCode = hashCode
	return c
}

func (c ObjectConverter[O]) SetToString(toString ToStringFn[O]) ObjectConverter[O] {
	c.ToString = toString
	return c
}
