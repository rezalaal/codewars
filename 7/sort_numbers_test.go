package kata7

import "testing"

func TestSortNumbers(t *testing.T) {
	testCases := []struct {
		numbers []int
		out     []int
	}{
		{[]int{1, 2, 10, 50, 5}, []int{1, 2, 5, 10, 50}},
		{[]int{}, []int{}},
	}

	for _, test := range testCases {
		got := SortNumbers(test.numbers)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = SortNumbers1(test.numbers)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = SortNumbers2(test.numbers)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
