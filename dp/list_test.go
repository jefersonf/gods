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
	}

	for _, tc := range testcases {
		n, _ := LIS(tc.input...)
		if n != uint64(tc.want) {
			t.Fatalf("Output array length mismath: got %v, want %v", n, tc.want)
		}
	}
}
