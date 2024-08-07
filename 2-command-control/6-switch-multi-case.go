// Exercise 2.06 – switch statements and multiple case values

package main

import (
	"fmt"
	"time"
)

func main() {
	dayBorn := time.Sunday

	switch dayBorn {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Born on the weekday")
	case time.Saturday, time.Sunday:
		fmt.Println("Born on the weekend")
	default:
		fmt.Println("Error, day born not valid!")
	}
}
