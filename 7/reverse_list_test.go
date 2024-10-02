package kata7

import "testing"


func TestReverseList(t *testing.T) {
	testCases := []struct{
		lst []int
		out []int
	}{
		{[]int{1,2,3,4}, []int{4,3,2,1}},
	}

	for _, test := range testCases {
		got := ReverseList(test.lst)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}