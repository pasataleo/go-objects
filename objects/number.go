package objects

import (
	"encoding/json"
	"strings"
)

var (
	_ Object              = (*Number)(nil)
	_ Comparable[*Number] = (*Number)(nil)
)

type Number struct {
	value json.Number
}

func WrapNumber(value json.Number) *Number {
	return &Number{
		value: value,
	}
}

func (n *Number) Unwrap() json.Number {
	return n.value
}

func (n *Number) Int64() (int64, error) {
	return n.value.Int64()
}

func (n *Number) Float64() (float64, error) {
	return n.value.Float64()
}

func (n *Number) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.value)
}

func (n *Number) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &n.value)
}

func (n *Number) Equals(other any) bool {
	if nOther, ok := other.(Number); ok {
		return nOther.value == n.value
	}

	if nOther, ok := other.(*Number); ok {
		return nOther.value == n.value
	}

	return false
}

func (n *Number) HashCode() uint64 {
	return hashFromString(n)
}

func (n *Number) String() string {
	return string(n.value)
}

func (n *Number) CompareTo(other *Number) int {
	return strings.Compare(string(n.value), string(other.value))
}
