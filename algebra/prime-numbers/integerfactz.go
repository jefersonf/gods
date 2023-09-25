package primenumbers

func TrialDivision(n uint64) []uint64 {
	factorization := make([]uint64, 0, 2)
	var d uint64
	for d = 2; d*d <= n; d++ {
		for n%d == 0 {
			factorization = append(factorization, d)
			n /= d
		}
	}

	if n > 1 {
		factorization = append(factorization, n)
	}

	return factorization
}
