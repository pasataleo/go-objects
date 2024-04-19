package objects

import "strconv"

var (
	_ Object               = (*Float32)(nil)
	_ Comparable[*Float32] = (*Float32)(nil)

	_ Object               = (*Float64)(nil)
	_ Comparable[*Float64] = (*Float64)(nil)
)

type Float32 struct {
	value float32
}

func WrapFloat32(value float32) *Float32 {
	return &Float32{
		value: value,
	}
}

func (f *Float32) Unwrap() float32 {
	return f.value
}

func (f *Float32) CompareTo(other *Float32) int {
	diff := other.value - f.value
	switch {
	case diff > 0:
		return 1
	case diff < 0:
		return -1
	default:
		return 0
	}
}

func (f *Float32) Equals(other any) bool {
	if fOther, ok := other.(Float32); ok {
		return fOther.value == f.value
	}

	if fOther, ok := other.(*Float32); ok {
		return fOther.value == f.value
	}

	return false
}

func (f *Float32) HashCode() uint64 {
	return hashFromString(f)
}

func (f *Float32) String() string {
	return strconv.FormatFloat(float64(f.value), 'g', -1, 32)
}

func (f *Float32) MarshalJSON() ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (f *Float32) UnmarshalJSON(bytes []byte) error {
	//TODO implement me
	panic("implement me")
}

type Float64 struct {
	value float64
}

func (f *Float64) MarshalJSON() ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (f *Float64) UnmarshalJSON(bytes []byte) error {
	//TODO implement me
	panic("implement me")
}

func WrapFloat64(value float64) *Float64 {
	return &Float64{
		value: value,
	}
}

func (f *Float64) Unwrap() float64 {
	return f.value
}

func (f *Float64) CompareTo(other *Float64) int {
	diff := other.value - f.value
	switch {
	case diff > 0:
		return 1
	case diff < 0:
		return -1
	default:
		return 0
	}
}

func (f *Float64) Equals(other any) bool {
	if fOther, ok := other.(Float64); ok {
		return fOther.value == f.value
	}

	if fOther, ok := other.(*Float64); ok {
		return fOther.value == f.value
	}

	return false
}

func (f *Float64) HashCode() uint64 {
	return hashFromString(f)
}

func (f *Float64) String() string {
	return strconv.FormatFloat(f.value, 'g', -1, 64)
}
