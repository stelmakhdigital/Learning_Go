package main

import "fmt"

// пример вызов defer в обратном порядке
func main() {
	fmt.Println("Hello")

	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Defer")
}
