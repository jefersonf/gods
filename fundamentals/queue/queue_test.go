package queue

import "testing"

func TestFront(t *testing.T) {
	q1 := New(1, 2)
	v, err := q1.Front()
	if err != nil {
		t.Errorf("got %v, want nil", err)
	}
	if v != 1 {
		t.Errorf("got %v, want %v", v, 1)
	}
	a := make([]int, 0)
	q2 := New(a...)
	_, err = q2.Front()
	if err == nil {
		t.Errorf("got nil, want %v", errIndexOutOfBound)
	}
}

func TestPopBack(t *testing.T) {
	q1 := New(1, 2)
	sz := q1.size
	v, err := q1.PopBack()
	if err != nil {
		t.Errorf("got %v, want nil", err)
	}
	if v != 2 {
		t.Errorf("got %v, want %v", v, 2)
	}
	if q1.size != uint64(sz-1) {
		t.Errorf("got %v, want %v", sz, sz-1)
	}
	a := make([]int, 0)
	q2 := New(a...)
	_, err = q2.PopBack()
	if err == nil {
		t.Errorf("got nil, want %v", errIndexOutOfBound)
	}
}

func TestPopFront(t *testing.T) {
	q1 := New(1, 2)
	sz := q1.size
	v, err := q1.PopFront()
	if err != nil {
		t.Errorf("got %v, want nil", err)
	}
	if v != 1 {
		t.Errorf("got %v, want %v", v, 1)
	}
	if q1.size != uint64(sz-1) {
		t.Errorf("got %v, want %v", sz, sz-1)
	}
	a := make([]int, 0)
	q2 := New(a...)
	_, err = q2.PopFront()
	if err == nil {
		t.Errorf("got nil, want %v", errIndexOutOfBound)
	}
}

func TestPush(t *testing.T) {
	a := make([]int, 0)
	q1 := New(a...)
	sz := q1.size
	v := 10
	q1.Push(v)
	front, err := q1.Front()
	if err != nil {
		t.Errorf("got %v, want nil", err)
	}
	if front != v {
		t.Errorf("got %v, want %v", front, v)
	}
	if q1.size != uint64(sz+1) {
		t.Errorf("got %v, want %v", q1.size, sz-1)
	}
}

func TestSize(t *testing.T) {
	q1 := New(1, 2, 3)
	sz := q1.size
	if q1.Size() != sz {
		t.Errorf("got %v, want %v", q1.Size(), sz)
	}
}
