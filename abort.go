package piscine

func Abort(a, b, c, d, e int) int {
	numbers := [5]int{a, b, c, d, e}

	for i := 0; i < 5; i++ {
		for j := i + 1; j < 5; j++ {
			if numbers[i] > numbers[j] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}

	return numbers[2]
}
