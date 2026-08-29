package main

import (
	"os"
)

func atoi(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}

	sign := 1
	i := 0

	if s[0] == '-' {
		sign = -1
		i++
	}

	if i == len(s) {
		return 0, false
	}

	n := 0

	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}

		n = n*10 + int(s[i]-'0')
	}

	return n * sign, true
}

func printNbr(n int) {
	if n < 0 {
		os.Stdout.Write([]byte("-"))
		n = -n
	}

	if n >= 10 {
		printNbr(n / 10)
	}

	os.Stdout.Write([]byte{byte(n%10) + '0'})
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	a, ok := atoi(os.Args[1])
	if !ok {
		return
	}

	b, ok := atoi(os.Args[3])
	if !ok {
		return
	}

	switch os.Args[2] {
	case "+":
		printNbr(a + b)
	case "-":
		printNbr(a - b)
	case "*":
		printNbr(a * b)
	case "/":
		if b == 0 {
			os.Stdout.Write([]byte("No division by 0\n"))
			return
		}
		printNbr(a / b)
	case "%":
		if b == 0 {
			os.Stdout.Write([]byte("No modulo by 0\n"))
			return
		}
		printNbr(a % b)
	default:
		return
	}

	os.Stdout.Write([]byte("\n"))
}
