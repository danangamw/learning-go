//  Activity 12.01 – Formatting a date according to user requirements

package main

import (
	"fmt"
	"strconv"
	"time"
)

func getTime(datetime time.Time) string {
	hours := strconv.Itoa(datetime.Hour())
	minutes := strconv.Itoa(datetime.Minute())
	seconds := strconv.Itoa(datetime.Second())
	year := strconv.Itoa(datetime.Year())
	month := strconv.Itoa(int(datetime.Month()))
	day := strconv.Itoa(datetime.Day())

	return fmt.Sprintf("%s:%s:%s %s/%s/%s", hours, minutes, seconds, year, month, day)
}

func main() {
	fmt.Println(getTime(time.Now()))
}
