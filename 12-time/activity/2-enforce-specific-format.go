// Activity 12.02 – Enforcing a specific format of date and time

package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	now := time.Date(2024, 8, 25, 20, 15, 21, 123121, time.UTC)

	hours := strconv.Itoa(now.Hour())
	minutes := strconv.Itoa(now.Minute())
	seconds := strconv.Itoa(now.Second())

	year := strconv.Itoa(now.Year())
	month := strconv.Itoa(int(now.Month()))
	day := strconv.Itoa(now.Day())

	fmt.Printf("%s:%s:%s %s/%s/%s", hours, minutes, seconds, year, month, day)
}
