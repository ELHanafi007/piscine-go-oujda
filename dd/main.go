package main

import (
	"fmt"

	piscine "github.com/01-edu/piscine-go"
)

const N = 6

func main() {
	fmt.Println(piscine.ReverseMenuIndex([]string{"desserts", "mains", "drinks", "starters"}))
}
