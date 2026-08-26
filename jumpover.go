package piscine

func JumpOver(str string) string {
	str2 := ""
	if str == "" {
		return "\n"
	}
	for i := 0; i < len(str); i++ {
		if (i+1)%3 == 0 {
			str2 += string(str[i])
		}
	}
	return str2 + "\n"
}
