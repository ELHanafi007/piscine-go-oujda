package piscine

func DescendAppendRange(max, min int) []int {
	slc := []int{}

	if min >= max {
		return slc
	}

	for i := max; i > min; i-- {
		slc = append(slc, i)
	}

	return slc
}
