package main

import "fmt"

func main() {
	fmt.Println(unintuitive())

	fmt.Println(intuitive())
}

// используя именованное возвращаемое значение можно value присвоить новое значение с defer
func unintuitive() (value string) {
	defer func() { value = "На самом деле" }() // () в конце означает вызов функции
	return "Казалось бы"
}

// в этом случае значение возвращаемое значение будет копироваться и defer не повлияет на его изменение 
func intuitive() string {
	value := "Казалось бы"
	defer func() { value = "На самом деле" }()
	return value
}
