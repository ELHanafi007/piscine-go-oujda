package main

import (
	"fmt"

	piscine "github.com/01-edu/piscine-go"
)

func main() {
	steps := piscine.CollatzCountdown(12)
	fmt.Println(steps)
}
