// Activity 2.01 – looping over map data using range

package main

import "fmt"

func main() {
	highestWordCount := 0
	var popularWord string

	words := map[string]int{
		"Gonna": 3,
		"You":   3,
		"Give":  2,
		"Never": 1,
		"Up":    4,
	}

	for key, value := range words {
		if highestWordCount < value {
			highestWordCount = value
			popularWord = key
		}
	}

	fmt.Println("Most popular word:", popularWord)
	fmt.Println("With a count of:", highestWordCount)
}
