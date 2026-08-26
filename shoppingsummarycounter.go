package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	mymap := make(map[string]int)
	word := ""
	for _, i := range str {
		if i == ' ' {
			mymap[word]++
			word = ""
		} else {
			word += string(i)
		}
	}
	mymap[word]++
	return mymap
}
