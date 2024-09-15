package kata7

import "testing"

func TestTwoOldestAges(t *testing.T) {
	testCases := []struct{
		ages []int
		out [2]int
	}{
		{[]int{6,5,83,5,3,18},[2]int{18,83}},
		{[]int{1,5,87,45,8,8},[2]int{45,87}	},
	}

	for _, test := range testCases {
		got := TwoOldestAges(test.ages)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = TwoOldestAges1(test.ages)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = TwoOldestAges2(test.ages)
		want = test.out
		assertCorrectMessage(t, want, got)

	}
}