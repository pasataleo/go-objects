package objects

type Comparable[T any] interface {
	CompareTo(other T) int
}

type ComparableObject[T any] interface {
	Object
	Comparable[T]
}

type Comparator[T any] interface {
	Compare(left, right T) int
}

func ReverseComparator[T any](comparator Comparator[T]) Comparator[T] {
	return reversedComparator[T]{
		comparator: comparator,
	}
}

type reversedComparator[T any] struct {
	comparator Comparator[T]
}

func (c reversedComparator[T]) Compare(left, right T) int {
	return -c.comparator.Compare(left, right)
}

func ComparableComparator[T Comparable[T]]() Comparator[T] {
	return comparableComparator[T]{}
}

type comparableComparator[T Comparable[T]] struct{}

func (c comparableComparator[T]) Compare(left, right T) int {
	return left.CompareTo(right)
}
