package kata7

import (
	"math"
	"testing"
)

func TestIsNegativeZero(t *testing.T) {
	testCases := []struct{
		n float64
		out bool
	}{
		{5.0, false},
		{4.0, false},
		{3.0, false},
		{2.0, false},
		{1.0, false},
		{-5.0, false},
		{-4.0, false},
		{-2.0, false},
		{-1.0, false},
		{math.Inf(-1), false},
		{math.Inf(1), false},
		{0.0, false},
		{math.NaN(), false},
		{math.MaxFloat64, false},
		{-math.MaxFloat64, false},
		{math.SmallestNonzeroFloat64, false},
		{math.Copysign(0, -1), true},
	}
	for _, test := range testCases {
		got := IsNegativeZero(test.n)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}