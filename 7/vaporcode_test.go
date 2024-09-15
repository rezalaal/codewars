package kata7

import "testing"

func TestVaporcode(t *testing.T) {
	testCases := []struct{
		s string
		out string
	}{
		{"Lets go to the movies", "L  E  T  S  G  O  T  O  T  H  E  M  O  V  I  E  S"},
       	{"Why isnt my code working", "W  H  Y  I  S  N  T  M  Y  C  O  D  E  W  O  R  K  I  N  G"},
	}

	for _, test := range testCases {
		got := Vaporcode(test.s)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = Vaporcode1(test.s)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}