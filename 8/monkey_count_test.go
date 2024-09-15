package kata8

import "testing"

func TestMonkeyCount(t *testing.T) {
	testCases := []struct {
		n   int
		out []int
	}{
		{1, []int{1}},
		{2, []int{1, 2}},
		{3, []int{1, 2, 3}},
		{4, []int{1, 2, 3, 4}},
		{5, []int{1, 2, 3, 4, 5}},
		{6, []int{1, 2, 3, 4, 5, 6}},
		{7, []int{1, 2, 3, 4, 5, 6, 7}},
		{8, []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{9, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{10, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, test := range testCases {
		got := MonkeyCount(test.n)
		expected := test.out

		for i, n := range got {
			if expected[i] != n {
				t.Errorf("Expected %v but got %v", expected[i], got)
			}
		}
	}
}
