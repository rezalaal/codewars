package kata8

import "testing"

func TestCorrectTail(t *testing.T) {
	testCases := []struct {
		body string
		tail rune
		out  bool
	}{
		{"Fox", 'x', true},
		{"Rhino", 'o', true},
		{"Meerkat", 't', true},
	}

	for _, test := range testCases {
		got := CorrectTail(test.body, test.tail)
		want := test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}

		got = CorrectTail2(test.body, test.tail)
		want = test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}

		got = CorrectTail3(test.body, test.tail)
		want = test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}
