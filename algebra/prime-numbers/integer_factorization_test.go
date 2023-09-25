package primenumbers

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestTrialDivision(t *testing.T) {

	var (
		n     = uint64(98)
		factz = []uint64{2, 7, 7}
	)

	factorization := TrialDivision(n)
	if slices.Compare(factz, factorization) != 0 {
		t.Errorf("Got %v, want %v", factorization, factz)
	}
}
