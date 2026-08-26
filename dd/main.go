package main

import (
	"fmt"

	piscine "github.com/01-edu/piscine-go"
)

const N = 6

func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(piscine.Join(toConcat, ":"))
}
