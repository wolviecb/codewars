package stvtonato

import (
	"strings"
)

var nato = map[string]string{
	"a": "Alfa",
	"b": "Bravo",
	"c": "Charlie",
	"d": "Delta",
	"e": "Echo",
	"f": "Foxtrot",
	"g": "Golf",
	"h": "Hotel",
	"i": "India",
	"j": "Juliett",
	"k": "Kilo",
	"l": "Lima",
	"m": "Mike",
	"n": "November",
	"o": "Oscar",
	"p": "Papa",
	"q": "Quebec",
	"r": "Romeo",
	"s": "Sierra",
	"t": "Tango",
	"u": "Uniform",
	"v": "Victor",
	"w": "Whiskey",
	"x": "Xray",
	"y": "Yankee",
	"z": "Zulu",
}

//ToNato tonatizes
func ToNato(words string) string {

	s := ""
	l := strings.Join(strings.Fields(words), "")
	for _, v := range l {
		if isLetter(string(v)) {
			s += nato[strings.ToLower(string(v))] + " "
		} else {
			s += string(v)
		}
	}
	return strings.TrimSpace(s)
}

func isLetter(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}
