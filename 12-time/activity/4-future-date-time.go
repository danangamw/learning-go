//  Activity 12.04 – Calculating the future date and time

package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("The current time:", now.Format(time.ANSIC))

	future := now.Add(21966 * time.Second)
	fmt.Println("6 hours, 6 minutes and 6 seconds from now the time will be:", future.Format(time.ANSIC))
}
