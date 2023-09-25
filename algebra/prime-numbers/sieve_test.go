package primenumbers

import (
	"testing"
)

func TestSieve(t *testing.T) {
	primeNumbersUnder100 := []int{
		2, 3, 5, 7, 11, 13, 17, 19, 23,
		29, 31, 37, 41, 43, 47, 53, 59,
		61, 67, 71, 73, 79, 83, 89, 97,
	}

	primes := make(map[int]struct{})
	for _, p := range primeNumbersUnder100 {
		primes[p] = struct{}{}
	}

	sieve := NewSieve(100)
	for i, isPrime := range sieve {
		_, tc := primes[i]
		if tc != isPrime {
			t.Errorf("Got %t, want %t", isPrime, tc)
		}
	}
}
