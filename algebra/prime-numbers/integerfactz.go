package primenumbers

func TrialDivision[T Number](n T) []T {
	factorization := make([]T, 0, 2)
	var d T
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
