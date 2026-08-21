package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	upper := false

	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	for _, arg := range args {
		n := 0
		valid := true

		if len(arg) == 0 {
			valid = false
		}

		for _, r := range arg {
			if r < '0' || r > '9' {
				valid = false
				break
			}
			n = n*10 + int(r-'0')
		}

		if !valid || n < 1 || n > 26 {
			z01.PrintRune(' ')
			continue
		}

		r := rune('a' + n - 1)
		if upper {
			r = rune('A' + n - 1)
		}

		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
