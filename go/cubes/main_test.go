package cube

import (
	"testing"
)

func TestFindNb(t *testing.T) {
	tt := []struct {
		name string
		i    int
		o    int
	}{
		{
			"Valid Cube",
			4183059834009,
			2022,
		},
		{
			"Invalid Cube",
			24723578342962,
			-1,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			r := FindNb(tc.i)
			if r != tc.o {
				t.Fatalf("expected '%v', got '%v'", r, tc.o)
			}
		})
	}
}
