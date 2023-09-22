package fundamentals

// RecursiveGCD returns GCD using recursive approach.
func RecursiveGCD(a, b int) int {
	if b == 0 {
		return a
	}
	return RecursiveGCD(b, a%b)
}

// IterativeGCD returns GCD using iterative approach.
func IterativeGCD(a, b int) int {
	for b > 0 {
		a %= b
		a, b = b, a
	}
	return a
}

// GCD returns the greated common divisor of a and b
func GCD(a, b int, isIterative bool) int {
	if isIterative {
		return IterativeGCD(a, b)
	}
	return RecursiveGCD(a, b)
}
