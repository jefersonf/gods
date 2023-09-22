package fundamentals

// RecursiveGCD returns GCD using recursive approach.
func RecursiveGCD[T Number](a, b T) T {
	if b == 0 {
		return a
	}
	return RecursiveGCD(b, a%b)
}

// IterativeGCD returns GCD using iterative approach.
func IterativeGCD[T Number](a, b T) T {
	for b > 0 {
		a %= b
		a, b = b, a
	}
	return a
}

// GCD returns the greated common divisor of a and b
func GCD[T Number](a, b T, isIterative bool) T {
	if isIterative {
		return IterativeGCD(a, b)
	}
	return RecursiveGCD(a, b)
}
