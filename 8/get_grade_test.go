package kata8

import "testing"

func TestGetGrade(t *testing.T) {
	testCases := []struct {
		a   int
		b   int
		c   int
		out rune
	}{
		{95, 90, 93, 'A'},
		{100, 85, 96, 'A'},
		{92, 93, 94, 'A'},
		{70, 70, 100, 'B'},
		{82, 85, 87, 'B'},
		{84, 79, 85, 'B'},
		{70, 70, 70, 'C'},
		{75, 70, 79, 'C'},
		{60, 82, 76, 'C'},
		{65, 70, 59, 'D'},
		{66, 62, 68, 'D'},
		{58, 62, 70, 'D'},
		{44, 55, 52, 'F'},
		{48, 55, 52, 'F'},
		{58, 59, 60, 'F'},
	}

	for _, test := range testCases {
		got := GetGrade(test.a, test.b, test.c)
		want := test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}
