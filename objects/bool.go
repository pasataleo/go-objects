package objects

import "fmt"

var _ Object = (*Bool)(nil)
var _ Comparable[Bool] = (*Bool)(nil)

type Bool struct {
	value bool
}

func WrapBool(value bool) Bool {
	return Bool{
		value: value,
	}
}

func (b Bool) Unwrap() bool {
	return b.value
}

func (b Bool) Equals(other any) bool {
	if bOther, ok := other.(Bool); ok {
		return bOther.value == b.value
	}
	if bOther, ok := other.(*Bool); ok {
		return bOther.value == b.value
	}
	return false
}

func (b Bool) HashCode() uint64 {
	if b.value {
		return 1
	}
	return 0
}

func (b Bool) ToString() string {
	return fmt.Sprintf("%t", b.value)
}

func (b Bool) CompareTo(other Bool) int {
	var this, o int
	if b.value {
		this = 1
	}

	if other.value {
		o = 1
	}

	return o - this
}
