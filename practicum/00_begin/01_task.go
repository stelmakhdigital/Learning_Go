// Создайте слайс и заполните его числами от 1 до 100.
// Оставьте в слайсе первые и последние 10 элементов и разверните слайс в обратном порядке.

// Подумайте, можно ли подобным образом развернуть строку?

package main

import "fmt"

func funcSlice() {

	array_size := 100

	slice := make([]int, 0, array_size) // создаем пустой слайс с базовым массивом с (len ==  100)

	for index := 0; index < array_size; index++ {
		slice = append(slice, index+1)
	}

	// оставляем первые и последние 10 элементов
	slice = append(slice[:10], slice[array_size-10:]...)

	array_size = len(slice) // задаем новый размер для слайс

	// переворачиваем слайс
	// slice[:array_size/2] - проходим до середени и переназначаем значения по индексам
	for i := range slice[:array_size/2] {
		slice[i], slice[array_size-i-1] = slice[array_size-i-1], slice[i]
	}

	fmt.Println(slice)
}

// Разернуть строку таким образом не получится, так как строка — неизменяемый тип данных.
// Строку можно преобразовать к слайсу байт ([]byte), но это опасно, если строка содержит Unicode-символы
// Лучше создать слайс рун из строки и развернуть его

func funcString() {
	input := "The quick brown 狐 jumped over the lazy 犬"
	points := 0

	// Создаем слайс рун
	slice_rune := make([]rune, len(input))

	for _, r := range input {
		slice_rune[points] = r // Назначаем руну
		points++
	}
	fmt.Println(slice_rune)
	slice_rune = slice_rune[0:points] // удаляет нуленые руны [... 121 32 29356 0 0 0 0 <- эти]
	fmt.Println(slice_rune)

	// разворачиваем руны
	for i := 0; i < points/2; i++ {
		slice_rune[i], slice_rune[points-1-i] = slice_rune[points-1-i], slice_rune[i]
	}

	//преобразовываем руны в текст
	output := string(slice_rune)

	fmt.Println(output)
}

func main() {
	funcSlice()
	funcString()

}
