package objects

var (
	_ Object             = (*String)(nil)
	_ Comparable[String] = (*String)(nil)
)

type String struct {
	value string
}

func WrapString(value string) String {
	return String{
		value: value,
	}
}

func (s String) Unwrap() string {
	return s.value
}

func (s String) CompareTo(other String) int {
	switch {
	case s.value < other.value:
		return -1
	case s.value > other.value:
		return 1
	default:
		return 0
	}
}

func (s String) Equals(other any) bool {
	if sOther, ok := other.(String); ok {
		return sOther.value == s.value
	}

	if sOther, ok := other.(*String); ok {
		return sOther.value == s.value
	}

	return false
}

func (s String) HashCode() uint64 {
	return hashFromString(s)
}

func (s String) ToString() string {
	return s.value
}
