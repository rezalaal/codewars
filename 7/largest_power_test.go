package kata7

import "testing"

func TestLargestPower(t *testing.T) {
	testCases := []struct {
		n   uint64
		out int
	}{
		{3, 0},
		{5, 1},
		{7, 1},
		{81, 3},
		{82, 4},
	}

	for _, test := range testCases {
		got := LargestPower(test.n)
		want := test.out

		assertCorrectMessage(t, want, got)
	}
}
