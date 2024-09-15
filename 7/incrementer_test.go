package kata7

import "testing"

func TestIncrementer(t *testing.T) {
	testCases := []struct {
		n   []int
		out []int
	}{
		{[]int{}, []int{}},
		{[]int{1, 2, 3}, []int{2, 4, 6}},
		{[]int{4, 6, 7, 1, 3}, []int{5, 8, 0, 5, 8}},
		{[]int{3, 6, 9, 8, 9}, []int{4, 8, 2, 2, 4}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 9, 9, 9, 8}, []int{2, 4, 6, 8, 0, 2, 4, 6, 8, 9, 0, 1, 2, 2}},
	}

	for _, test := range testCases {
		got := Incrementer(test.n)
		want := test.out
		assertCorrectMessage(t, want, got)		
	}
}
