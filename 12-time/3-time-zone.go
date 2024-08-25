//  Exercise 12.03 – What is the time in your zone?

package main

import (
	"fmt"
	"time"
)

func timeDiff(timezone string) (string, string) {
	current := time.Now()
	remoteZone, err := time.LoadLocation(timezone)
	if err != nil {
		fmt.Println(err)
		return "", ""
	}

	remoteTime := current.In(remoteZone)
	fmt.Print("The current time is:", current.Format(time.ANSIC))
	fmt.Println("the timezone:", timezone, "time is:", remoteTime)
	return current.Format(time.ANSIC), remoteTime.Format(time.ANSIC)
}

func main() {
	fmt.Println(timeDiff("Asia/Jakarta"))
	fmt.Println(timeDiff("America/New_York"))
}
