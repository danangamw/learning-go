// Exercise 5.04 – mapping a CSV index to a column header with
// return values

package main

import (
	"fmt"
	"strings"
)

func csvHdrColl(hdr []string) map[int]string {
	csvIdxToCol := make(map[int]string)
	for i, v := range hdr {
		v = strings.TrimSpace(v)

		switch strings.ToLower(v) {
		case "employee":
			csvIdxToCol[i] = v
		case "hours worked":
			csvIdxToCol[i] = v
		case "hourly rate":
			csvIdxToCol[i] = v
		}
	}

	return csvIdxToCol
}

func main() {
	hdr := []string{"empid", "employee", "address", "hours worked", "hourly rate", "manager"}
	result := csvHdrColl(hdr)

	fmt.Println("Result: ")
	fmt.Println(result)
	fmt.Println()

	hdr2 := []string{"employee", "empid", "hours worked", "address", "manager", "hourly rate"}
	result2 := csvHdrColl(hdr2)
	fmt.Println("Result2: ")
	fmt.Println(result2)
	fmt.Println()
}
