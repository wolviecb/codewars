package cube

// FindNb find cubes
func FindNb(m int) int {
	for n := 0; ; n++ {
		if m > 0 {
			m -= (n + 1) * (n + 1) * (n + 1)
		} else if m == 0 {
			return n
		} else if m < 0 {
			return -1
		}
	}
}
