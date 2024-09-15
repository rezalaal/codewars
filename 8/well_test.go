package kata8

import "testing"

func TestWell(t *testing.T) {
	testCases := []struct {
		x   []string
		out string
	}{
		{[]string{"good", "bad", "good", "good", "bad", "good", "bad", "bad", "good", "bad", "bad"}, "I smell a series!"},
		{[]string{"bad", "bad", "bad", "bad", "good", "good", "bad", "bad", "bad"}, "Publish!"},
		{[]string{"bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "good", "bad", "bad", "bad"}, "Publish!"},
		{[]string{"bad", "bad", "bad", "good", "bad", "bad", "good", "bad", "bad", "bad"}, "Publish!"},
		{[]string{"bad", "bad", "bad", "bad", "good", "bad", "bad"}, "Publish!"},
		{[]string{"bad", "bad"}, "Fail!"},
		{[]string{"bad", "bad", "bad", "bad", "bad"}, "Fail!"},
		{[]string{"bad", "bad", "bad", "bad", "good", "bad"}, "Publish!"},
		{[]string{"bad"}, "Fail!"},
		{[]string{"bad", "bad", "bad", "good", "bad", "good", "good", "bad", "bad", "bad", "bad", "good", "good"}, "I smell a series!"},
		{[]string{"bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "good", "bad", "bad", "bad", "bad", "good"}, "Publish!"},
		{[]string{"good", "bad", "bad", "bad", "bad", "bad", "bad", "good", "bad", "bad", "good", "bad"}, "I smell a series!"},
		{[]string{"bad", "bad", "bad", "bad", "bad", "good", "bad", "good", "good", "good", "bad", "bad", "good"}, "I smell a series!"},
		{[]string{"bad", "good", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad"}, "Publish!"},
		{[]string{"bad", "bad", "bad", "good", "bad", "bad"}, "Publish!"},
		{[]string{"good", "bad", "bad", "bad", "bad", "good", "bad"}, "Publish!"},
		{[]string{"good"}, "Publish!"},
		{[]string{"good", "good"}, "Publish!"},
		{[]string{"bad", "bad", "bad", "good", "bad", "bad", "bad", "good", "bad", "bad", "bad", "bad", "bad"}, "Publish!"},
		{[]string{"good", "bad", "good", "good", "bad", "bad", "bad"}, "I smell a series!"},
		{[]string{"bad", "bad", "bad", "bad", "bad"}, "Fail!"},
		{[]string{"bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "good", "bad", "bad", "bad", "good"}, "Publish!"},
		{[]string{"bad", "good", "bad", "bad", "good", "bad", "good", "bad", "bad", "bad", "bad", "bad", "bad", "bad"}, "I smell a series!"},
		{[]string{"bad", "good"}, "Publish!"},
		{[]string{"bad", "good", "bad", "bad"}, "Publish!"},
		{[]string{"good", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "good"}, "Publish!"},
		{[]string{"bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad"}, "Fail!"},
		{[]string{"bad", "bad", "good", "bad", "bad", "good", "bad", "bad", "bad", "bad", "bad", "good", "good", "bad", "good", "bad"}, "I smell a series!"},
		{[]string{"bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad"}, "Fail!"},
		{[]string{"bad", "bad", "bad", "bad", "good", "bad", "bad", "bad"}, "Publish!"},
		{[]string{"bad", "bad", "good", "bad", "bad", "good", "bad", "good", "bad", "bad", "bad"}, "I smell a series!"},
		{[]string{"bad", "bad", "bad", "bad", "bad"}, "Fail!"},
		{[]string{"good", "bad", "bad", "bad", "bad", "bad", "good", "good", "bad", "bad", "bad", "bad", "good", "bad", "bad"}, "I smell a series!"},
		{[]string{"bad", "bad", "bad", "bad", "bad", "good", "bad", "bad"}, "Publish!"},
	}

	for _, test := range testCases {
		got := Well(test.x)
		want := test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}
