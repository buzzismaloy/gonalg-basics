package main

import "fmt"

func countOrderCost(order []string, prices map[string]int) int {
	total := 0

	for _, item := range order {
		price, ok := prices[item]

		if ok {
			total += price
		}
	}

	return total
}

func main() {
	var number int = 500
	priceList := map[string]int{"bread": 50, "milk": 100, "butter": 200, "sausage": 500, "salt": 20, "cucumbers": 200, "cheese": 600, "ham": 700, "pork": 900, "tomatoes": 250, "fish": 300, "hamon": 1500}
	order := []string{"bread", "pork", "cheese", "cucumbers"}

	for food, price := range priceList {
		if price > number {
			fmt.Printf("%s is more expensive than 500 rubles\n", food)
		}
	}

	totalPrice := countOrderCost(order, priceList)

	fmt.Printf("\n\nThe cost of the order is %d", totalPrice)

}
