package piscine

func LoafOfBread(str string) string {
	result := ""
	count := 0
	skip := false

	for _, r := range str {
		if skip {
			skip = false
			continue
		}

		if r == ' ' {
			continue
		}

		result += string(r)
		count++

		if count == 5 {
			result += " "
			count = 0
			skip = true
		}
	}

	if count < 5 && count > 0 {
		return result + "\n"
	}

	if count == 0 && len(result) > 0 {
		result = result[:len(result)-1]
	}

	return result + "\n"
}
