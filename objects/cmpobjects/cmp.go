package cmpobjects

import (
	"github.com/google/go-cmp/cmp"

	"github.com/pasataleo/go-objects/objects"
)

// ObjectEquals returns a cmp.Option that compares objects.Object instances using their Equals method.
func ObjectEquals() cmp.Option {
	return cmp.Comparer(func(a, b objects.Object) bool {
		return a.Equals(b)
	})
}

func IterableTransform[T any]() cmp.Option {
	return cmp.Transformer("cmpobjects.IterableTransform", func(i objects.Iterable[T]) []T {
		iter := i.Iterator()
		var items []T
		for iter.HasNext() {
			items = append(items, iter.Next())
		}
		return items
	})
}
