package objects

import "iter"

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

func SequenceFrom[T any](iterable Iterable[T]) iter.Seq[T] {
	iterator := iterable.Iterator()
	return func(yield func(T) bool) {
		for iterator.HasNext() {
			if !yield(iterator.Next()) {
				return
			}
		}
	}
}

var _ Iterator[Object] = (*SliceIterator[Object])(nil)

type SliceIterator[T any] struct {
	current int
	slice   []T
}

func NewSliceIterator[T any](slice []T) *SliceIterator[T] {
	return &SliceIterator[T]{slice: slice}
}

func (s *SliceIterator[T]) HasNext() bool {
	return s.current < len(s.slice)
}

func (s *SliceIterator[T]) Next() T {
	item := s.slice[s.current]
	s.current++
	return item
}
