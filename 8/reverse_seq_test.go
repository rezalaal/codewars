package kata8

import "testing"

func TestReverseSeq(t *testing.T) {
	testCases := []struct {
		n   int
		out []int
	}{
		{5, []int{5, 4, 3, 2, 1}},
	}

	for _, test := range testCases {
		got := ReverseSeq(test.n)
		want := test.out

		for i, n := range got {
			if want[i] != n {
				t.Errorf("Expected %v but got %v", want[i], n)
			}
		}
	}
}
