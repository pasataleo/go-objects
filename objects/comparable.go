package objects

type Comparable[T any] interface {
	CompareTo(other T) int
}

type Comparator[T any] interface {
	Compare(left, right T) int
}
