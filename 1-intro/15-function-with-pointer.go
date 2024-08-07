// Exercise 1.15 – function design with pointers

package main

import "fmt"

// pass by value
func add5Value(count int) {
	count += 5
	fmt.Println("add5Value    :", count)
}

// pass by pointer
func add5Point(count *int) {
	*count += 5
	fmt.Println("add5Point    ", *count)
}

func main() {
	var count int

	add5Value(count)

	fmt.Println("add5Value post: ", count)

	add5Point(&count)

	fmt.Println("add5Point post: ", count)
}
