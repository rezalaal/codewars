package kata8

import "testing"

func TestStringToNumber(t *testing.T) {
	testCases := []struct {
		s   string
		out int
	}{
		{"1234", 1234},
		{"605", 605},
		{"1405", 1405},
		{"-7", -7},
	}

	for _, test := range testCases {
		got := StringToNumber(test.s)
		expected := test.out

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}
	}
}
