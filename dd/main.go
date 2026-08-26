package main

import (
	"fmt"

	piscine "github.com/01-edu/piscine-go"
)

func main() {
	fmt.Println(piscine.FoodDeliveryTime("burger"))
	fmt.Println(piscine.FoodDeliveryTime("chips"))
	fmt.Println(piscine.FoodDeliveryTime("nuggets"))
	fmt.Println(piscine.FoodDeliveryTime("burger") + piscine.FoodDeliveryTime("chips") + piscine.FoodDeliveryTime("nuggets"))
}
