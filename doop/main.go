package main

import "os"

func atoi(s string) (int64, bool) {
	if len(s) == 0 {
		return 0, false
	}

	sign := int64(1)
	i := 0

	if s[0] == '-' {
		sign = -1
		i++
	} else if s[0] == '+' {
		i++
	}

	if i == len(s) {
		return 0, false
	}

	var n int64

	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}

		d := int64(s[i] - '0')

		if n > (9223372036854775807-d)/10 {
			return 0, false
		}

		n = n*10 + d
	}

	if sign == -1 {
		if n > 9223372036854775808 {
			return 0, false
		}
		return -n, true
	}

	return n, true
}

func printNbr(n int64) {
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
		if b > 0 && a > 9223372036854775807-b {
			return
		}
		if b < 0 && a < -9223372036854775808-b {
			return
		}
		printNbr(a + b)

	case "-":
		if b < 0 && a > 9223372036854775807+b {
			return
		}
		if b > 0 && a < -9223372036854775808+b {
			return
		}
		printNbr(a - b)

	case "*":
		if a != 0 && (a*b)/a != b {
			return
		}
		printNbr(a * b)

	case "/":
		if b == 0 {
			os.Stdout.Write([]byte("No division by 0\n"))
			return
		}
		if a == -9223372036854775808 && b == -1 {
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
