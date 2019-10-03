package coord

import (
	"testing"
)

func TestIsValidCoordinates(t *testing.T) {
	tt := []struct {
		name string
		i    string
		o    bool
	}{
		{
			"True One",
			"-23, 25",
			true,
		},
		{
			"True Two",
			"4, -3",
			true,
		},
		{
			"True Three",
			"3, -4",
			true,
		},
		{
			"True Four",
			"24.53525235, 23.45235",
			true,
		},
		{
			"True Five",
			"04, -23.234235",
			true,
		},
		{
			"True Six",
			"43.91343345, 143",
			true,
		},
		{
			"False One",
			"23.234, - 23.4234",
			false,
		},
		{
			"False Two",
			"2342.43536, 34.324236",
			false,
		},
		{
			"False Three",
			"N23.43345, E32.6457",
			false,
		},
		{
			"False Four",
			"99.234, 12.324",
			false,
		},
		{
			"False Five",
			"23.245, 1e1",
			false,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			r := IsValidCoordinates(tc.i)
			if r != tc.o {
				t.Fatalf("expected '%v', got '%v'", tc.o, r)
			}
		})
	}
}
