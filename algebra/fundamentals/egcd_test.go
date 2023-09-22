package fundamentals

import "testing"

func TestExtendedGCD(t *testing.T) {
	var (
		a         = 55
		b         = 80
		x0        = 3
		y0        = -2
		gcdResult = 5
	)

	var x, y int
	gcd := ExtendedGCD(a, b, &x, &y)
	if gcd != gcdResult {
		t.Errorf("Got GCD=%d x=%d y=%d, want GCD=%d x=%d y=%d", gcd, x, y, gcdResult, x0, y0)
	}
}

func TestIterativeExtendedGCD(t *testing.T) {
	var (
		a         = 55
		b         = 80
		x0        = 3
		y0        = -2
		gcdResult = 5
	)

	var x, y int
	gcd := IterativeExtendedGCD(a, b, &x, &y)
	if gcd != gcdResult {
		t.Errorf("Got GCD=%d x=%d y=%d, want GCD=%d x=%d y=%d", gcd, x, y, gcdResult, x0, y0)
	}
}
