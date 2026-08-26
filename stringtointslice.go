package piscine

func StringToIntSlice(str string) []int {
	var slc []int

	for _, r := range str {
		slc = append(slc, int(r))
	}
	return slc
}
