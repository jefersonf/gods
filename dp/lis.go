package dp

import (
	"golang.org/x/exp/constraints"
)

// Number represents a Numeric value.
type Number interface {
	constraints.Integer
}

// LIS returns the longest inscreasing subsequence in the given array of numbers.
func LIS[T Number](a ...T) (uint64, []T) {
	inputSize := len(a)
	aux := make([]uint64, inputSize)
	mapAux := make(map[int]int)
	var k int
	for i := 1; i < inputSize; i++ {
		for j := 0; j < i; j++ {
			if a[i] > a[j] && aux[j]+1 > aux[i] {
				aux[i] = aux[j] + 1
				mapAux[i] = j
				k = i
			}
		}
	}

	n := uint64(1)
	for i := range aux {
		if aux[i]+1 > n {
			n = aux[i] + 1
		}
	}
	sample := make([]T, n)
	var ok bool
	nc := n - 1
	sample[nc] = a[k]
	for {
		k, ok = mapAux[k]
		if !ok {
			break
		}
		nc -= 1
		sample[nc] = a[k]
	}
	return n, sample
}
