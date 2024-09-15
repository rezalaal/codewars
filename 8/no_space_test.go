package kata8

import "testing"

func TestNoSpace(t *testing.T) {
	testCases := []struct {
		word string
		out  string
	}{
		{"8 j 8   mBliB8g  imjB8B8  jl  B", "8j8mBliB8gimjB8B8jlB"},
		{"8 8 Bi fk8h B 8 BB8B B B  B888 c hl8 BhB fd", "88Bifk8hB8BB8BBBB888chl8BhBfd"},
		{"8aaaaa dddd r     ", "8aaaaaddddr"},
		{"jfBm  gk lf8hg  88lbe8 ", "jfBmgklf8hg88lbe8"},
		{"8j aam", "8jaam"},
	}
	for _, test := range testCases {
		got := NoSpace(test.word)
		want := test.out
		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}

		got = NoSpace2(test.word)
		want = test.out
		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}

		got = NoSpace3(test.word)
		want = test.out
		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}
