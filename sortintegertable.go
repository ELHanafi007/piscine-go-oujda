package piscine

func SortIntegerTable(table []int) {
	for index := 1; index < len(table); index++ {
		now := table[index]
		prev := index - 1

		for prev >= 0 && table[prev] > now {
			table[prev+1] = table[prev]
			prev = prev - 1
		}
		table[prev+1] = now
	}
}
