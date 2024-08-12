// Activity 4.01 – Filling an array

package main

import "fmt"

func createArr() [10]int {
	var arr [10]int
	for i := 1; i <= 10; i++ {
		arr[i-1] = i
	}

	return arr
}

func main() {
	fmt.Println(createArr())
}
