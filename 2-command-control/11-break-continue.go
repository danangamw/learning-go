// Exercise 2.11 – using break and continue to control loops

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for {
		r := rand.Intn(8)

		if r%3 == 0 {
			fmt.Println("Skip")
			continue
		} else if r%2 == 0 {
			fmt.Println("Stop")
			break
		}

		fmt.Println(r)
	}
}
