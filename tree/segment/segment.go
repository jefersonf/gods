package segment

import (
	"errors"
)

// Integer represents a generic interface for integer values.
type Integer interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// Float represents a generic interface for floating-point values.
type Float interface {
	float32 | float64
}

// Number represents a Numeric value.
type Number interface {
	Integer | Float
}

// SegNode is a segment tree node whose value could be anything comparable.
type SegNode[T any] struct {
	value T
}

// Segment represents a slice of SegNodes with support for efficient
// range queries like sum, minimum, maximum, and any function that combines
// two nodes into a new one called segment.
type Segment[T any] struct {
	data     []SegNode[T]
	size     uint
	joinFunc func(n1, n2 *SegNode[T]) SegNode[T]
}

var (
	// errSegNodeDoNotExist is an error that accurs whenever an invalid segnode id is accessed.
	errSegNodeDoNotExist = errors.New("segnode id does not exist")
)

// New creates a new SegmentTree with the given elements.
func New[T any](a []T, f func(n1, n2 *SegNode[T]) SegNode[T]) *Segment[T] {
	n := len(a)
	var b uint = 1
	for b < uint(n) {
		b *= 2
	}
	t := make([]SegNode[T], b*2)
	if f == nil {
		f = func(n1, n2 *SegNode[T]) SegNode[T] {
			return SegNode[T]{value: n1.value}
		}
	}
	segtree := &Segment[T]{
		data:     t,
		joinFunc: f,
		size:     uint(n),
	}
	segtree.Build(1, 1, uint(n), a)
	return segtree
}

// Build builds a segment tree with the given elements.
func (t *Segment[T]) Build(id uint, l, r uint, a []T) {
	if l == r {
		t.data[id].value = a[l-1]
		return
	}
	m := (l + r) / 2
	t.Build(id*2, l, m, a)
	t.Build(id*2+1, m+1, r, a)
	t.data[id] = t.joinFunc(&t.data[id*2], &t.data[id*2+1])
}

// Size returns the size of the segment tree.
func (t *Segment[T]) Size() uint {
	return t.size
}

// GetNodeByID returns a SegNode for a given id and an error
// when the given id is invalid.
func (t *Segment[T]) GetNodeByID(id uint) (SegNode[T], error) {
	if id == 0 || id > uint(len(t.data)) {
		return SegNode[T]{}, errSegNodeDoNotExist
	}
	return t.data[id], nil
}

// Set updates the segment tree data by its position.
func (t *Segment[T]) Set(pos int, val T) {
	b := len(t.data) / 2
	t.update(1, 1, b, pos, val)
}

// QueryRange queries the segnode from the segment tree that is
// responsible for the given range [l, r] and an error when that
// interval is invalid.
func (t *Segment[T]) QueryRange(l, r int) (SegNode[T], error) {
	if l < 1 || r > int(t.size) {
		return SegNode[T]{}, errSegNodeDoNotExist
	}
	sn := t.queryRange(1, 1, int(t.size), l, r)
	return sn, nil
}

func (t *Segment[T]) queryRange(id uint, i, j, l, r int) SegNode[T] {
	if i == l && j == r {
		return t.data[id]
	}

	m := (i + j) / 2

	if r <= m {
		return t.queryRange(id*2, i, m, l, r)
	}
	if l > m {
		return t.queryRange(id*2+1, m+1, j, l, r)
	}

	n1 := t.queryRange(id*2, i, m, l, m)
	n2 := t.queryRange(id*2+1, m+1, j, m+1, r)
	return t.joinFunc(&n1, &n2)
}

func (t *Segment[T]) update(id uint, l, r, pos int, val T) {
	if pos == l && l == r {
		t.data[id].value = val
		return
	}
	m := (l + r) / 2
	if pos <= m {
		t.update(id*2, l, m, pos, val)
	} else {
		t.update(id*2+1, m+1, r, pos, val)
	}
	t.data[id] = t.joinFunc(&t.data[id*2], &t.data[id*2+1])
}
