// Exercise 4.09 – Appending multiple items to a slice

package main

import (
	"fmt"
	"os"
)

func getPassedArgs1() []string {
	var args []string

	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}

	return args
}

func getLocales(extraLocales []string) []string {
	var locales []string
	locales = append(locales, "en_US", "fr_FR")
	locales = append(locales, extraLocales...)

	return locales
}

func main() {
	locales := getLocales(getPassedArgs1())
	fmt.Println("Locales to use:", locales)
}
