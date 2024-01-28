package dp

import (
	"testing"
)

func TestLIS(t *testing.T) {

	testcases := []struct {
		input []int
		want  int
	}{
		{
			input: []int{3, 10, 2, 1, 20},
			want:  3,
		},
		{
			input: []int{3, 2},
			want:  1,
		},
		{
			input: []int{50, 3, 10, 7, 40, 80},
			want:  4,
		},
		{
			input: []int{1},
			want:  1,
		},
		{
			input: []int{1, 2, 3, 4, 5},
			want:  5,
		},
		{
			input: []int{50, 3, 10, 10, 7, 40, 80, 80},
			want:  4,
		},
	}

	for _, tc := range testcases {
		n, sample := LIS(tc.input...)
		if n != uint64(tc.want) {
			t.Fatalf("Output array length mismath: got %v, want %v", n, tc.want)
		}
		for i := 1; i < len(sample); i++ {
			if sample[i] < sample[i-1] {
				t.Fatalf("Output array does not match contraint: got %v", sample)
			}
		}

	}
}
