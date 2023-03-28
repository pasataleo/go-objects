package objects

import "fmt"

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
