// Activity 3.01 – Sales tax calculator

package main

import "fmt"

type Item struct {
	Cost         float32
	SalesTaxRate float32
}

func main() {
	var taxTotal float32 = 0
	salesData := map[string]Item{
		"Cake":   {Cost: 0.99, SalesTaxRate: 0.075},
		"Milk":   {Cost: 2.75, SalesTaxRate: 0.015},
		"Butter": {Cost: 0.87, SalesTaxRate: 0.02},
	}

	for _, value := range salesData {
		taxAmount := value.Cost * value.SalesTaxRate
		taxTotal += taxAmount
	}

	fmt.Println("Sales Tax Total: ", taxTotal)
}
