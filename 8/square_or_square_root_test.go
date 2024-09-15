package kata8

import "testing"

func TestSquareOrSquareRoot(t *testing.T) {
	testCases := []struct {
		arr []int
		out []int
	}{
		{[]int{4, 3, 9, 7, 2, 1}, []int{2, 9, 3, 49, 4, 1}},
		{[]int{100, 101, 5, 5, 1, 1}, []int{10, 10201, 25, 25, 1, 1}},
		{[]int{1, 2, 3, 4, 5, 6}, []int{1, 4, 9, 2, 25, 36}},
	}

	for _, test := range testCases {
		got := SquareOrSquareRoot(test.arr)
		want := test.out
		for i, n := range got {
			if want[i] != n {
				t.Errorf("Expected %v but got %v", want[i], n)
			}
		}
	}
}
