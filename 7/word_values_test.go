package kata7

import "testing"

func TestNameValue(t *testing.T) {
	testCases := []struct {
		my_list []string
		out     []int
	}{
		{[]string{"abc", "abc", "abc", "abc"}, []int{6, 12, 18, 24}},
		{[]string{"codewars", "abc", "xyz"}, []int{88, 12, 225}},
	}

	for _, test := range testCases {
		got := NameValue(test.my_list)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
