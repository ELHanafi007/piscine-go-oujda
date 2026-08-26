package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	list := make(map[string]int)
	currW := ""

	for _, s := range str {
		if isalpha(s) {
			currW += string(s)
		} else if s == ' ' {
			if currW != "" {
				list[currW]++
				currW = ""
			}
		}
	}

	if currW != "" {
		list[currW]++
	}

	return list
}

func isalpha(s rune) bool {
	if s >= 'a' && s <= 'z' {
		return true
	} else if s >= 'A' && s <= 'Z' {
		return true
	}
	return false
}
