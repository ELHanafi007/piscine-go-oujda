package piscine

func Max(a []int) int {
	maksi := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > maksi {
			maksi = a[i]
		}
	}
	return maksi
}
