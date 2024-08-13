// Activity 5.01 – calculating the working hours of employees

package main

import "fmt"

type Employee struct {
	ID        int
	FirstName string
	LastName  string
}

type Developer struct {
	Individual Employee
	HourlyRate int
	WorkWeek   [7]int
}

type Weekday int

const (
	Minggu Weekday = iota
	Senin
	Selasa
	Rabu
	Kamis
	Jumat
	Sabtu
)

func (d *Developer) LogHours(day Weekday, workHours int) {
	d.WorkWeek[day] = workHours
}

func (d *Developer) HoursWorked() int {
	total := 0
	for _, v := range d.WorkWeek {
		total += v
	}

	return total
}

func main() {
	d := Developer{Individual: Employee{ID: 1, FirstName: "Danang", LastName: "Ari"}, HourlyRate: 10}
	d.LogHours(Senin, 10)
	d.LogHours(Selasa, 8)
	d.LogHours(Rabu, 8)
	d.LogHours(Kamis, 8)
	d.LogHours(Jumat, 8)

	fmt.Println("Hours worked on Senin: ", d.WorkWeek[Senin])
	fmt.Println("Hours worked on Selasa: ", d.WorkWeek[Selasa])
	fmt.Println("Hours worked on Rabu: ", d.WorkWeek[Rabu])
	fmt.Println("Hours worked on Kamis: ", d.WorkWeek[Kamis])
	fmt.Println("Hours worked on Jumat: ", d.WorkWeek[Jumat])

	fmt.Println("Hours worked this week: ", d.HoursWorked())
}
