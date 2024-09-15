package kata6

import "testing"

func TestParse(t *testing.T) {
	testCases := []struct {
		data string
		out  []int
	}{
		{"ooo", []int{0, 0, 0}},
		{"ioioio", []int{1, 2, 3}},
		{"idoiido", []int{0, 1}},
		{"isoisoiso", []int{1, 4, 25}},
		{"codewars", []int{0}},
	}

	for _, test := range testCases {
		got := Parse(test.data)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
