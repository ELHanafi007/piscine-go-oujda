package piscine


func UltimateDivMod(a *int, b *int) {

	tmp := *a / *b
	tmp2 := *a % *b
	*a = tmp
	*b =tmp2
}
/*func main() {
	a := 13
	b := 2
	UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}/*