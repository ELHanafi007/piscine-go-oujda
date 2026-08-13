package piscine

func UltimateDivMod(a *int, b *int) {
	tmp := *a / *b
	tmp2 := *a % *b
	*a = tmp
	*b = tmp2
}
