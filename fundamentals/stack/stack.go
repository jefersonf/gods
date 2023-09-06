package stack

import (
	"errors"
)

type Stack[T any] struct {
	size uint64
	data []T
}

var (
	// errIndexOutOfBound is an error that accurs whenever an invalid invalid index is accessed.
	errIndexOutOfBound = errors.New("index out of bound")
)

// New returns a new stack
func New[T any](nums ...T) *Stack[T] {
	return &Stack[T]{
		size: uint64(len(nums)),
		data: nums,
	}
}

// Top returns the top element of the stack but keeps it on stack
func (s *Stack[T]) Top() (v T, err error) {
	if s.size < 1 {
		err = errIndexOutOfBound
		return v, err
	}
	top := s.data[s.size-1]
	return top, err
}

// Pop returns the top element of the stack and also removes it from the stack
func (s *Stack[T]) Pop() (v T, err error) {
	if s.size < 1 {
		err = errIndexOutOfBound
		return v, err
	}
	top := s.data[s.size-1]
	s.size -= 1
	return top, err
}

// Push insert a new element on top of the stack
func (s *Stack[T]) Push(x T) {
	s.data = append(s.data, x)
	s.size += 1
}

// Size returns the number of elements in the stack
func (s Stack[T]) Size() uint64 {
	return s.size
}
