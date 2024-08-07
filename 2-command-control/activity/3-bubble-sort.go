// Activity 2.03 – bubble sort

package main

import (
	"fmt"
)

func main() {
	// var numList []int = []int{5, 8, 2, 4, 0, 1, 3, 7, 9, 6}
	// var sortedNumList []int = make([]int, len(numList))
	// copy(sortedNumList, numList)
	// sort.Ints(sortedNumList)

	// fmt.Println("Before:", numList)
	// fmt.Println("After:", sortedNumList)

	numbers := []int{5, 8, 2, 4, 0, 1, 3, 7, 9, 6}
	fmt.Println("Before:", numbers)

	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	fmt.Println("After:", numbers)
}
