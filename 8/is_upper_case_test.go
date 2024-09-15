package kata8

import "testing"

func TestIsUpperCase(t *testing.T) {
	
	testCases := []struct {
		s   MyString
		out bool
	}{
		{"a", false},
		{"b", false},
		{"c", false},
		{"d", false},
		{"e", false},
		{"f", false},
		{"g", false},
		{"h", false},
		{"i", false},
		{"j", false},
		{"k", false},
		{"l", false},
		{"m", false},
		{"n", false},
		{"o", false},
		{"p", false},
		{"q", false},
		{"r", false},
		{"s", false},
		{"t", false},
		{"u", false},
		{"v", false},
		{"w", false},
		{"x", false},
		{"y", false},
		{"z", false},
		{"abcdefghijklmnopqrstuvwxyz", false},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYz", false},
		{"false", false},
		{"true", false},
		{"False", false},
		{"True", false},
		{"WHAT DOES THE FOX SAY", true},
		{"HTML CSS JAVASCRIPT PYTHON C PERL LISP JAVA GO RUBY NODEJS RUST SCALA", true},
	}

	for _, test := range testCases {
		got := test.s.IsUpperCase()
		want := test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}

		got = test.s.IsUpperCase2()
		want = test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}

		got = test.s.IsUpperCase3()
		want = test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}
