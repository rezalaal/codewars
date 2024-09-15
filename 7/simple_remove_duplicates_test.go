package kata7

import "testing"

func TestSimpleRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		arr []int
		out []int
	}{
		{[]int{3, 4, 4, 3, 6, 3}, []int{4, 6, 3}},
		{[]int{1, 2, 1, 2, 1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{1, 1, 4, 5, 1, 2, 1}, []int{4, 5, 2, 1}},
		{[]int{1, 2, 1, 2, 1, 1, 3}, []int{2, 1, 3}},
		{[]int{0, 4, 4, 3, 0, 3}, []int{4, 0, 3}},
	}

	for _, test := range testCases {
		got := SimpleRemoveDuplicates(test.arr)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = SimpleRemoveDuplicates1(test.arr)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
