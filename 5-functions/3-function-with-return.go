// Exercise 5.03 – creating a checkNumbers function with return values

package main

import "fmt"

func checkNumbers(i int) (int, string) {
	switch {
	case i%2 == 0:
		return i, "Even"
	default:
		return i, "Odd"
	}
}

func main() {
	for i := 1; i <= 15; i++ {
		num, result := checkNumbers(i)
		fmt.Printf("Results: %d %s\n", num, result)
	}
}
