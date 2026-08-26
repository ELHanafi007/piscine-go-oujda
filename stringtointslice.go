package piscine

func StringToIntSlice(str string) []int {
	slc := make([]int, 0)
	for _, r := range str {
		slc = append(slc, int(r))
	}
	return slc
}
