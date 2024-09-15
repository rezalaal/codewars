package kata8

import "testing"

func TestCountBy(t *testing.T) {
	testCases := []struct {
		x int
		n int
		out []int
	}{
		{1, 5,   []int{1, 2, 3, 4, 5}},
		{2, 5,   []int{2, 4, 6, 8, 10}},
		{3, 5,   []int{3, 6, 9, 12, 15}},
		{50, 5,  []int{50, 100, 150, 200, 250}},
		{100, 5, []int{100, 200, 300, 400, 500}},
	}

	for _, test :=range testCases {
		got := CountBy(test.x, test.n)
		expected := test.out

		for i, n := range got {
			if expected[i] != n {
				t.Errorf("Expected %v but got %v", expected[i], n)
			}
		}
	}
}