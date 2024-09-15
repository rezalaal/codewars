package kata7

import "testing"

func TestMinMax(t *testing.T) {
	testCases := []struct{
		arr []int
		out [2]int
	}{
		{[]int{1,2,3,4,5},[2]int{1,5}},
		{[]int{2334454,5},[2]int{5,2334454}},
	}

	for _, test := range testCases {
		got := MinMax(test.arr)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = MinMax1(test.arr)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = MinMax2(test.arr)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}