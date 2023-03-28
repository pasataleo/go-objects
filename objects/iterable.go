package objects

type Iterable[O Object] interface {
	Iterator() Iterator[O]
}

type Iterator[O Object] interface {
	HasNext() bool
	Next() O
}

func Slice[O Object](iterable Iterable[O]) []O {
	iterator := iterable.Iterator()

	var slice []O
	for iterator.HasNext() {
		slice = append(slice, iterator.Next())
	}
	return slice
}
