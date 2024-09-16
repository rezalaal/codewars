package kata7

import "testing"

func TestArrowArea(t *testing.T) {
	testCases := []struct {
		a   int
		b   int
		out float64
	}{
		{4, 2, 2.0},
		{7, 6, 10.5},
		{25, 25, 156.25},
	}

	for _, test := range testCases {
		got := ArrowArea(test.a, test.b)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
