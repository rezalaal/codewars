package kata7

import "testing"

func TestMaxRot(t *testing.T) {
	testCases := []struct{
		n int64
		out int64 
	}{
		{38458215, 85821534},
        {195881031, 988103115},
        {896219342, 962193428},
		{56789, 68957},
	}

	for _, test := range testCases {
		got := MaxRot(test.n)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}