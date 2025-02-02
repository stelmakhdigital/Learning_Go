package main

import (
	"fmt"
	"time"
)

func check_time_working(start time.Time) {
	// получаем текущее время и с помощью Sub находим разницу между ними
	value := time.Now().Sub(start)
	fmt.Printf("Функция отработала за %v\n", value)
}

func main() {
	defer check_time_working(time.Now())

	// Эмитация бурной деятельности
	time.Sleep(2)
}
