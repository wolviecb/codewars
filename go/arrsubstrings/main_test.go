package arrsubstrings

import (
	"testing"
)

func testEq(a, b []string) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
func TestInArray(t *testing.T) {
	tt := []struct {
		name string
		i1   []string
		i2   []string
		o    []string
	}{
		{
			"Simple",
			[]string{"live", "arp", "strong"},
			[]string{"lively", "alive", "harp", "sharp", "armstrong"},
			[]string{"arp", "live", "strong"},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			r := InArray(tc.i1, tc.i2)
			if !testEq(r, tc.o) {
				t.Fatalf("expected '%v', got '%v'", tc.o, r)
			}
		})
	}
}
