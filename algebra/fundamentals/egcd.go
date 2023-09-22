package fundamentals

// ExtendedGCD returns the GCD and points to x and y such that ax + by = gcd(a, b).
func ExtendedGCD(a, b int, x, y *int) int {
	if b == 0 {
		*x = 1
		*y = 0
		return a
	}
	var x1, y1 int
	d := ExtendedGCD(b, a%b, &x1, &y1)
	*x = y1
	*y = x1 - y1*(a/b)
	return d
}

// IterativeExtendedGCD returns the GCD and points to x and y such that ax + by = gcd(a, b)
// using iterative approach.
func IterativeExtendedGCD(a, b int, x, y *int) int {
	*x, *y = 1, 0
	var (
		x1 = 0
		y1 = 1
		a1 = a
		b1 = b
	)
	for b1 > 0 {
		q := a1 / b1
		*x, x1 = x1, (*x)-q*x1
		*y, y1 = y1, (*y)-q*y1
		a1, b1 = b1, a1-q*b1
	}
	return a1
}
