package piscine

func AppendRange(min, max int) []int {
	var slc []int
	if min > max || min == max {
		return slc
	}

	for i := min; i < max; i++ {
		slc = append(slc, i)
	}

	return slc
}
