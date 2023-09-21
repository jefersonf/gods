package fundamentals

import "testing"

func TestRecursiveGCD(t *testing.T) {
	var (
		a         = 30
		b         = 16
		gcdResult = 2
	)
	gcd := RecursiveGCD(a, b)
	if gcdResult != gcd {
		t.Errorf("RecursiveGCD(%d, %d) got %d, want %d", a, b, gcd, gcdResult)
	}
}
