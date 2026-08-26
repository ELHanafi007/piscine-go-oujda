package piscine

func Compact(ptr *[]string) int {
	i := 0
	for _, r := range *ptr {
		if r != "" {
			(*ptr)[i] = r
			i++
		}
	}
	*ptr = (*ptr)[:i]
	return i
}
