package kata8

import "testing"

func TestRepeatStr(t *testing.T) {
	testCases := []struct {
		repetitions int
		value       string
		out         string
	}{
		{4, "a", "aaaa"},
		{3, "hello ", "hello hello hello "},
		{2, "abc", "abcabc"},
	}

	for _, test := range testCases {
		got := RepeatStr(test.repetitions, test.value)
		want := test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}

		got = RepeatStr2(test.repetitions, test.value)
		want = test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}
