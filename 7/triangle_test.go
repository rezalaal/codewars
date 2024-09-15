package kata7

import "testing"

func TestTriangle(t *testing.T){
	testCases := []struct{
		row string
		out rune
	}{
		{"GB", 'R'},
		{"RRR", 'R'},
		{"RGBG", 'B'},
		{"RBRGBRB", 'G'},
		{"RBRGBRBGGRRRBGBBBGG", 'G'},
		{"B", 'B'},
	}

	for _, test := range testCases {
		got := Triangle(test.row)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = Triangle2(test.row)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = Triangle3(test.row)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = Triangle4(test.row)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = Triangle5(test.row)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = Triangle6(test.row)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
