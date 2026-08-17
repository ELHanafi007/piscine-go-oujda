package piscine

/*
import (

	"fmt"

)
*/
func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 20{
		return 0
	}

	r := 1

	for i := 1; i <= nb; i++ {
		r *= i
	}

	return r
}

/*func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}*/
