package main

import (
	"os"
	"piscine"

	"github.com/01-edu/z01"
)

func printNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n >= 10 {
		printNbr(n / 10)
	}
	z01.PrintRune(rune(n%10) + '0')
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	a := piscine.Atoi(os.Args[1])
	b := piscine.Atoi(os.Args[3])

	switch os.Args[2] {
	case "+":
		printNbr(a + b)
	case "-":
		printNbr(a - b)
	case "*":
		printNbr(a * b)
	case "/":
		if b == 0 {
			return
		}
		printNbr(a / b)
	case "%":
		if b == 0 {
			return
		}
		printNbr(a % b)
	default:
		return
	}

	z01.PrintRune('\n')
}
