//Напишите программу, которая определяет название языка по его коду. Правила:
//
//en → English
//fr → French
//ru или rus → Russian
//иначе → Unknown


package main

import (
	"fmt"
)

func main() {
	var code string
	fmt.Scan(&code)

	switch code {
	case "en":
		fmt.Println("English")
	case "fr":
		fmt.Println("French")
	case "ru", "rus":
		fmt.Println("Russian")
	default:
		fmt.Println("Unknown")
	}
}

// run `echo "ru" | go run 1_basic_4.go`
// run `echo "rus" | go run 1_basic_4.go`
// run `echo "en" | go run 1_basic_4.go`
// run `echo "fr" | go run 1_basic_4.go`