package kata8

import "testing"

func TestHowManyDalmatians(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{"26 dogs", 26, "More than a handful!"},
		{"8 dogs", 8, "Hardly any"},
		{"14 dogs", 14, "More than a handful!"},
		{"80 dogs", 80, "Woah that's a lot of dogs!"},
		{"100 dogs", 100, "Woah that's a lot of dogs!"},
		{"50 dogs", 50, "More than a handful!"},
		{"10 dogs", 10, "Hardly any"},
		{"101 dogs", 101, "101 DALMATIONS!!!"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got1 := HowManyDalmatians(test.input)
			if got1 != test.expected {
				t.Errorf("HowManyDalmatians(%d): expected %v but got %v", test.input, test.expected, got1)
			}

			got2 := HowManyDalmatians1(test.input)
			if got2 != test.expected {
				t.Errorf("HowManyDalmatians1(%d): expected %v but got %v", test.input, test.expected, got2)
			}
		})
	}
}
