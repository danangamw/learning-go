// Exercise 3.05 – Safely looping over a string

package main

import "fmt"

func main() {
	logLevel := "デバッグ"

	for index, runeVal := range logLevel {
		fmt.Println(index, string(runeVal))
	}
}
