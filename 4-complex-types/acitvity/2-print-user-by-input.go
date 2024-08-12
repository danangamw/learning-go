// Activity 4.02 – Printing a user’s name based on user input
package main

import (
	"fmt"
	"os"
)

func getUsers() map[string]string {
	return map[string]string{
		"305": "Sue",
		"204": "Bob",
		"631": "Jake",
		"073": "Tracy",
	}
}

func getUser(id string) (string, bool) {
	users := getUsers()
	user, exists := users[id]

	return user, exists
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("User ID not passed")
		os.Exit(1)
	}

	userId := os.Args[1]
	user, exists := getUser(userId)
	if !exists {
		fmt.Printf("Passed user id (%v) not found.\nUsers :\n", userId)
		for key, v := range getUsers() {
			fmt.Println("ID :", key, "Name :", v)
		}

		os.Exit(1)
	}

	fmt.Println("Hi,", user)
}
