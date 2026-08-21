package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[0]
	for i := len(arg) - 1; i >= 0; i-- {
		if arg[i] == '/' {
			arg = arg[i+1:]
			break
		}
	}
	for _, r := range arg {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
