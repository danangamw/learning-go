package main

import (
	"fmt"
	"strings"
)

var (
	ErrInvalidLastName      = "invalid last name"
	ErrInvalidRoutingNumber = "invalid routing number"
)

type DirectDeposit struct {
	FirstName     string
	LastName      string
	BankName      string
	RoutingNumber int
	AccountNumber int
}

func (d *DirectDeposit) validateRoutingNumber() {
	defer func() {
		if r := recover(); r != nil {
			if r == ErrInvalidRoutingNumber {
				fmt.Println(r)
			}
		}
	}()

	if d.RoutingNumber < 100 {
		panic(ErrInvalidRoutingNumber)
	}

}

func (d *DirectDeposit) validateLastName() error {
	defer func() {
		if r := recover(); r != nil {
			if r == ErrInvalidLastName {
				fmt.Println(r)
			}
		}
	}()

	if len(strings.TrimSpace(d.LastName)) <= 0 {
		panic(ErrInvalidLastName)
	}

	return nil
}

func (d *DirectDeposit) report() {
	fmt.Println("Last Name: ", d.LastName)
	fmt.Println("First Name: ", d.FirstName)
	fmt.Println("Bank Name: ", d.BankName)
	fmt.Println("Routing Number: ", d.RoutingNumber)
	fmt.Println("Account Number: ", d.AccountNumber)
}

func main() {
	deposit := DirectDeposit{FirstName: "Abe", LastName: "", BankName: "XYZ Inc", RoutingNumber: 17, AccountNumber: 1809}
	deposit.validateRoutingNumber()
	deposit.validateLastName()
	fmt.Println("**********************")
	deposit.report()
}
