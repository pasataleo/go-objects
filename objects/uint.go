package objects

import (
	"encoding/json"
	"strconv"
)

var (
	_ Object            = (*UInt)(nil)
	_ Comparable[*UInt] = (*UInt)(nil)

	_ Object             = (*UInt8)(nil)
	_ Comparable[*UInt8] = (*UInt8)(nil)

	_ Object              = (*UInt16)(nil)
	_ Comparable[*UInt16] = (*UInt16)(nil)

	_ Object              = (*UInt32)(nil)
	_ Comparable[*UInt32] = (*UInt32)(nil)

	_ Object              = (*UInt64)(nil)
	_ Comparable[*UInt64] = (*UInt64)(nil)
)

type UInt struct {
	value uint
}

func WrapUInt(value uint) *UInt {
	return &UInt{
		value: value,
	}
}

func (i *UInt) Unwrap() uint {
	return i.value
}

func (i *UInt) CompareTo(other *UInt) int {
	switch {
	case i.value < other.value:
		return -1
	case other.value < i.value:
		return 1
	default:
		return 0
	}
}

func (i *UInt) Equals(other any) bool {
	if iOther, ok := other.(UInt); ok {
		return iOther.value == i.value
	}

	if iOther, ok := other.(*UInt); ok {
		return iOther.value == i.value
	}

	return false
}

func (i *UInt) HashCode() uint64 {
	return uint64(i.value)
}

func (i *UInt) String() string {
	return strconv.FormatUint(uint64(i.value), 10)
}

func (i *UInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.value)
}

func (i *UInt) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &i.value)
}

type UInt8 struct {
	value uint8
}

func WrapUInt8(value uint8) *UInt8 {
	return &UInt8{
		value: value,
	}
}

func (i *UInt8) Unwrap() uint8 {
	return i.value
}

func (i *UInt8) CompareTo(other *UInt8) int {
	switch {
	case i.value < other.value:
		return -1
	case other.value < i.value:
		return 1
	default:
		return 0
	}
}

func (i *UInt8) Equals(other any) bool {
	if iOther, ok := other.(UInt8); ok {
		return iOther.value == i.value
	}

	if iOther, ok := other.(*UInt8); ok {
		return iOther.value == i.value
	}

	return false
}

func (i *UInt8) HashCode() uint64 {
	return uint64(i.value)
}

func (i *UInt8) String() string {
	return strconv.FormatUint(uint64(i.value), 10)
}

func (i *UInt8) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.value)
}

func (i *UInt8) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &i.value)
}

type UInt16 struct {
	value uint16
}

func WrapUInt16(value uint16) *UInt16 {
	return &UInt16{
		value: value,
	}
}

func (i *UInt16) Unwrap() uint16 {
	return i.value
}

func (i *UInt16) CompareTo(other *UInt16) int {
	switch {
	case i.value < other.value:
		return -1
	case other.value < i.value:
		return 1
	default:
		return 0
	}
}

func (i *UInt16) Equals(other any) bool {
	if iOther, ok := other.(UInt16); ok {
		return iOther.value == i.value
	}

	if iOther, ok := other.(*UInt16); ok {
		return iOther.value == i.value
	}

	return false
}

func (i *UInt16) HashCode() uint64 {
	return uint64(i.value)
}

func (i *UInt16) String() string {
	return strconv.FormatUint(uint64(i.value), 10)
}

func (i *UInt16) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.value)
}

func (i *UInt16) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &i.value)
}

type UInt32 struct {
	value uint32
}

func WrapUInt32(value uint32) *UInt32 {
	return &UInt32{
		value: value,
	}
}

func (i *UInt32) Unwrap() uint32 {
	return i.value
}

func (i *UInt32) CompareTo(other *UInt32) int {
	switch {
	case i.value < other.value:
		return -1
	case other.value < i.value:
		return 1
	default:
		return 0
	}
}

func (i *UInt32) Equals(other any) bool {
	if iOther, ok := other.(UInt32); ok {
		return iOther.value == i.value
	}

	if iOther, ok := other.(*UInt32); ok {
		return iOther.value == i.value
	}

	return false
}

func (i *UInt32) HashCode() uint64 {
	return uint64(i.value)
}

func (i *UInt32) String() string {
	return strconv.FormatUint(uint64(i.value), 10)
}

func (i *UInt32) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.value)
}

func (i *UInt32) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &i.value)
}

type UInt64 struct {
	value uint64
}

func WrapUInt64(value uint64) *UInt64 {
	return &UInt64{
		value: value,
	}
}

func (i *UInt64) Unwrap() uint64 {
	return i.value
}

func (i *UInt64) CompareTo(other *UInt64) int {
	switch {
	case i.value < other.value:
		return -1
	case other.value < i.value:
		return 1
	default:
		return 0
	}
}

func (i *UInt64) Equals(other any) bool {
	if iOther, ok := other.(UInt64); ok {
		return iOther.value == i.value
	}

	if iOther, ok := other.(*UInt64); ok {
		return iOther.value == i.value
	}

	return false
}

func (i *UInt64) HashCode() uint64 {
	return i.value
}

func (i *UInt64) String() string {
	return strconv.FormatUint(i.value, 10)
}

func (i *UInt64) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.value)
}

func (i *UInt64) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &i.value)
}
