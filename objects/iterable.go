package objects

type Iterable[T any] interface {
	Iterator() Iterator[T]
}

type Iterator[T any] interface {
	HasNext() bool
	Next() (T, error)
}

func SliceOf[T any](iterable Iterable[T]) []T {
	iterator := iterable.Iterator()

	var slice []T
	for iterator.HasNext() {
		value, err := iterator.Next()
		if err != nil {
			// Assuming the implementation of iterator is correct, this
			// shouldn't happen, so we'll panic.
			panic(err)
		}

		slice = append(slice, value)
	}
	return slice
}
