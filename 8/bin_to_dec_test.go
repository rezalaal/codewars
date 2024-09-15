package kata8

import "testing"

func TestBinToDec(t *testing.T) {
	testCases := []struct {
		input string
		expected int
	}{
		{"0",      0},
		{"1",      1},
		{"10",     2},
		{"11",     3},
		{"101010", 42},
		{"1001001",73},
	}
	for _, test :=range testCases {
		got := BinToDec(test.input)
		
		if got != test.expected {
			t.Errorf("BinToDec test expected %v but got %v", test.expected, got)
		}

		got = BinToDec2(test.input)

		if got != test.expected {
			t.Errorf("BinToDec test expected %v but got %v", test.expected, got)
		}

		got = BinToDec3(test.input)

		if got != test.expected {
			t.Errorf("BinToDec test expected %v but got %v", test.expected, got)
		}
	}
}