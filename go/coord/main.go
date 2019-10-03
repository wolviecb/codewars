package coord

import (
	"strconv"
	"strings"
)

const alpha = "abcdefghijklmnopqrstuvwxyz"

func hasAlpha(s string) bool {
	for _, char := range s {
		if strings.Contains(alpha, strings.ToLower(string(char))) {
			return true
		}
	}
	return false
}

//IsValidCoordinates validates
func IsValidCoordinates(coordinates string) bool {
	c := strings.Split(coordinates, ",")
	if len(c) != 2 {
		return false
	}

	for _, v := range c {
		if hasAlpha(v) {
			return false
		}
	}
	var vlt bool
	lt, err := strconv.ParseFloat(c[0], 64)
	if err != nil {
		return false
	}
	if (lt >= -90) && (lt <= 90) {
		vlt = true
	}

	var vlg bool
	lg, err := strconv.ParseFloat(strings.TrimSpace(c[1]), 64)
	if err != nil {
		return false
	}
	if (lg >= -180) && (lg <= 180) {
		vlg = true
	}

	if vlt && vlg {
		return true
	}
	return false
}
