package kata8

import "testing"

func TestMultipleOfIndex(t *testing.T) {
	testCases := []struct{
		ints []int
		out []int
	}{
		{[]int{22, -6, 32, 82, 9, 25},                                                         []int{-6, 32, 25}},
		{[]int{68, -1, 1, -7, 10, 10},                                                         []int{-1, 10}},
		{[]int{11, -11},                                                                       []int{-11}},
		{[]int{-56, -85, 72, -26, -14, 76, -27, 72, 35, -21, -67, 87, 0, 21, 59, 27, -92, 68}, []int{-85, 72, 0, 68}},
		{[]int{28, 38, -44, -99, -13, -54, 77, -51},                                           []int{38, -44, -99}},
		{[]int{-1, -49, -1, 67, 8, -60, 39, 35},                                               []int{-49, 8, -60, 35}},
	}

	for _, test := range testCases {
		got := MultipleOfIndex(test.ints)
		want := test.out
		for i, n :=range got {
			if want[i] != n {
				t.Errorf("Expected %v but got %v", want[i], n)
			}
		}
	}
}