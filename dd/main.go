package main

import (
	"fmt"

	piscine "github.com/01-edu/piscine-go"
)

const N = 6

func main() {
	slice := []string{"Pineapple", "Honey", "Mushroom", "Tea", "Pepper", "Milk"}
	fmt.Println(piscine.ShoppingListSort(slice))
}
