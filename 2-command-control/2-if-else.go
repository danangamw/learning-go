// Exercise 2.02 – using an if else statement

package main

import "fmt"

func main() {
	input := 4

	if input%2 == 0 {
		fmt.Println(input, "is even")
	} else {
		fmt.Println(input, "is odd")
	}
}
