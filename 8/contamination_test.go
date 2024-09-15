package kata8

import "testing"

func TestContamination(t *testing.T) {
	testCases := []struct {
		text string
		char string
		out  string
	}{
		{"abc", "z", "zzz"},
		{"", "z", ""},
		{"abc", "", ""},
		{"_3ebzgh4", "&", "&&&&&&&&"},
		{"//case", " ", "      "},
	}

	for _, test := range testCases {
		got := Contamination(test.text, test.char)
		expected := test.out

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}

		got = Contamination1(test.text, test.char)
		expected = test.out

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}

		got = Contamination2(test.text, test.char)
		expected = test.out

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}

	}
}
