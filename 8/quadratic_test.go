package kata8

import "testing"

func TestQuadratic(t *testing.T) {
	testCases := []struct {
		x1  int
		x2  int
		out [3]int
	}{
		{0, 1, [3]int{1, -1, 0}},
		{4, 9, [3]int{1,-13, 36}},
		{2, 6, [3]int{1,-8, 12}},
		{-5, -4, [3]int{1, 9, 20}},
	}
	for _, test := range testCases {
		got := Quadratic(test.x1, test.x2)
		want := test.out
		for i, n := range got {
			if want[i] != n {
				t.Errorf("Expected %v but got %v", want[i], n)
			}
		}
	}
}
