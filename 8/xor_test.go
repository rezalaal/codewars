package kata8

import "testing"

func TestXor(t *testing.T) {
	testCases := []struct {
		a   bool
		b   bool
		out bool
	}{
		{false, false, false}, // "false xor false should be false"
		{true, false, true},   // "true xor false should be true"
		{false, true, true},   // "false xor true should be true"
		{true, true, false},   // "true xor true should be false"
	}

	for _, test := range testCases {
		got := Xor(test.a, test.b)
		expected := test.out

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}

		got = Xor1(test.a, test.b)
		expected = test.out

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}

		got = Xor2(test.a, test.b)
		expected = test.out

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}
	}
}
