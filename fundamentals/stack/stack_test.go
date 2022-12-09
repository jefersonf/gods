package stack

import (
	"testing"
)

func TestTop(t *testing.T) {
	s1 := New(1, 2)
	v, err := s1.Top()
	if err != nil {
		t.Error("Non-empty stack, must return nil")
	}
	if v != 2 {
		t.Errorf("Top() = %v, want %v", v, 2)
	}
	a := make([]int, 0)
	s2 := New(a...)
	_, err = s2.Top()
	if err == nil {
		t.Error("Empty stack")
	}
}

func TestPop(t *testing.T) {
	s1 := New(1, 2, 3)
	sz := s1.size
	v1, err := s1.Pop()
	if err != nil {
		t.Error("Non-empty stack, must return nil")
	}
	if v1 != 3 {
		t.Errorf("Pop() = %v, want %v", v1, 3)
	}
	if s1.size != int64(sz-1) {
		t.Errorf("Stack size is still %v, want %v", sz, sz-1)
	}
	s1.Pop()
	v2, _ := s1.Pop()
	if v2 != 1 {
		t.Errorf("Pop() = %v, want %v", v2, 1)
	}
	_, err = s1.Pop()
	if err == nil {
		t.Errorf("Empty stack, must return an error: %v", err)
	}
}

func TestPush(t *testing.T) {
	a := make([]int, 0)
	s1 := New(a...)
	sz := s1.size
	v := 10
	s1.Push(v)
	top, err := s1.Top()
	if err != nil {
		t.Error("Non-empty stack, must result nil")
	}
	if top != v {
		t.Errorf("Get %v, want %v", top, v)
	}
	if s1.size != int64(sz+1) {
		t.Errorf("Stack size is still %v, want %v", s1.size, sz-1)
	}
}

func TestSize(t *testing.T) {
	s1 := New(1, 2, 3)
	sz := s1.size
	if s1.Size() != sz {
		t.Errorf("stack size must be %v", sz)
	}
}
