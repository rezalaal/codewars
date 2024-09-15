package kata7

import "testing"

func TestCountRedBeads(t *testing.T) {
	testCases := []struct {
		n   int
		out int
	}{
		{0, 0},
		{1, 0},
		{3, 4},
		{5, 8},
	}

	for _, test := range testCases {
		got := CountRedBeads(test.n)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
