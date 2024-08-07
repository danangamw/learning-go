// Activity 1.03 – message bug

package main

import "fmt"

// problem in variable message, that can't be accessed on condition level

func main() {
	count := 5
	var message string

	if count > 5 {
		message = "Greater than 5"
	} else {
		message = "Not greater than 5"
	}

	fmt.Println(message)
}
