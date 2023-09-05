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

func New[T any](nums ...T) *Stack[T] {
	return &Stack[T]{
		size: uint64(len(nums)),
		data: nums,
	}
}

func (s *Stack[T]) Top() (v T, err error) {
	if s.size < 1 {
		err = errIndexOutOfBound
		return v, err
	}
	top := s.data[s.size-1]
	return top, err
}

func (s *Stack[T]) Pop() (v T, err error) {
	if s.size < 1 {
		err = errIndexOutOfBound
		return v, err
	}
	top := s.data[s.size-1]
	s.size -= 1
	return top, err
}

func (s *Stack[T]) Push(x T) {
	s.data = append(s.data, x)
	s.size += 1
}

func (s Stack[T]) Size() uint64 {
	return s.size
}
