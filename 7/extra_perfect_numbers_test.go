package kata7

import "testing"

func TestExtraPerfect(t *testing.T) {
	testCases := []struct {
		n   int
		out []int
	}{
		{3, []int{1, 3}},
		{5, []int{1, 3, 5}},
		{7, []int{1, 3, 5, 7}},
		{28, []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27}},
		{39, []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 33, 35, 37, 39}},
	}

	for _, test := range testCases {
		got := ExtraPerfect(test.n)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
