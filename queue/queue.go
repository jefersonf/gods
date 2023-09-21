package queue

import (
	"errors"
)

var (
	// errIndexOutOfBound is an error that accurs whenever an invalid invalid index is accessed.
	errIndexOutOfBound = errors.New("index out of bound")
)

type Queue[T any] struct {
	size uint64
	data []T
}

// Newn returns a new queue
func New[T any](nums ...T) *Queue[T] {
	return &Queue[T]{
		size: uint64(len(nums)),
		data: nums,
	}
}

// Front returns the first element inserted into the queue
func (q *Queue[T]) Front() (v T, err error) {
	if q.size < 1 {
		err = errIndexOutOfBound
		return v, err
	}
	v = q.data[0]
	return v, nil
}

// Push inserts a new element into queue
func (q *Queue[T]) Push(x T) {
	q.data = append(q.data, x)
	q.size += 1
}

// PopBack returns the last element inserted by removing it from queue
func (q *Queue[T]) PopBack() (v T, err error) {
	if q.size < 1 {
		err = errIndexOutOfBound
		return v, err
	}
	t := make([]T, q.size-1)
	b := q.data[q.size-1]
	copy(t, q.data[:q.size-1])
	q.size -= 1
	q.data = t
	return b, err
}

// PopFront returns the first element inserted by removing it from queue
func (q *Queue[T]) PopFront() (v T, err error) {
	if q.size < 1 {
		err = errIndexOutOfBound
		return v, err
	}
	t := make([]T, q.size-1)
	f := q.data[0]
	copy(t, q.data[1:])
	q.size -= 1
	q.data = t
	return f, err
}

// Size returns the number of elements in the queue
func (q Queue[T]) Size() uint64 {
	return q.size
}

// IsEmpty checks whether the queue is empty or not
func (q Queue[T]) IsEmpty() bool {
	return q.size == 0
}
