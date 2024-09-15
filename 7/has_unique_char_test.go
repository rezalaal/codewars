package kata7

import "testing"

func TestHasUniqueChar(t *testing.T) {
	testCases := []struct {
		str string
		out bool
	}{
		{"  nAa", false},
		{"abcde", true},
		{"++-", false},
		{"AaBbC", true},
	}

	for _, test := range testCases {
		got := HasUniqueChar(test.str)
		want := test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}

		got = HasUniqueChar2(test.str)
		want = test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}

		got = HasUniqueChar3(test.str)
		want = test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}

		got = HasUniqueChar4(test.str)
		want = test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}

		got = HasUniqueChar5(test.str)
		want = test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}
