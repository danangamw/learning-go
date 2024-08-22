//  Exercise 11.03 – Printing the Go representation of a variable

package main

import "fmt"

type Person struct {
	lname  string
	age    int
	salary float64
}

func main() {
	fname := "Jon"
	grades := []int{100, 87, 47}
	states := map[string]string{"KY": "Kentucky", "WV": "West Virginia", "VA": "Virginia"}
	p := Person{lname: "Snow", age: 210, salary: 25000.00}

	fmt.Printf("fname value %#v\n", fname)
	fmt.Printf("grades value %#v\n", grades)
	fmt.Printf("states value %#v\n", states)
	fmt.Printf("p value %#v\n", p)
}
