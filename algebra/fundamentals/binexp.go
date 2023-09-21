package fundamentals

import "golang.org/x/exp/constraints"

// Number represents a Numeric value.
type Number interface {
	constraints.Integer
}

func RecursivePow[T Number](a, b T) T {
	if b == 0 {
		return 1
	}
	p := RecursivePow(a, b/2)
	if b%2 != 0 {
		return p * p * a
	}
	return p * p
}

func RecursivePowMod[T Number](a, b, m T) T {
	if b == 0 {
		return 1
	}
	p := RecursivePow(a, b/2) % m
	if b%2 != 0 {
		return p * p * a % m
	}
	return p * p % m
}

func IterativePow[T Number](a, b T) T {
	var ans T = 1
	for b > 0 {
		if b%2 != 0 {
			ans = ans * a
		}
		a *= a
		b >>= 1
	}
	return ans
}

func IterativePowMod[T Number](a, b, m T) T {
	var ans T = 1
	for b > 0 {
		if b%2 != 0 {
			ans = ans * a % m
		}
		a = a * a % m
		b >>= 1
	}
	return ans % m
}

func Pow[T Number](a, b T, isIterative bool) T {
	if isIterative {
		return IterativePow(a, b)
	}
	return RecursivePow(a, b)
}
