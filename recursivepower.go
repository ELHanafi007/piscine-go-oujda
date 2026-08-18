package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	sum := 1
	if sum <= power {
		sum = nb * RecursivePower(nb, (power-1))
	}
	return sum
}
