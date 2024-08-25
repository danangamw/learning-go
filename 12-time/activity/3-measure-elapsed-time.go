//  Activity 12.03 – Measuring elapsed time

package main

import (
	"fmt"
	"time"
)

func elapsedTime(start time.Time, end time.Time) {
	elapsed := end.Sub(start)

	fmt.Println("The execution took exactly ", elapsed.Seconds(), " seconds!")
}

func main() {
	start := time.Now()
	time.Sleep(2 * time.Second)
	end := time.Now()
	elapsedTime(start, end)
}
