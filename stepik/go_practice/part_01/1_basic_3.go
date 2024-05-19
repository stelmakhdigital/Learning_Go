package main

import (
	"fmt"
)

func main() {
	var source string
	var result string
	var resultAlt string
	var times int
	// гарантируется, что значения корректные
	fmt.Scan(&source, &times)

	// возьмите строку `source` и повторите ее `times` раз
	// запишите результат в `result`
	for i := 1; i <= times; i++ {
		result += source
	}

	// Оптимизированный вариант без временного инкремента
	for ; times > 0; times-- {
		resultAlt += source
	}

	// либо сразу вывесли fmt.Println(strings.Repeat(source,times))

	fmt.Println(result)
	fmt.Println(resultAlt)
}

// run `echo "a 5" | go run 1_basic_3.go`