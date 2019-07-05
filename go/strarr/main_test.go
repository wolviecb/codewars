package strarr

import (
	"testing"
)

func TestLongestConsec(t *testing.T) {
	tt := []struct {
		name string
		i    []string
		t    int
		o    string
	}{
		{
			"Words",
			[]string{"zone", "abigail", "theta", "form", "libe", "zas"},
			2,
			"abigailtheta",
		},
		{
			"Single K",
			[]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"},
			1,
			"oocccffuucccjjjkkkjyyyeehh",
		},
		{
			"Empty",
			[]string{},
			3,
			"",
		},
		{
			"Long strings",
			[]string{"itvayloxrp", "wkppqsztdkmvcuwvereiupccauycnjutlv", "vweqilsfytihvrzlaodfixoyxvyuyvgpck"},
			2,
			"wkppqsztdkmvcuwvereiupccauycnjutlvvweqilsfytihvrzlaodfixoyxvyuyvgpck",
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			r := LongestConsec(tc.i, tc.t)
			if r != tc.o {
				t.Fatalf("expected '%v', got '%v'", tc.o, r)
			}
		})
	}
}
