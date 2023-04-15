package objects

type Iterable[T any] interface {
	Iterator() Iterator[T]
}

type IterableObject[T any] interface {
	Object
	Iterable[T]
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
