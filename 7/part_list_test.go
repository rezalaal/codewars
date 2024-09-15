package kata7

import "testing"

func TestPartList(t *testing.T) {
	testCases := []struct {
		arr []string
		out string
	}{
		{[]string{"I", "wish", "I", "hadn't", "come"}, "(I, wish I hadn't come)(I wish, I hadn't come)(I wish I, hadn't come)(I wish I hadn't, come)"},
		{[]string{"cdIw", "tzIy", "xDu", "rThG"}, "(cdIw, tzIy xDu rThG)(cdIw tzIy, xDu rThG)(cdIw tzIy xDu, rThG)"},
	}

	for _, test := range testCases {
		got := PartList(test.arr)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = PartList2(test.arr)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = PartList3(test.arr)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
