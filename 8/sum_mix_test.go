package kata8

import "testing"

func TestSumMix(t *testing.T) {
	testCases := []struct {
		arr []any
		out int
	}{
		{[]any{9, 3, "7", "3"}, 22},
		{[]any{"5", "0", 9, 3, 2, 1, "9", 6, 7}, 42},
		{[]any{"3", 6, 6, 0, "5", 8, 5, "6", 2, "0"}, 41},
		{[]any{"1", "5", "8", 8, 9, 9, 2, "3"}, 45},
		{[]any{8, 0, 0, 8, 5, 7, 2, 3, 7, 8, 6, 7}, 61},
	}

	for _, test := range testCases {
		got := SumMix(test.arr)
		want := test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}
