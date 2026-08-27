package piscine

func ReverseMenuIndex(menu []string) []string {
	result := make([]string, len(menu))

	for i := 0; i < len(menu); i++ {
		result[len(menu)-1-i] = menu[i]
	}

	return result
}
