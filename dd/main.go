package main

import (
	"fmt"

	piscine "github.com/01-edu/piscine-go"
)

func main() {
	summary := "Burger Water Carrot Coffee Water Water Chips Carrot Carrot Burger Carrot Water"
	for index, element := range piscine.ShoppingSummaryCounter(summary) {
		fmt.Println(index, "=>", element)
	}
}
