// Создайте новую директорию и новый файл main.go. Напишите код, в котором:
// объявите два новых типа AmericanVelocity и EuropeanVelocity
// выполните преобразование скорости 120.4 м/сек в км/ч и присвойте результат переменной с типом EuropeanVelocity;
// выполните преобразование скорости 130 м/с в мили/ч и присвойте результат переменной с типом AmericanVelocity;
// примечание: 1 миля = 1.609 км. Если потребуется, округлите значение до 2 знаков после запятой, для округления обратитесь к пакету math.

package main

import (
	"fmt"
	"math"
)

type EuropeanVelocity float64
type AmericanVelocity float64

func roundValue(value float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))  // Возведение в степень в десятичной системе
	return math.Round(value * ratio) / ratio
}

func main() {
	var europeanValue EuropeanVelocity = 120.4 * 3.6
	var americanValue AmericanVelocity = (130 * 3.6) / 1.609

	fmt.Printf("Value: %.6f\n", europeanValue)
	fmt.Printf("Равно: %.2f км/ч\n", europeanValue) // просто печатаем до двух знаков после запятой

	fmt.Printf("Value: %.6f\n", americanValue)
	fmt.Printf("Равно: %.2f мили/ч\n", americanValue) // просто печатаем до двух знаков после запятой

	// c применением math
	fmt.Println(roundValue(float64(europeanValue), 2))
	fmt.Println(roundValue(float64(americanValue), 2))

}

