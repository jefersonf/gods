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
