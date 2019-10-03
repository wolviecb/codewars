package stvtonato

import (
	"testing"
)

func TestInArray(t *testing.T) {
	tt := []struct {
		name string
		i    string
		o    string
	}{
		{
			"Simple",
			"If you can read.",
			"India Foxtrot Yankee Oscar Uniform Charlie Alfa November Romeo Echo Alfa Delta .",
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			r := ToNato(tc.i)
			if r != tc.o {
				t.Fatalf("expected '%v', got '%v'", tc.o, r)
			}
		})
	}
}
