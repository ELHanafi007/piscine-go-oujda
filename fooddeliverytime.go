package piscine

type food struct {
	name string
	preptime int
	
}

func FoodDeliveryTime(order string) int {
	foods := []food{
		{"burger", 15},
		{"chips", 10},
		{"nuggets", 12},
	}
	for _, f := range foods{
		if order == f.name{
			return f.preptime
		}
	}
	
	return 404
}
