package objects

type Comparable[O Object] interface {
	CompareTo(other O) int
}

type Comparator[O Object] interface {
	Compare(left, right O) int
}
