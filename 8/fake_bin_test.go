package kata8

import "testing"

func TestFakeBin(t *testing.T) {
	testCases := []struct {
		x   string
		out string
	}{
		{"45385593107843568", "01011110001100111"},
		{"509321967506747", "101000111101101"},
		{"366058562030849490134388085", "011011110000101010000011011"},
		{"15889923", "01111100"},
		{"800857237867", "100111001111"},
	}

	for _, test := range testCases {
		got := FakeBin(test.x)
		expected := test.out

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}

		got = FakeBin2(test.x)
		expected = test.out

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}

		got = FakeBin3(test.x)
		expected = test.out

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}

		got = FakeBin4(test.x)
		expected = test.out

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}
	}
}
