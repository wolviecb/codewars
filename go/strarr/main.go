package strarr

import (
	"strings"
)

//LongestConsec strings
func LongestConsec(strarr []string, k int) string {
	var (
		s string
	)
	n := len(strarr)
	ss := 0
	if n == 0 || k > n || k <= 0 {
		return ""
	}

	for i := range strarr {
		if i+k <= n {
			t := strings.Join(strarr[i:i+(k)], "")
			ts := len(t)
			if ts > ss {
				s = t
				ss = ts
			}
		}
	}
	return s
}
