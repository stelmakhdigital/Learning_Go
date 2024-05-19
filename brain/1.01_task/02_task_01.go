// Создайте новую директорию с файлом main.go. Напишите код, в котором:
//объявите две переменные, первая - строка со значением 104, вторая - целое число со значением 35;
//приведите строку к целому числу, а целое число - к строке;

package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var valueString string = "104"
	var valueUnsignedNumber uint8 = 35
	fmt.Printf("Значение: %s Тип %T\n", valueString, valueString) // %T - вывод типа без использования reflect
	// плейсхолдер на %v для использования значений разных типов
	fmt.Printf("Значение: %v Тип %s\n", valueUnsignedNumber, reflect.TypeOf(valueUnsignedNumber))

	// Преобразование строки в число
	newFormatValueString, err := strconv.ParseUint(valueString, 10, 8)  // 10 - десятичная система исчислений и 8 - размерность в uint8
	if err != nil {
		fmt.Println("Ошибка преобразования строки valueString:", err)
		return
	}
	valueStringUint8 := uint8(newFormatValueString)

	fmt.Printf("Преобразованное значение из строки в число: %d Тип %T\n", valueStringUint8, valueStringUint8)

	// Преобразование числа в строку
	// 1 способ
	valueUnsignedNumberStringFirst := strconv.Itoa(int(valueUnsignedNumber))

	// 2 способ
	valueUnsignedNumberStringSecond := fmt.Sprintf("%d", valueUnsignedNumber)

	// плейсхолдер на %d для использования значений числовых типов
	fmt.Printf("Преобразованное значение число в строку (1 способ): %s Тип %T\n", valueUnsignedNumberStringFirst, valueUnsignedNumberStringFirst)
	fmt.Printf("Преобразованное значение число в строку (2 способ): %s Тип %T\n", valueUnsignedNumberStringSecond, valueUnsignedNumberStringSecond)
}
