package main

// Exercise 1.06 – declaring multiple variables from a function

import (
	"fmt"
	"time"
)

func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}

func main() {
	Debug, LogLevel, startUpTime := getConfig()

	fmt.Println(Debug, LogLevel, startUpTime)
}
