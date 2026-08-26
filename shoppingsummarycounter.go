package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	list := make(map[string]int)
	currW := ""

	for _, r := range str {
		if isalpha(r) {
			currW += string(r)
		} else if r == ' ' {
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

func isalpha(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	} else if r >= 'A' && r <= 'Z' {
		return true
	}
	return false
}
