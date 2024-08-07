// Exercise 1.13 – getting a pointer

package main

import (
	"fmt"
	"time"
)

func main() {
	// Declare a pointer using a var statement:
	var count1 *int

	count2 := new(int)

	countTemp := 5

	count3 := &countTemp
	// count1 = &countTemp

	countTemp = 2

	//  It’s possible to create a pointer from some types without a temporary variable. Here, we’re using
	// our trusty time struct:
	t := &time.Time{}

	fmt.Printf("count1 : %#v\n", count1)
	fmt.Printf("count2 : %#v\n", count2)
	fmt.Printf("count3 : %#v\n", count3)
	fmt.Printf("time : %#v\n", t)

}
