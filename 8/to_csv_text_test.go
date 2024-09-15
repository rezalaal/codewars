package kata8

import (
	"testing"
)

func TestToCsvText(t *testing.T) {
	testCases := []struct {
		array [][]int
		out   string
	}{
		{
			[][]int{
				{0, 1, 2, 3, 45},
				{10, 11, 12, 13, 14},
				{20, 21, 22, 23, 24},
				{30, 31, 32, 33, 34},
			},
			"0,1,2,3,45\n10,11,12,13,14\n20,21,22,23,24\n30,31,32,33,34",
		},
		{
			[][]int{
				{-25, 21, 2, -33, 48},
				{30, 31, -32, 33, -34},
			},
			"-25,21,2,-33,48\n30,31,-32,33,-34",
		},
		{
			[][]int{
				{5, 55, 5, 5, 55},
				{6, 6, 66, 23, 24},
				{666, 31, -66 - 22, 33, 7},
			},
			"5,55,5,5,55\n6,6,66,23,24\n666,31,-88,33,7",
		},
	}

	for _, test := range testCases {
		got := ToCsvText(test.array)
		expected := test.out

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}

		got = ToCsvText1(test.array)
		expected = test.out

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}
	}
}
