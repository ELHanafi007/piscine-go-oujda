package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) < 2 {
		return true
	}

	direction := f(a[0], a[1])

	for i := 1; i < len(a)-1; i++ {
		if direction > 0 && f(a[i], a[i+1]) < 0 {
			return false
		}
		if direction < 0 && f(a[i], a[i+1]) > 0 {
			return false
		}
	}

	return true
}
