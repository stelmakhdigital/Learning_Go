package main

import (
	"fmt"
	"reflect"
)

const someoneString = "Будьте вежливы и соблюдайте наши принципы сообщества."

func main() {
	arrayString := [3]string{"a", "b", "c"}
	arrayDynamic := [...]string{"a", "b"} // создаст массив под кол-во элементов
	arrayEmpty := [5]string{}

	sliceString := []string{"a", "b", "c"} // slice определяется типом элементов, но кол-вом как array


	fmt.Printf("arrayString: %s, type: %v\n", arrayString, reflect.TypeOf(arrayString))
	fmt.Printf("arrayDynamic: %s, type: %v\n", arrayDynamic, reflect.TypeOf(arrayDynamic))
	fmt.Printf("arrayEmpty: %s, type: %v\n", arrayEmpty, reflect.TypeOf(arrayEmpty))

	// передаем массив по ссылке и присваиваем значение по индексу
	arrayString2 := &arrayString
	arrayString2[2] = "2"
	fmt.Printf("arrayString2: %v, type: %v\n", arrayString2, reflect.TypeOf(arrayString2))

	fmt.Printf("sliceString: %v, type: %v\n", sliceString, reflect.TypeOf(sliceString))
	// добавим элемент в срез
	newSliceString := append(sliceString, "SECOND") // обязательно назначать переменную в append - возвращает копию
	fmt.Printf("newSliceString: %v, type: %v\n", newSliceString, reflect.TypeOf(newSliceString))
	// ... так как сам первоначальный срез не меняется
	fmt.Printf("sliceString: %v, type: %v\n", sliceString, reflect.TypeOf(sliceString))

	// создание пустого среза
	emptySliceString := make([]string, len(sliceString))
	fmt.Printf("emptySliceString: %v, type: %v\n", emptySliceString, reflect.TypeOf(emptySliceString))

	// копирование среза в первый из второго
	copy(emptySliceString, newSliceString)
	fmt.Printf("Copy (newSliceString) %v (%v) in (emptySliceString) %v (%v)\n", newSliceString, len(newSliceString), emptySliceString, len(emptySliceString))
	fmt.Printf("emptySliceString: %v, type: %v\n", emptySliceString, reflect.TypeOf(emptySliceString))
	fmt.Println("SECOND было отброшено, тк элементов в срезе emptySliceString всего 3")

	// ну и сами срезы
	sliceFromSlice := sliceString[:1] // значение(включительно) : значение до (не включительно)
	fmt.Printf("sliceFromSlice: %v, type: %v\n", sliceFromSlice, reflect.TypeOf(sliceFromSlice))

	newSliceFromSlice := sliceString[1:] // значение(включительно) : до конца
	fmt.Printf("newSliceFromSlice: %v, type: %v\n", newSliceFromSlice, reflect.TypeOf(newSliceFromSlice))

	massiveSlice := []string{"a", "b", "c", "d", "e", "f"} // [2:5] -> [c d e]
	fmt.Printf("newSliceFromSlice: %v, type: %v\n", massiveSlice[2:5], reflect.TypeOf(massiveSlice))

	str := "Строку можно преобразовать в срез байт и обратно"

	strBytes := []byte(str) // uint8
	fmt.Printf("strBytes: %v, type: %v, string: %s\n", strBytes, reflect.TypeOf(strBytes), string(strBytes))

	// срез unicode-символов - руны
	strRunes := []rune(str)  //int32
	fmt.Printf("strRunes: %v, type: %v, string: %s\n", strRunes, reflect.TypeOf(strRunes), string(strRunes))

	// в Go строка реализована как массив байт, а не рун
	newStr := "в Go строка реализована как массив байт, а не рун"
	bytes := []byte(newStr)
	fmt.Println(bytes[3]) // 71

	runes := []rune(newStr)
	fmt.Println(runes[3]) // 111

	fmt.Println(newStr[3]) // обращение к строке по индексу вернет 71 - байт в индексе
	fmt.Println(string(newStr[3])) // а вот так вернет 'G'
}