package cube

import (
	"math"
)

// FindNb find cubes
func FindNb(m int) int {
	for n := 0; ; n++ {
		if m > 0 {
			c := math.Pow(float64(n+1), float64(3))
			m = m - int(c)
		} else if m == 0 {
			return n
		} else if m < 0 {
			return -1
		}
	}
}
