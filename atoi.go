package piscine

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	sn := 1
	i := 0

	if s[0] == '+' {
		i++
	} else if s[0] == '-' {
		sn = -1
		i++
	}

	if i == len(s) {
		return 0
	}

	res := 0

	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}

		num := int(s[i] - '0')
		res = res*10 + num
	}

	return res * sn
}
