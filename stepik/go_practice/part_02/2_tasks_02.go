// Напишите программу, которая считает, сколько раз каждая цифра встречается в числе.
//Гарантируется, что на вход подаются только положительные целые числа, не выходящие за диапазон int.
// 12823
// 1:1 2:2 3:1 8:1


package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var input string
	fmt.Scanln(&input)

	// Посчитайте, сколько раз каждая цифра встречается
	// в числе `input`. Запишите результат в карту `counter`,
	// где ключом является цифра числа, а значением -
	// сколько раз она встречается
	counter := make(map[int]int)
	for _, char := range input {
		digit, _ := strconv.Atoi(string(char))
		counter[digit]++
	}

	printCounter(counter)
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// printCounter печатает карту в формате
// key1:val1 key2:val2 ... keyN:valN
func printCounter(counter map[int]int) {
	digits := make([]int, 0)
	for d := range counter {
		digits = append(digits, d)
	}
	sort.Ints(digits)
	for idx, digit := range digits {
		fmt.Printf("%d:%d", digit, counter[digit])
		if idx < len(digits)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
}

// run `echo 12823 | go run 2_tasks_02.go`