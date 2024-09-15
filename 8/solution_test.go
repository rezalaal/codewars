package kata8

import "testing"

func TestSolution(t *testing.T) {
	testCases := []struct {
		word string
		out  string
	}{
		{"world", "dlrow"},
		{"word", "drow"},
	}

	for _, test := range testCases {
		got := Solution(test.word)
		want := test.out

		if got != want {
			t.Errorf("Expected %v but %v", want, got)
		}

		got = Solution2(test.word)
		want = test.out

		if got != want {
			t.Errorf("Expected %v but %v", want, got)
		}
		
		got = Solution3(test.word)
		want = test.out

		if got != want {
			t.Errorf("Expected %v but %v", want, got)
		}
	}
}
