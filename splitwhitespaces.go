package piscine

func SplitWhiteSpaces(s string) []string {
	var res []string
	currW := ""

	for _, r := range s {
		if r != ' ' && r != '\t' && r != '\n' {
			currW += string(r)
		} else {
			if currW != "" {
				res = append(res, currW)
				currW = ""
			}
		}
	}

	if currW != "" {
		res = append(res, currW)
	}

	return res
}
