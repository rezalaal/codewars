package kata6

import (
	"testing"
)

func TestArrayDiff(t *testing.T) {
	testCases := []struct {
		a []int
		b []int
		r []int
	}{
		{[]int{1, 2}, []int{1}, []int{2}},
		{[]int{1, 2, 2}, []int{1}, []int{2, 2}},
		{[]int{1, 2, 2}, []int{2}, []int{1}},
		{[]int{1, 2, 2}, []int{}, []int{1, 2, 2}},
		{[]int{}, []int{1, 2}, []int{}},
		{[]int{1, 2, 3}, []int{1, 2}, []int{3}},
	}

	for _, test := range testCases {
		got := ArrayDiff(test.a, test.b)
		want := test.r
		for i, n := range got {
			if want[i] != n {
				t.Errorf("expected %v but got %v", want[i], n)
			}
		}

		got = ArrayDiff1(test.a, test.b)
		want = test.r
		for i, n := range got {
			if want[i] != n {
				t.Errorf("expected %v but got %v", want[i], n)
			}
		}
	}
}
