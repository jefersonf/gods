package fenwick

import (
	"testing"
)

func TestGet(t *testing.T) {
	n := 100
	a := make([]int64, n)
	for i := range a {
		a[i] = int64(i)
	}
	tree := New(a...)
	for i := range a {
		v := tree.Get(i)
		if a[i] != v {
			t.Errorf("Get(%d) = %v, want %v", i, v, a[i])

		}
	}
}

func TestSet(t *testing.T) {
	n := 100
	a := make([]int64, n)
	for i := range a {
		a[i] = int64(i)
	}
	tree := New(a...)
	for i := range a {
		tree.Set(i, 42)
		v := tree.Get(i)
		if v != 42 {
			t.Errorf("Set(%d, %v) results in Get(%d) = %v, want %v", i, 42, i, v, 42)

		}
	}
}

func TestAdd(t *testing.T) {
	n := 100
	a := make([]int64, n)
	for i := range a {
		a[i] = int64(i)
	}
	tree := New(a...)
	for i := range a {
		tree.Add(i, 1)
		v := tree.Get(i)
		if v != a[i]+1 {
			t.Errorf("Add(%d, %v) must result in Get(%d) = %v, want %v", i, 1, i, v, a[i]+1)
		}
	}
}

func TestAppend(t *testing.T) {
	n := 100
	a := make([]int64, n)
	tree := new(BinaryIndexed[int64])
	if tree.Len() != 0 {
		t.Errorf("Len() = %v, want %v", tree.Len(), 0)
	}
	for i := range a {
		a[i] = int64(i)
		tree.Append(int64(i))
	}
	if tree.Len() != n {
		t.Errorf("Len() = %v, want %v", tree.Len(), n)
	}
	for i := range a {
		var res int64
		for j := 0; j < i; j++ {
			res += a[j]
		}
		if tree.Get(i) != a[i] {
			t.Errorf("Get(%d) = %v, want %v", i, tree.Get(i), a[i])
		}
		if tree.Sum(i) != res {
			t.Errorf("Sum(%d) = %v, want %v", i, tree.Sum(i), res)
		}
	}
}

func TestSum(t *testing.T) {
	n := 100
	a := make([]int64, n)
	for i := range a {
		a[i] = int64(i)
	}
	tree := New(a...)
	for i := range a {
		var res int64
		for j := 0; j < i; j++ {
			res += a[j]
		}
		sum := tree.Sum(i)
		if sum != res {
			t.Errorf("Sum(%d) = %v, want %v", i, sum, res)
		}
	}
}

func TestSumRange(t *testing.T) {
	n := 100
	a := make([]int64, n)
	for i := range a {
		a[i] = int64(i)
	}
	tree := New(a...)
	for i := range a {
		for j := i; j < n; j++ {
			var res int64
			for k := i; k < j; k++ {
				res += a[k]
			}
			if tree.SumRange(i, j) != res {
				t.Errorf("SumRange(%d, %d) = %v, want %v", i, j, tree.SumRange(i, j), res)
			}
		}
	}
}
