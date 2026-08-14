package piscine


func BasicAtoi(s string) int {

	i := 0
	for _, r := range s {
		
		i = (i *10) + int(r - '0')
	}
	return i


}
