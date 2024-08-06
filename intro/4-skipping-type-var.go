package main

// Exercise 1.04 – skipping the type or value when declaring variables

import (
	"fmt"
	"time"
)

var (
	Debug       = false
	logLevel    = "info"
	startUpTime = time.Now()
)

func main() {
	fmt.Println(Debug, logLevel, startUpTime)
}
