package kata8

import "testing"

func TestFindMultiples(t *testing.T) {
	testCases := []struct {
		integer int
		limit   int
		out     []int
	}{
		{5, 25, []int{5, 10, 15, 20, 25}},
		{1, 2, []int{1, 2}},
	}

	for _, test := range testCases {
		got := FindMultiples(test.integer, test.limit)
		expected := test.out

		for i, n := range got {
			if expected[i] != n {
				t.Errorf("Expected %v but got %v", expected[i], got)
			}
		}
	}
}
