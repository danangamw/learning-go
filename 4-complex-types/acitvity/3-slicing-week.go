// Activity 4.03 – Slicing the week

package main

import "fmt"

func getWeekDays() []string {
	week := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	week = append(week[6:], week[:6]...)
	return week
}

func main() {
	fmt.Println(getWeekDays())
}
