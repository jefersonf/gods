package primenumbers

type sieve []bool

func NewSieve(n int) sieve {
	isPrime := make([]bool, n+1)
	for i := range isPrime {
		if i >= 2 {
			isPrime[i] = true
		}
	}
	for i := 2; i <= n; i++ {
		if isPrime[i] && i*i <= n {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	return isPrime
}
