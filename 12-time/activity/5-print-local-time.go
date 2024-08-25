//  Activity 12.05 – Printing the local time in different time zones

package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	nyLoc, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println(err)
	}

	laLoc, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		fmt.Println(err)
	}

	idLoc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println(err)
	}

	nyTime := now.In(nyLoc)
	laTime := now.In(laLoc)
	idTime := now.In(idLoc)

	fmt.Println("The local current time is:", now.Format(time.ANSIC))
	fmt.Println("The time in New York is:", nyTime.Format(time.ANSIC))
	fmt.Println("The time in Los Angeles is:", laTime.Format(time.ANSIC))
	fmt.Println("The time in Jakarta is:", idTime.Format(time.ANSIC))
}
