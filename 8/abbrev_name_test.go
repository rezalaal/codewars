package kata8

import "testing"

func TestAbbrevName(t *testing.T) {
	testCases := []struct {
        input    string
        expected string
    }{
        {"Sam Harris", "S.H"},
        {"Patrick Feenan", "P.F"},
        {"Evan Cole", "E.C"},
        {"P Favuzzi", "P.F"},
        {"David Mendieta", "D.M"},
    }

    for _, tc := range testCases {
        t.Run(tc.input, func(t *testing.T) {
            result := AbbrevName(tc.input)
            if result != tc.expected {
                t.Errorf("Expected %v but got %v", tc.expected, result)
            }
        })
    }
}
