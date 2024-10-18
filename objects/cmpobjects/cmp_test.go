package cmpobjects

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/pasataleo/go-objects/objects"
)

var (
	_ objects.Iterable[objects.Object] = List{}
)

// List is a simple list implementation.
type List struct {
	items []objects.Object
}

func (l List) Iterator() objects.Iterator[objects.Object] {
	return objects.NewSliceIterator(l.items)
}

func TestObjectEquals(t *testing.T) {
	a := objects.WrapString("a")
	b := objects.WrapString("b")

	diff := cmp.Diff(a, b, ObjectEquals())
	t.Logf("diff: %v", diff)
	if diff == "" {
		t.Errorf("expected diff, got empty")
	}

	if !cmp.Equal(a, a, ObjectEquals()) {
		t.Errorf("expected equal, got not equal")
	}
}

func TestIterableTransform(t *testing.T) {
	a := List{
		items: []objects.Object{
			objects.WrapString("a"),
			objects.WrapString("b"),
			objects.WrapString("c"),
		},
	}

	b := List{
		items: []objects.Object{
			objects.WrapString("d"),
			objects.WrapString("e"),
			objects.WrapString("f"),
		},
	}

	diff := cmp.Diff(a, b, IterableTransform[objects.Object](), ObjectEquals())
	t.Logf("diff: %v", diff)
	if diff == "" {
		t.Errorf("expected diff, got empty")
	}

	if !cmp.Equal(a, a, IterableTransform[objects.Object](), ObjectEquals()) {
		t.Errorf("expected equal, got not equal")
	}
}
