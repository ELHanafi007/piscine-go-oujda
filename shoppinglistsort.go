package piscine

func ShoppingListSort(slice []string) []string {
	for index := 1; index < len(slice); index++ {
		now := slice[index]
		prev := index - 1

		for prev >= 0 && len(slice[prev]) > len(now) {
			slice[prev+1] = slice[prev]
			prev--
		}

		slice[prev+1] = now
	}

	return slice
}
