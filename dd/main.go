package main

import (
	"fmt"

	piscine "github.com/01-edu/piscine-go"
)

const N = 6

func main() {
	fmt.Print(piscine.LoafOfBread("deliciousbread"))
	fmt.Print(piscine.LoafOfBread("This is a loaf of bread"))
	fmt.Print(piscine.LoafOfBread("loaf"))
}
