package main

import (
	"fmt"

	"github.com/danangamw/printer"
)

func main() {
	msg := printer.PrintNewUUID()
	fmt.Println(msg)
}
