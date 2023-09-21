package fundamentals

func RecursivePow(a, b uint64) uint64 {
	if b == 0 {
		return 1
	}
	p := RecursivePow(a, b/2)
	if b%2 != 0 {
		return p * p * a
	}
	return p * p
}

func IterativePow(a, b uint64) uint64 {
	var ans uint64 = 1
	for b > 0 {
		if b%2 != 0 {
			ans = ans * a
		}
		a *= a
		b >>= 1
	}
	return ans
}

func Pow(a, b uint64, isIterative bool) uint64 {
	if isIterative {
		return IterativePow(a, b)
	}
	return RecursivePow(a, b)
}
