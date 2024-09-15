package kata7

import "testing"

func TestSimpleStringCharacters(t *testing.T) {
	testCases := []struct {
		s   string
		out []int
	}{
		{"Codewars@codewars123.com", []int{1, 18, 3, 2}},
		{"bgA5<1d-tOwUZTS8yQ", []int{7, 6, 3, 2}},
		{"P*K4%>mQUDaG$h=cx2?.Czt7!Zn16p@5H", []int{9, 9, 6, 9}},
		{"RYT'>s&gO-.CM9AKeH?,5317tWGpS<*x2ukXZD", []int{15, 8, 6, 9}},
		{"$Cnl)Sr<7bBW-&qLHI!mY41ODe", []int{10, 7, 3, 6}},
	}

	for _, test := range testCases {
		got := SimpleStringCharacters(test.s)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = SimpleStringCharacters2(test.s)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = SimpleStringCharacters3(test.s)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = SimpleStringCharacters4(test.s)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
