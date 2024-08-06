package main

// Exercise 1.08 – changing multiple values at once

import "fmt"

func main() {
	query, limit, offset := "bat", 10, 0

	query, limit, offset = "ball", offset, 20

	fmt.Println(query, limit, offset)
}
