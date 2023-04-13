package objects

type Iterable[T any] interface {
	Iterator() Iterator[T]
}

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}

func SliceFrom[T any](iterable Iterable[T]) []T {
	iterator := iterable.Iterator()

	var slice []T
	for iterator.HasNext() {
		slice = append(slice, iterator.Next())
	}
	return slice
}
