// Exercise 2.01 – a simple if statement

package main

import "fmt"

func main() {
	input := 6

	if input%2 == 0 {
		fmt.Println(input, "is even")
	}

	if input%2 != 0 {
		fmt.Println(input, "is odd")
	}
}
