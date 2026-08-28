package piscine

func Map(f func(int) bool, a []int) []bool {
	result := []bool{}

	for _, n := range a {
		result = append(result, f(n))
	}

	return result
}
