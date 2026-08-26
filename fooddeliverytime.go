package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	orderT := 0
	if order == "burger" {
		orderT += 15
	}
	if order == "chips" {
		orderT += 10
	}
	if order == "nuggets" {
		orderT += 12
	}
	return orderT
}
