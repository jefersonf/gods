package stack

import (
	"testing"
)

func TestTop(t *testing.T) {
	s1 := New(1, 2)
	v, err := s1.Top()
	if err != nil {
		t.Errorf("got %v, want nil", err)
	}
	if v != 2 {
		t.Errorf("got %v, want %v", v, 2)
	}
	a := make([]int, 0)
	s2 := New(a...)
	_, err = s2.Top()
	if err == nil {
		t.Errorf("got nil, want %v", errIndexOutOfBound)
	}
}

func TestPop(t *testing.T) {
	s1 := New(1, 2, 3)
	sz := s1.size
	v1, err := s1.Pop()
	if err != nil {
		t.Errorf("got %v, want nil", err)
	}
	if v1 != 3 {
		t.Errorf("got %v, want %v", v1, 3)
	}
	if s1.size != int64(sz-1) {
		t.Errorf("got %v, want %v", sz, sz-1)
	}
	s1.Pop()
	v2, _ := s1.Pop()
	if v2 != 1 {
		t.Errorf("got %v, want %v", v2, 1)
	}
	_, err = s1.Pop()
	if err == nil {
		t.Errorf("got nil, want %v", errIndexOutOfBound)
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
		t.Errorf("got %v, want nil", err)
	}
	if top != v {
		t.Errorf("got %v, want %v", top, v)
	}
	if s1.size != int64(sz+1) {
		t.Errorf("got %v, want %v", s1.size, sz-1)
	}
}

func TestSize(t *testing.T) {
	s1 := New(1, 2, 3)
	sz := s1.size
	if s1.Size() != sz {
		t.Errorf("got %v, want %v", s1.Size(), sz)
	}
}
