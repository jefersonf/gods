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

func TestFindSet(t *testing.T) {

	testcases := []struct {
		n        uint64
		id       uint64
		parentId uint64
		unions   [][2]uint64
	}{
		{
			n:        1,
			id:       1,
			parentId: 1,
			unions:   make([][2]uint64, 0),
		},
		{
			n:        3,
			id:       2,
			parentId: 1,
			unions:   [][2]uint64{{1, 2}},
		},
		{
			n:        3,
			id:       3,
			parentId: 3,
			unions:   [][2]uint64{{1, 2}},
		},
		{
			n:        3,
			id:       3,
			parentId: 1,
			unions:   [][2]uint64{{1, 2}, {2, 3}},
		},
		{
			n:        4,
			id:       4,
			parentId: 3,
			unions:   [][2]uint64{{1, 2}, {3, 4}},
		},
		{
			n:        4,
			id:       4,
			parentId: 3,
			unions:   [][2]uint64{{2, 1}, {4, 3}},
		},
		{
			n:        5,
			id:       3,
			parentId: 1,
			unions:   [][2]uint64{{5, 1}, {4, 3}, {2, 1}, {4, 5}},
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("ArrayOfLength=%d-Unions=%d", tc.n, len(tc.unions)), func(t *testing.T) {
			dsu := New(tc.n)
			for _, p := range tc.unions {
				dsu.UnionSets(p[0], p[1])
			}
			parentId := dsu.FindSet(tc.id)
			if parentId != tc.parentId {
				t.Errorf("Got %v, want %v", parentId, tc.parentId)
			}
		})
	}
}

func TestUnionSets(t *testing.T) {

	testcases := []struct {
		n       uint64
		id1     uint64
		id2     uint64
		unions  [][2]uint64
		sameSet bool
	}{
		{
			n:       1,
			id1:     1,
			id2:     1,
			unions:  [][2]uint64{},
			sameSet: true,
		},
		{
			n:       2,
			id1:     1,
			id2:     2,
			unions:  [][2]uint64{},
			sameSet: false,
		},
		{
			n:       2,
			id1:     1,
			id2:     2,
			unions:  [][2]uint64{{2, 1}},
			sameSet: true,
		},
		{
			n:       2,
			id1:     1,
			id2:     2,
			unions:  [][2]uint64{{1, 2}},
			sameSet: true,
		},
		{
			n:       3,
			id1:     1,
			id2:     3,
			unions:  [][2]uint64{{1, 2}},
			sameSet: false,
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("SameSet=%t", tc.sameSet), func(t *testing.T) {
			dsu := New(tc.n)
			for _, p := range tc.unions {
				dsu.UnionSets(p[0], p[1])
			}
			a := dsu.FindSet(tc.id1)
			b := dsu.FindSet(tc.id2)
			sameSet := a == b
			if sameSet != tc.sameSet {
				t.Errorf("Got %t, want %t", sameSet, tc.sameSet)
			}
		})
	}
}

func TestSize(t *testing.T) {

	testcases := []struct {
		n    uint64
		size uint64
	}{
		{
			n:    0,
			size: 0,
		},
		{
			n:    1,
			size: 1,
		},
		{
			n:    5,
			size: 5,
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("ArrayOfLength=%d", tc.n), func(t *testing.T) {
			dsu := New(tc.n)
			dsuLen := dsu.Size()
			if dsuLen != tc.size {
				t.Errorf("Got %v, want %v", dsuLen, tc.size)
			}
		})
	}

}
