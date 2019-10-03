package arrsubstrings

import (
	"sort"
	"strings"
)

//InArray arrays
func InArray(array1 []string, array2 []string) []string {
	r := []string{}
	for _, a := range array1 {
		if contains(array2, a) {
			r = append(r, a)
		}
	}
	sort.Strings(r)
	return unique(r)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if strings.Contains(a, e) {
			return true
		}
	}
	return false
}

func unique(ss []string) []string {
	k := make(map[string]bool)
	l := []string{}
	for _, v := range ss {
		if _, value := k[v]; !value {
			k[v] = true
			l = append(l, v)
		}
	}
	return l
}
