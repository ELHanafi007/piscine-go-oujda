package piscine

func DescendAppendRange(max, min int) []int {
	slc := make([]int, 0)
	if min >= max {
		return slc
	}

	for i := max; i > min; i-- {
		slc = append(slc, i)
	}

	return slc
}
