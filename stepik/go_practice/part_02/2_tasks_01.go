//Напишите программу, которая укорачивает строку до указанной длины и добавляет в конце многоточие:
//
//text = Eyjafjallajokull, width = 6 → Eyjafj...
//Если строка не превышает указанной длины, менять ее не следует:
//
//text = hello, width = 6 → hello
//Гарантируется, что в исходной строке text используются только однобайтовые символы без пробелов, а длина width строго больше 0
// Возьмите первые `width` байт строки `text`,
// допишите в конце `...` и сохраните результат
// в переменную `res`
// ...

package main

import (
	"fmt"
)

func main() {
	var text string
	var width int
	fmt.Scanf("%s %d", &text, &width)

	res := ""
	if len(text) <= width {
		res = text
	} else {
		res = text[:width] + "..."
	}
	fmt.Println(res)
}