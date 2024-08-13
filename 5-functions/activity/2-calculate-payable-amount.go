// Activity 5.02 – calculating the payable amount for employees
// based on working hours

package main

import "fmt"

type Employeee struct {
	ID        int
	FirstName string
	LastName  string
}

type Programmer struct {
	Individual Employeee
	HourlyRate int
	WorkWeek   [7]int
}

type WeekDay int

const (
	Sunday WeekDay = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (p *Programmer) logHours(day WeekDay, hours int) {
	p.WorkWeek[day] = hours
}

func (p *Programmer) HoursWorked() int {
	total := 0
	for _, v := range p.WorkWeek {
		total += v
	}

	return total
}

func (p *Programmer) PayDay() (int, bool) {
	if p.HoursWorked() > 40 {
		hourseOver := p.HoursWorked() - 40
		overTime := hourseOver * 2 * p.HourlyRate
		regularPay := 40 * p.HourlyRate
		return regularPay + overTime, true
	}

	return p.HoursWorked() * p.HourlyRate, false
}

func nonLoggedHours() func(int) int {
	total := 0
	return func(i int) int {
		total += i
		return total
	}
}

func (p *Programmer) PayDetails() {
	for i, v := range p.WorkWeek {
		switch i {
		case 0:
			fmt.Println("Sunday hours: ", v)
		case 1:
			fmt.Println("Monday hours: ", v)
		case 2:
			fmt.Println("Tuesday hours: ", v)
		case 3:
			fmt.Println("Wednesday hours: ", v)
		case 4:
			fmt.Println("Thursday hours: ", v)
		case 5:
			fmt.Println("Friday hours: ", v)
		case 6:
			fmt.Println("Saturday hours: ", v)
		}
	}

	fmt.Printf("\nHours worked this week: %d\n", p.HoursWorked())
	pay, overtime := p.PayDay()
	fmt.Println("Pay for the weekL $", pay)
	fmt.Println("Is this overtime pay: ", overtime)
	fmt.Println()
}

func main() {
	p := Programmer{Individual: Employeee{ID: 1, FirstName: "Danang", LastName: "Ari"}, HourlyRate: 10}
	x := nonLoggedHours()

	fmt.Println("Tracking hours worked thus far today: ", x(2))
	fmt.Println("Tracking hours worked thus far today: ", x(3))
	fmt.Println("Tracking hours worked thus far today: ", x(5))
	fmt.Println()

	p.logHours(Monday, 10)
	p.logHours(Tuesday, 10)
	p.logHours(Wednesday, 15)
	p.logHours(Thursday, 14)
	p.logHours(Friday, 13)
	p.logHours(Saturday, 10)

	p.PayDetails()
}
