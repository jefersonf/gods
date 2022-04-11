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
type SegNode[T any] struct {
	value T
}

type Segment[T any] struct {
	data     []SegNode[T]
	size     uint
	joinFunc func(n1, n2 *SegNode[T]) SegNode[T]
}

var (
	errSegNodeDoNotExist = errors.New("segnode id does not exist")
)

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

func (t *Segment[T]) Size() uint {
	return t.size
}

func (t *Segment[T]) GetNodeByID(id uint) (SegNode[T], error) {
	if id == 0 || id > uint(len(t.data)) {
		return SegNode[T]{}, errSegNodeDoNotExist
	}
	return t.data[id], nil
}

func (t *Segment[T]) Set(pos int, val T) {
	b := len(t.data) / 2
	t.update(1, 1, b, pos, val)
}

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
