package primenumbers

import "golang.org/x/exp/constraints"

// Number represents a Numeric value.
type Number interface {
	constraints.Integer
}

type sieve []bool

func NewSieve[T Number](n T) sieve {
	isPrime := make([]bool, n+1)
	for i := range isPrime {
		if i >= 2 {
			isPrime[i] = true
		}
	}
	var i T
	for i = 2; i <= n; i++ {
		if isPrime[i] && i*i <= n {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	return isPrime
}
