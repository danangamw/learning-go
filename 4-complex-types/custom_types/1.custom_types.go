// Exercise 4.16 – Creating a simple custom type
package main

import "fmt"

type id string

func getIDs() (id, id, id) {
	var id1 id
	var id2 id = "1234-5678"
	var id3 id
	id3 = "1234-5678"

	return id1, id2, id3
}

func main() {
	id1, id2, id3 := getIDs()
	fmt.Println("id1 == id2\t\t:", id1 == id2)
	fmt.Println("id2 == id3\t\t:", id2 == id3)
	fmt.Println("id2 == \"1234-5678\"\t:", id2 == "1234-5678")
}
