// Activity 7.01 – calculating pay and performance review

package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

type Developer struct {
	Individual        Employee
	HourlyRate        float64
	HoursWorkedInYear float64
	Review            map[string]interface{}
}

type Manager struct {
	Individual     Employee
	Salary         float64
	CommissionRate float64
}

type Payer interface {
	Pay() (string, float64)
}

func (d Developer) Pay() (string, float64) {
	name := d.FullName()
	totalPayment := d.HourlyRate * d.HoursWorkedInYear

	return name, float64(totalPayment)
}

func (d Developer) FullName() string {
	return d.Individual.FirstName + " " + d.Individual.LastName
}

func (d Developer) ReviewRating() error {
	total := 0
	for _, v := range d.Review {
		rating, err := OverallReview(v)
		if err != nil {
			return err
		}

		total += rating
	}

	averageRating := float64(total) / float64(len(d.Review))
	fmt.Printf("%s got a review rating of %2.f\n", d.FullName(), averageRating)
	return nil
}

func OverallReview(i interface{}) (int, error) {
	switch v := i.(type) {
	case int:
		return v, nil
	case string:
		rating, err := convertReviewToInt(v)
		if err != nil {
			return 0, err
		}

		return rating, nil
	default:
		return 0, errors.New("unknown type")
	}
}

func convertReviewToInt(str string) (int, error) {
	switch strings.ToLower(str) {
	case "excellent":
		return 5, nil
	case "good":
		return 4, nil
	case "fair":
		return 3, nil
	case "poor":
		return 2, nil
	case "unsatisfactory":
		return 1, nil
	default:
		return 0, errors.New("Invalid rating: " + str)
	}
}

func (m Manager) Pay() (string, float64) {
	name := m.Individual.FirstName + " " + m.Individual.LastName
	totalPayment := m.Salary + (m.Salary * m.CommissionRate)

	return name, float64(totalPayment)
}

func payDetails(p Payer) {
	fullname, yearPay := p.Pay()
	fmt.Printf("%s got paid %2.f for the year\n", fullname, yearPay)
}

func main() {
	employeeReview := make(map[string]interface{})
	employeeReview["WorkQuality"] = 5
	employeeReview["TeamWork"] = 2
	employeeReview["Communication"] = "Poor"
	employeeReview["Problem-solving"] = 4
	employeeReview["Dependability"] = "Good"

	d := Developer{Individual: Employee{Id: 1, FirstName: "Danang", LastName: "Ari"}, HourlyRate: 35, HoursWorkedInYear: 2400, Review: employeeReview}
	m := Manager{Individual: Employee{Id: 2, FirstName: "Rani", LastName: "Putriana"}, Salary: 150000, CommissionRate: 0.07}

	err := d.ReviewRating()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	payDetails(d)
	payDetails(m)
}
