package kata7

import "testing"

func TestSeven(t *testing.T) {
	testCases := []struct {
		n   int64
		out []int
	}{
		{371, []int{35, 1}},
		{1603, []int{7, 2}},
		{477557101, []int{28, 7}},
		{477557102, []int{47, 7}},
	}

	for _, test := range testCases {
		got := Seven(test.n)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
