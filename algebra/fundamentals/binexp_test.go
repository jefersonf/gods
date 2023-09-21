package fundamentals

import "testing"

func TestRecursivePow(t *testing.T) {

	testcases := []struct {
		base uint64
		exp  uint64
		pow  uint64
	}{
		{
			base: 1,
			exp:  0,
			pow:  1,
		},
		{
			base: 1,
			exp:  1,
			pow:  1,
		},
		{
			base: 1,
			exp:  3,
			pow:  1,
		},
		{
			base: 2,
			exp:  5,
			pow:  32,
		},
		{
			base: 2,
			exp:  8,
			pow:  256,
		},
		{
			base: 3,
			exp:  4,
			pow:  81,
		},
		{
			base: 3,
			exp:  7,
			pow:  2187,
		},
	}

	for _, tc := range testcases {
		pow := RecursivePow(tc.base, tc.exp)
		if pow != tc.pow {
			t.Errorf("Pow(%d, %d), Got %d, want %d", tc.base, tc.exp, pow, tc.pow)
		}
	}
}

func TestIterativePow(t *testing.T) {

	testcases := []struct {
		base uint64
		exp  uint64
		pow  uint64
	}{
		{
			base: 1,
			exp:  0,
			pow:  1,
		},
		{
			base: 1,
			exp:  1,
			pow:  1,
		},
		{
			base: 1,
			exp:  3,
			pow:  1,
		},
		{
			base: 2,
			exp:  5,
			pow:  32,
		},
		{
			base: 2,
			exp:  8,
			pow:  256,
		},
		{
			base: 3,
			exp:  4,
			pow:  81,
		},
		{
			base: 3,
			exp:  7,
			pow:  2187,
		},
	}

	for _, tc := range testcases {
		pow := IterativePow(tc.base, tc.exp)
		if pow != tc.pow {
			t.Errorf("Pow(%d, %d), Got %d, want %d", tc.base, tc.exp, pow, tc.pow)
		}
	}
}
