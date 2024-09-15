package kata8

import "testing"

func TestDigitize(t *testing.T) {
	testCases := []struct {
		n int
		out []int
	}{
		{35231, []int{1,3,2,5,3}},
		{0,     []int{0}},
	}

	for _, test :=range testCases {
		got := Digitize(test.n)
		expected := test.out
		for i, n :=range got {
			if expected[i] != n {
				t.Errorf("Expected %v but got %v", expected[i], n)
			}
		}

		got = Digitize2(test.n)
		expected = test.out
		for i, n :=range got {
			if expected[i] != n {
				t.Errorf("Expected %v but got %v", expected[i], n)
			}
		}

		got = Digitize3(test.n)
		expected = test.out
		for i, n :=range got {
			if expected[i] != n {
				t.Errorf("Expected %v but got %v", expected[i], n)
			}
		}
	}
}