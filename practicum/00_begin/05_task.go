package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Name        string     `json:"Имя"`
	Email       string     `json:"Почта"`
	DateOfBirth *time.Time `json:"-"` // - означает, что это поле не будет сериализовано

}

func NewPerson(name, email string, dob ...int) (Person, error) {

	if name == "" || email == "" {
		return Person{}, fmt.Errorf("Name and email must be non-empty")
	}

	var dobPointer *time.Time // чтобы передать указатель в return конструктора

	if len(dob) == 3 {
		year, month, day := dob[0], dob[1], dob[2]
		dobValue := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
		dobPointer = &dobValue
	} else {
		dobValue := time.Now()
		dobPointer = &dobValue
	}
	return Person{
		Name:        name,
		Email:       email,
		DateOfBirth: dobPointer,
	}, nil
}

func main() {

	result, err := NewPerson("Alex", "test@test.ru")
	// result, err := NewPerson("John", "john@example.com", 1990, 5, 14)

	if err == nil {
		fmt.Printf("%+v\n", result)
		fmt.Printf("%v\n", result)
	} else {
		// сбор лога о провальной валидации
		fmt.Println(err)
	}

	jsonResult, err := json.Marshal(result)

	if err == nil {
		fmt.Println(string(jsonResult))
		// {"Имя":"Alex","Почта":"test@test.ru"}
	}

}
