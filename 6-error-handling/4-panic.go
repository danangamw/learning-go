// Exercise 6.04 – Crashing the program on errors using a panic

package main

import (
	"errors"
	"fmt"
)

var (
	errHourlyRate  = errors.New("invalid hours rate")
	errHoursWorked = errors.New("invalid hours worked per week")
)

func PayDay(hoursWorked, hourlyRate int) int {
	report := func() {
		fmt.Printf("HoursWorked: %d\nHourlyRate: %d\n", hoursWorked, hourlyRate)
	}
	defer report()

	if hourlyRate < 10 || hourlyRate > 75 {
		panic(errHourlyRate)
	}
	if hoursWorked < 0 || hoursWorked > 80 {
		panic(errHoursWorked)
	}

	if hoursWorked > 40 {
		hoursOver := hoursWorked - 40
		overTime := hoursOver * 2
		regularPay := hoursWorked * hourlyRate
		return regularPay + overTime
	}
	return hoursWorked * hourlyRate

}

func main() {
	pay := PayDay(81, 50)
	fmt.Println(pay)
}
