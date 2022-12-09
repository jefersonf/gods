package stack

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

type Stack[T Number] struct {
	size int64
	data []T
}

var (
	// errIndexOutOfBound is an error that accurs whenever an invalid invalid index is accessed.
	errIndexOutOfBound = errors.New("index out of bound")
)

func New[T Number](nums ...T) *Stack[T] {
	return &Stack[T]{
		size: int64(len(nums)),
		data: nums,
	}
}

func (s *Stack[T]) Top() (T, error) {
	if s.size < 1 {
		return 0, errIndexOutOfBound
	}
	top := s.data[s.size-1]
	return top, nil
}

func (s *Stack[T]) Pop() (T, error) {
	if s.size < 1 {
		return 0, errIndexOutOfBound
	}
	top := s.data[s.size-1]
	s.size -= 1
	return top, nil
}

func (s *Stack[T]) Push(x T) {
	s.data = append(s.data, x)
	s.size += 1
}

func (s Stack[T]) Size() int64 {
	return s.size
}
