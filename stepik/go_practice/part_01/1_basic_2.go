// Напишите программу, которая рассчитает евклидово расстояние между точками на плоскости (x1, y1) и (x2, y2):

package main

import (
	"fmt"
	"math"
)

func main() {
	// объявите переменные x1, y1, x2, y2 типа float64
	var x1, y1, x2, y2 float64

	// считываем числа из os.Stdin
	// гарантируется, что значения корректные
	// не меняйте этот блок
	fmt.Scan(&x1, &y1, &x2, &y2)

	// рассчитайте result по формуле эвклидова расстояния
	// используйте math.Pow(x, 2) для возведения в квардрат
	// используйте math.Sqrt(x) для извлечения корня
	result := math.Sqrt(math.Pow(x1 - x2, 2) + math.Pow(y1-y2, 2))

	// Можно не использовать Sqrt а также возвести в степень 0.5
	resultAlt := math.Pow(math.Pow(x1 - x2, 2) + math.Pow(y1-y2, 2), 0.5)

	// выводим результат в os.Stdout
	fmt.Println(result)
	fmt.Println(resultAlt)
}

// run `echo "1 1 4 5" | go run 1_basic_2.go`