package kata8

import "testing"

func TestPowersOfTwo(t *testing.T) {
	testCases := []struct {
		n   int
		out []uint64
	}{
		{0, []uint64{1}},
		{1, []uint64{1, 2}},
		{4, []uint64{1, 2, 4, 8, 16}},
	}
	for _, test := range testCases {
		got := PowersOfTwo(test.n)
		want := test.out

		for i, n := range got {
			if want[i] != n {
				t.Errorf("Expected %v but got %v", want[i], n)
			}
		}
	}
}
