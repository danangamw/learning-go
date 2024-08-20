package main

// edit for connecting github

import (
	"fmt"

	"github.com/google/uuid"
	"rsc.io/quote"
)

func main() {
	id := uuid.New()
	quote := quote.Go()

	fmt.Println("Generated UUID: ", id)
	fmt.Println("Random Quote: ", quote)
}
