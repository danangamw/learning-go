// Exercise 1.12 – zero values
// bool : false
// Numbers (integer and float): 0
// String : "" (empty string)
// pointer, functions, interfaces, slices, channels, and maps : nil

// printing subtitution
// %v : Any value, use this if you don't care about the type you're printing.
// %+v : vakues with extra information, such as struct field names.
// %#v : Go syntax, such as %+v with the addition of the name of the type of the variable.
// %T : Print the variable's type.
// %d : Decimal (base 10).
// %s : String.

package main

import (
	"fmt"
	"time"
)

func main() {
	// Print an integer
	var count int
	fmt.Printf("Count    : %#v\n", count)

	// Print a float value
	var discount float64
	fmt.Printf("Discount    : %#v\n", discount)

	// Print a bool value
	var debug bool
	fmt.Printf("Debug    : %#v\n", debug)

	// Print a string value
	var message string
	fmt.Printf("Message    : %#v\n", message)

	// print a collective strings
	var emails []string
	fmt.Printf("Emails    : %#v\n", emails)

	// Print a struct
	var startTime time.Time
	fmt.Printf("Start    : %#v\n", startTime)
}
