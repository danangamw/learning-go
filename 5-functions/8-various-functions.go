// Exercise 5.08 – creating various functions to calculate salary

package main

import "fmt"

func salary(x, y int, f func(int, int) int) int {
	pay := f(x, y)
	return pay
}

func managerSalary(baseSalary, bonus int) int {
	return baseSalary + bonus
}

func developerSalary(hourlyRate, hoursWorked int) int {
	return hourlyRate * hoursWorked
}

func main() {
	devSalary := salary(50, 2080, developerSalary)
	bossSalary := salary(150000, 25000, managerSalary)
	fmt.Printf("Boss salary: %d\n", bossSalary)
	fmt.Printf("Developer Salary: %d\n", devSalary)
}
