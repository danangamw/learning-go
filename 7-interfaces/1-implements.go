// Exercise 7.01 – implementing an interface

package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Person struct {
	name      string
	age       int
	isMarried bool
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years old).\nMarried status: %v", p.name, p.age, p.isMarried)
}

func (p Person) Speak() string {
	return "Hi my name is: " + p.name
}

func main() {
	p := Person{name: "Cailyn", age: 44, isMarried: false}
	fmt.Println(p.Speak())
	fmt.Println(p)
}
