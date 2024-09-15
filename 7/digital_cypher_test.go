package kata7

import "testing"

func TestEncode(t *testing.T) {
	testCases := []struct {
		str string
		key int
		out []int
	}{
		{"scout", 1939, []int{20, 12, 18, 30, 21}},
		{"masterpiece", 1939, []int{14, 10, 22, 29, 6, 27, 19, 18, 6, 12, 8}},
	}

	for _, test := range testCases {
		got := Encode(test.str, test.key)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
