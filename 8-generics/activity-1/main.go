// Activity 8.01 – a minimum value

package main

import "fmt"

func findMinGeneric[Num int | float64](nums []Num) Num {
	if len(nums) <= 0 {
		return 0
	}

	minNumber := nums[0]
	for _, num := range nums {
		if num < minNumber {
			minNumber = num
		}
	}

	return minNumber
}

func main() {
	nums := []float64{1.1, 32.1, 5.1, 8.1, 10.1, 11.1}
	fmt.Println("min value:", findMinGeneric(nums))
}
