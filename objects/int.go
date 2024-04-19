package objects

import (
	"encoding/json"
	"strconv"
)

var (
	_ Object           = (*Int)(nil)
	_ Comparable[*Int] = (*Int)(nil)

	_ Object            = (*Int8)(nil)
	_ Comparable[*Int8] = (*Int8)(nil)

	_ Object             = (*Int16)(nil)
	_ Comparable[*Int16] = (*Int16)(nil)

	_ Object             = (*Int32)(nil)
	_ Comparable[*Int32] = (*Int32)(nil)

	_ Object             = (*Int64)(nil)
	_ Comparable[*Int64] = (*Int64)(nil)
)

type Int struct {
	value int
}

func WrapInt(value int) *Int {
	return &Int{
		value: value,
	}
}

func (i *Int) Unwrap() int {
	return i.value
}

func (i *Int) CompareTo(other *Int) int {
	return other.value - i.value
}

func (i *Int) Equals(other any) bool {
	if iOther, ok := other.(Int); ok {
		return iOther.value == i.value
	}

	if iOther, ok := other.(*Int); ok {
		return iOther.value == i.value
	}

	return false
}

func (i *Int) HashCode() uint64 {
	return uint64(i.value)
}

func (i *Int) String() string {
	return strconv.FormatInt(int64(i.value), 10)
}

func (i *Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.value)
}

func (i *Int) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &i.value)
}

type Int8 struct {
	value int8
}

func WrapInt8(value int8) *Int8 {
	return &Int8{
		value: value,
	}
}

func (i *Int8) Unwrap() int8 {
	return i.value
}

func (i *Int8) CompareTo(other *Int8) int {
	return int(other.value - i.value)
}

func (i *Int8) Equals(other any) bool {
	if iOther, ok := other.(Int8); ok {
		return iOther.value == i.value
	}

	if iOther, ok := other.(*Int8); ok {
		return iOther.value == i.value
	}

	return false
}

func (i *Int8) HashCode() uint64 {
	return uint64(i.value)
}

func (i *Int8) String() string {
	return strconv.FormatInt(int64(i.value), 10)
}

func (i *Int8) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.value)
}

func (i *Int8) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &i.value)
}

type Int16 struct {
	value int16
}

func WrapInt16(value int16) *Int16 {
	return &Int16{
		value: value,
	}
}

func (i *Int16) Unwrap() int16 {
	return i.value
}

func (i *Int16) CompareTo(other *Int16) int {
	return int(other.value - i.value)
}

func (i *Int16) Equals(other any) bool {
	if iOther, ok := other.(Int16); ok {
		return iOther.value == i.value
	}

	if iOther, ok := other.(*Int16); ok {
		return iOther.value == i.value
	}

	return false
}

func (i *Int16) HashCode() uint64 {
	return uint64(i.value)
}

func (i *Int16) String() string {
	return strconv.FormatInt(int64(i.value), 10)
}

func (i *Int16) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.value)
}

func (i *Int16) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &i.value)
}

type Int32 struct {
	value int32
}

func WrapInt32(value int32) *Int32 {
	return &Int32{
		value: value,
	}
}

func (i *Int32) Unwrap() int32 {
	return i.value
}

func (i *Int32) CompareTo(other *Int32) int {
	return int(other.value - i.value)
}

func (i *Int32) Equals(other any) bool {
	if iOther, ok := other.(Int32); ok {
		return iOther.value == i.value
	}

	if iOther, ok := other.(*Int32); ok {
		return iOther.value == i.value
	}

	return false
}

func (i *Int32) HashCode() uint64 {
	return uint64(i.value)
}

func (i *Int32) String() string {
	return strconv.FormatInt(int64(i.value), 10)
}

func (i *Int32) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.value)
}

func (i *Int32) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &i.value)
}

type Int64 struct {
	value int64
}

func WrapInt64(value int64) *Int64 {
	return &Int64{
		value: value,
	}
}

func (i *Int64) Unwrap() int64 {
	return i.value
}

func (i *Int64) CompareTo(other *Int64) int {
	return int(other.value - i.value)
}

func (i *Int64) Equals(other any) bool {
	if iOther, ok := other.(Int64); ok {
		return iOther.value == i.value
	}

	if iOther, ok := other.(*Int64); ok {
		return iOther.value == i.value
	}

	return false
}

func (i *Int64) HashCode() uint64 {
	return uint64(i.value)
}

func (i *Int64) String() string {
	return strconv.FormatInt(i.value, 10)
}

func (i *Int64) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.value)
}

func (i *Int64) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &i.value)
}
