// Activity 4.05 – Creating a locale checker

package main

import (
	"fmt"
	"os"
	"strings"
)

type locale struct {
	language string
	country  string
}

func getLocales() map[locale]struct{} {
	supportedLocales := make(map[locale]struct{})
	supportedLocales[locale{language: "en", country: "US"}] = struct{}{}
	supportedLocales[locale{language: "id", country: "ID"}] = struct{}{}
	supportedLocales[locale{language: "en", country: "CN"}] = struct{}{}
	supportedLocales[locale{language: "fr", country: "FR"}] = struct{}{}
	supportedLocales[locale{language: "in", country: "IN"}] = struct{}{}

	return supportedLocales
}

func localeExists(l locale) bool {
	_, exists := getLocales()[l]
	return exists
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No locale passed")
		os.Exit(1)
	}

	localeParts := strings.Split(os.Args[1], "_")
	if len(localeParts) != 2 {
		fmt.Printf("Invalid locale passed: %v\n", os.Args[1])
		os.Exit(1)
	}

	passedLocale := locale{
		language: localeParts[0],
		country:  localeParts[1],
	}
	exists := localeExists(passedLocale)
	if !exists {
		fmt.Printf("Locale not supported: %v\n", os.Args[1])
		os.Exit(1)
	}

	fmt.Println("Locale passed is supported")

}
