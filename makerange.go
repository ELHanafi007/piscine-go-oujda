package piscine

func MakeRange(min, max int) []int {
	if min > max || min == max {
		return nil
	}
	slc := make([]int, max-min)
	for i := 0; i < len(slc); i++ {
		slc[i] = min + i
	}
	return slc
}
