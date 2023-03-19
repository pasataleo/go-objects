package objects

import "fmt"

type Object interface {
	Equals(other any) bool
	HashCode() int64
	ToString() string
}

var _ Object = (*AbstractObject)(nil)

type AbstractObject struct{}

func (obj *AbstractObject) Equals(other any) bool {
	return memaddr(obj) == memaddr(other)
}

func (obj *AbstractObject) HashCode() int64 {
	return int64(memaddr(obj))
}

func (obj *AbstractObject) ToString() string {
	return fmt.Sprintf("%d", memaddr(obj))
}
