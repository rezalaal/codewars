package kata8

import "testing"

func TestDerive(t *testing.T) {
	testCases := []struct {
		coefficient int
		exponent    int
		out         string
	}{
		{7, 8, "56x^7"},
		{5, 9, "45x^8"},
	}

	for _, test := range testCases {
		got := Derive(test.coefficient, test.exponent)
		want := test.out

		if got != want {
			t.Errorf("Expect %v but got %v", want, got)
		}
	}
}
