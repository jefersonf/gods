package dsu

import (
	"fmt"
	"testing"

	"golang.org/x/exp/slices"
)

func TestMakeSet(t *testing.T) {

	testcases := []struct {
		n      uint64
		parent []uint64
		size   []uint64
	}{
		{
			n:      0,
			parent: make([]uint64, 1),
			size:   make([]uint64, 1),
		},
		{
			n:      1,
			parent: []uint64{0, 1},
			size:   []uint64{1, 1},
		},
		{
			n:      5,
			parent: []uint64{0, 1, 2, 3, 4, 5},
			size:   []uint64{1, 1, 1, 1, 1, 1},
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("ArrayOfLength=%d", tc.n), func(t *testing.T) {
			dsu := New(tc.n)
			if slices.Compare(dsu.parent, tc.parent) != 0 {
				t.Errorf("Got %v, want %v", dsu.parent, tc.parent)
			}
		})
	}

}
