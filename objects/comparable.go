package objects

// Comparable is an interface for objects that can be compared.
type Comparable[T any] interface {
	CompareTo(other T) int
}

// ComparableObject is an interface for objects that can be compared.
type ComparableObject[T any] interface {
	Object
	Comparable[T]
}

// Comparator is an interface for objects that can compare two objects.
type Comparator[T any] interface {
	Compare(left, right T) int
}

// ReverseComparator returns a comparator that reverses the order of the given comparator.
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

// ComparableComparator returns a comparator that uses the CompareTo method of the given comparable object.
func ComparableComparator[T Comparable[T]]() Comparator[T] {
	return comparableComparator[T]{}
}

type comparableComparator[T Comparable[T]] struct{}

func (c comparableComparator[T]) Compare(left, right T) int {
	return left.CompareTo(right)
}
