package arrsum

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	tt := []struct {
		name string
		i    []int
		t    int
		o    [2]int
	}{
		{
			"Simple",
			[]int{1, 2, 3},
			4,
			[2]int{0, 2},
		},
		{
			"Bigger",
			[]int{1234, 5678, 9012},
			14690,
			[2]int{1, 2},
		},
		{
			"Smaller",
			[]int{2, 2, 3},
			4,
			[2]int{0, 1},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			r := TwoSum(tc.i, tc.t)
			if r != tc.o {
				t.Fatalf("expected '%v', got '%v'", tc.o, r)
			}
		})
	}
}
