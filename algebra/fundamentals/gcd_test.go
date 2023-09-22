package fundamentals

import "testing"

type GCDTestCase struct {
	a   int
	b   int
	gcd int
	lcm int
}

func loadGCDTestCases() []GCDTestCase {

	testcases := []GCDTestCase{
		{
			a:   1,
			b:   1,
			gcd: 1,
			lcm: 1,
		},
		{
			a:   2,
			b:   4,
			gcd: 2,
			lcm: 4,
		},
		{
			a:   4,
			b:   16,
			gcd: 4,
			lcm: 16,
		},
		{
			a:   336,
			b:   3130,
			gcd: 2,
			lcm: 525840,
		},
	}

	return testcases
}

func TestRecursiveGCD(t *testing.T) {
	for _, tc := range loadGCDTestCases() {
		gcd := RecursiveGCD(tc.a, tc.b)
		if tc.gcd != gcd {
			t.Errorf("RecursiveGCD(%d, %d) got %d, want %d", tc.a, tc.b, tc.gcd, gcd)
		}
	}
}

func TestIterativeGCD(t *testing.T) {
	for _, tc := range loadGCDTestCases() {
		gcd := IterativeGCD(tc.a, tc.b)
		if tc.gcd != gcd {
			t.Errorf("IterativeGCD(%d, %d) got %d, want %d", tc.a, tc.b, tc.gcd, gcd)
		}
	}
}

func TestLCM(t *testing.T) {
	for _, tc := range loadGCDTestCases() {
		lcm := LCM(tc.a, tc.b)
		if tc.lcm != lcm {
			t.Errorf("LCM(%d, %d) got %d, want %d", tc.a, tc.b, tc.lcm, lcm)
		}
	}
}
