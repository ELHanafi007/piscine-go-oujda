package main

import (
	"fmt"

	piscine "github.com/01-edu/piscine-go"
)

func main() {
	middle := piscine.Abort(2, 3, 8, 5, 7)
	fmt.Println(middle)
}
