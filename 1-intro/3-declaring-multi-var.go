package main

// Exercise 1.03 Declaring multiple variables at once with var

import (
	"fmt"
	"time"
)

var (
	Debug       bool      = false
	logLevel    string    = "info"
	startUpTime time.Time = time.Now()
)

func main() {
	fmt.Println(Debug, logLevel, startUpTime)
}
