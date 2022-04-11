package segment

import (
	"math"
	"reflect"
	"testing"
)

func TestBuild(t *testing.T) {
	n := 5
	a := make([]int32, n)
	for i := range a {
		a[i] = int32(i + 1)
	}

	testcases := []struct {
		in       []int32
		joinType string
		joinFunc func(n1, n2 *SegNode[int32]) SegNode[int32]
		rootNode SegNode[int32]
	}{
		{
			in:       a,
			joinType: "sum",
			joinFunc: func(n1, n2 *SegNode[int32]) SegNode[int32] {
				return SegNode[int32]{
					value: n1.value + n2.value,
				}
			},
			rootNode: SegNode[int32]{
				value: 15,
			},
		},
		{
			in:       a,
			joinType: "minimun",
			joinFunc: func(n1, n2 *SegNode[int32]) SegNode[int32] {
				return min(n1, n2)
			},
			rootNode: SegNode[int32]{
				value: 1,
			},
		},
		{
			in:       a,
			joinType: "maximum",
			joinFunc: func(n1, n2 *SegNode[int32]) SegNode[int32] {
				return max(n1, n2)
			},
			rootNode: SegNode[int32]{
				value: 5,
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.joinType, func(t *testing.T) {
			st := New(tc.in, tc.joinFunc)
			sn, err := st.GetNodeByID(1)
			if err != nil {
				t.Errorf("Invalid id=%d; got %v", 1, err)
			}
			if !reflect.DeepEqual(sn, tc.rootNode) {
				t.Errorf("Get(%d) on %s segtree got %v; want %v", 1, tc.joinType, sn, tc.rootNode)
			}
		})
	}
}

func TestGetNodeByID(t *testing.T) {
	n := 5
	a := make([]int32, n)
	for i := range a {
		a[i] = int32(i + 1)
	}

	testcases := []struct {
		in       []int32
		joinType string
		joinFunc func(n1, n2 *SegNode[int32]) SegNode[int32]
		rootNode SegNode[int32]
	}{
		{
			in:       a,
			joinType: "product",
			joinFunc: func(n1, n2 *SegNode[int32]) SegNode[int32] {
				sn := SegNode[int32]{value: 1}
				if n1.value != 0 {
					sn.value *= n1.value
				}
				if n2.value != 0 {
					sn.value *= n2.value
				}
				return sn
			},
			rootNode: SegNode[int32]{
				value: 120,
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.joinType, func(t *testing.T) {
			st := New(tc.in, tc.joinFunc)
			sn, err := st.GetNodeByID(1)
			if err != nil {
				t.Errorf("Invalid id=%d; got %q", 1, err)
			}
			if sn.value != tc.rootNode.value {
				t.Errorf("Get(%d) got %v; want %v", 1, sn.value, tc.rootNode.value)
			}
		})
	}
}

func TestSize(t *testing.T) {
	n := 100
	a := make([]int32, n)
	st := New(a, nil)
	if st.Size() != 100 {
		t.Errorf("Size() got %v; want %v", st.Size(), 100)
	}
}

func TestSet(t *testing.T) {
	n := 5
	a := make([]int32, n)
	for i := range a {
		a[i] = int32(i + 1)
	}
	testcases := []struct {
		in       []int32
		joinType string
		joinFunc func(n1, n2 *SegNode[int32]) SegNode[int32]
		rootNode SegNode[int32]
	}{
		{
			in:       a,
			joinType: "sum",
			joinFunc: func(n1, n2 *SegNode[int32]) SegNode[int32] {
				return SegNode[int32]{
					value: n1.value + n2.value,
				}
			},
			rootNode: SegNode[int32]{
				value: 114,
			},
		},
		{
			in:       a,
			joinType: "min",
			joinFunc: func(n1, n2 *SegNode[int32]) SegNode[int32] {
				return min(n1, n2)
			},
			rootNode: SegNode[int32]{
				value: 2,
			},
		},
		{
			in:       a,
			joinType: "max",
			joinFunc: func(n1, n2 *SegNode[int32]) SegNode[int32] {
				return max(n1, n2)
			},
			rootNode: SegNode[int32]{
				value: 100,
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.joinType, func(t *testing.T) {
			st := New(tc.in, tc.joinFunc)
			st.Set(1, 100)
			sn, err := st.GetNodeByID(1)
			if err != nil {
				t.Errorf("Invalid id=%d; got %q", 1, err)
			}
			if !reflect.DeepEqual(sn, tc.rootNode) {
				t.Errorf("Set(%d) on %s segtree got %v; want %v", 1, tc.joinType, sn, tc.rootNode)
			}
		})
	}
}

func TestQueryRange(t *testing.T) {
	n := 5
	a := make([]int32, n)
	for i := range a {
		a[i] = int32(i + 1)
	}

	testcases := []struct {
		in           []int32
		joinType     string
		joinFunc     func(n1, n2 *SegNode[int32]) SegNode[int32]
		rangeResults [5][5]int32
	}{
		{
			in:       a,
			joinType: "product",
			joinFunc: func(n1, n2 *SegNode[int32]) SegNode[int32] {
				sn := SegNode[int32]{value: 1}
				if n1.value != 0 {
					sn.value *= n1.value
				}
				if n2.value != 0 {
					sn.value *= n2.value
				}
				return sn
			},
			rangeResults: [5][5]int32{
				{1, 2, 6, 24, 120},
				{0, 2, 6, 24, 120},
				{0, 0, 3, 12, 60},
				{0, 0, 0, 4, 20},
				{0, 0, 0, 0, 5},
			},
		},
		{
			in:       a,
			joinType: "sum",
			joinFunc: func(n1, n2 *SegNode[int32]) SegNode[int32] {
				return SegNode[int32]{
					value: n1.value + n2.value,
				}
			},
			rangeResults: [5][5]int32{
				{1, 3, 6, 10, 15},
				{0, 2, 5, 9, 14},
				{0, 0, 3, 7, 12},
				{0, 0, 0, 4, 9},
				{0, 0, 0, 0, 5},
			},
		},
		{
			in:       a,
			joinType: "min",
			joinFunc: func(n1, n2 *SegNode[int32]) SegNode[int32] {
				return min(n1, n2)
			},
			rangeResults: [5][5]int32{
				{1, 1, 1, 1, 1},
				{0, 2, 2, 2, 2},
				{0, 0, 3, 3, 3},
				{0, 0, 0, 4, 4},
				{0, 0, 0, 0, 5},
			},
		},
		{
			in:       a,
			joinType: "max",
			joinFunc: func(n1, n2 *SegNode[int32]) SegNode[int32] {
				return max(n1, n2)
			},
			rangeResults: [5][5]int32{
				{1, 2, 3, 4, 5},
				{0, 2, 3, 4, 5},
				{0, 0, 3, 4, 5},
				{0, 0, 0, 4, 5},
				{0, 0, 0, 0, 5},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.joinType, func(t *testing.T) {
			st := New(tc.in, tc.joinFunc)
			for i := 1; i <= n; i++ {
				for j := i; j <= n; j++ {
					res, _ := st.QueryRange(i, j)
					want := tc.rangeResults[i-1][j-1]
					if want != res.value {
						t.Errorf("QueryRange(%d, %d) on %q got %v; want %v", i, j, tc.joinType, res.value, want)
					}
				}
			}
		})
	}
}

func min[T int32](n1, n2 *SegNode[T]) SegNode[T] {
	sn := SegNode[T]{
		value: math.MaxInt32,
	}
	if n1.value != 0 && n1.value < sn.value {
		sn.value = n1.value
	}
	if n2.value != 0 && n2.value < sn.value {
		sn.value = n2.value
	}
	return sn
}

func max[T int32](n1, n2 *SegNode[T]) SegNode[T] {
	sn := SegNode[T]{
		value: math.MinInt32,
	}
	if n1.value != 0 && n1.value > sn.value {
		sn.value = n1.value
	}
	if n2.value != 0 && n2.value > sn.value {
		sn.value = n2.value
	}
	return sn
}
