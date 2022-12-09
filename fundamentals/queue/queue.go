package queue

import "errors"

// Integer represents a generic interface for integer values.
type Integer interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// Float represents a generic interface for floating-point values.
type Float interface {
	float32 | float64
}

// Number represents a numeric value.
type Number interface {
	Integer | Float
}

var (
	// errSegNodeDoNotExist is an error that accurs whenever an invalid segnode id is accessed.
	errIndexOutOfBound = errors.New("index out of bound")
)

type Queue[T Number] struct {
	size int64
	data []T
}

func New[T Number](nums ...T) *Queue[T] {
	return &Queue[T]{
		size: int64(len(nums)),
		data: nums,
	}
}

func (q *Queue[T]) Front() (T, error) {
	if q.size < 1 {
		return 0, errIndexOutOfBound
	}
	return q.data[0], nil
}

func (q *Queue[T]) Push(x T) {
	q.data = append(q.data, x)
	q.size += 1
}

func (q *Queue[T]) PopBack() (T, error) {
	if q.size < 1 {
		return 0, errIndexOutOfBound
	}
	t := make([]T, q.size-1)
	b := q.data[q.size-1]
	copy(t, q.data[:q.size-1])
	q.size -= 1
	q.data = t
	return b, nil
}

func (q *Queue[T]) PopFront() (T, error) {
	if q.size < 1 {
		return 0, errIndexOutOfBound
	}
	t := make([]T, q.size-1)
	f := q.data[0]
	copy(t, q.data[1:])
	q.size -= 1
	q.data = t
	return f, nil
}

func (q Queue[T]) Size() int64 {
	return q.size
}
