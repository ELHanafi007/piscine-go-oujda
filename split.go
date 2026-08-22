package piscine

func Split(s, sep string) []string {
	var res []string
	currW := ""

	for i := 0; i < len(s); i++ {
		match := true

		if i+len(sep) > len(s) {
			match = false
		} else {
			for j := 0; j < len(sep); j++ {
				if s[i+j] != sep[j] {
					match = false
					break
				}
			}
		}

		if match {
			res = append(res, currW)
			currW = ""
			i += len(sep) - 1
		} else {
			currW += string(s[i])
		}
	}

	res = append(res, currW)

	return res
}