/*
На основе функции PrintAllFiles из предыдущего примера реализуйте функцию
PrintAllFilesWithFilter(path string, filter string), которая будет печатать только путь со строкой filter.
*/

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	PrintAllFilesWithFilter(".", "08")
}

func PrintAllFilesWithFilter(path string, filter string) {
	// получаем список всех элементов в папке (и файлов, и директорий)
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("unable to get list of files", err)
		return
	}
	//  проходим по списку
	for _, f := range files {
		// получаем имя элемента
		// filepath.Join — функция, которая собирает путь к элементу с разделителями
		filename := filepath.Join(path, f.Name())

		if strings.Contains(filename, filter) {
			// печатаем имя элемента
			fmt.Println(filepath.Dir(filename))
		}
		// если элемент — директория, то вызываем для него рекурсивно ту же функцию
		if f.IsDir() {
			PrintAllFilesWithFilter(filename, filter)
		}
	}
}

/*
# Вариант из практикума
func PrintAllFilesWithFilterClosure(path string, filter string) {
    // создаём переменную, содержащую функцию обхода
    // мы создаём её заранее, а не через оператор :=, чтобы замыкание могло сослаться на него
    var walk func(string)
    walk = func(path string) {
        // получаем список всех элементов в папке (и файлов, и директорий)
        files, err := os.ReadDir(path)
        if err != nil {
            fmt.Println("unable to get list of files", err)
            return
        }
        //  проходим по списку
        for _, f := range files {
            // получаем имя элемента
            // filepath.Join — функция, которая собирает путь к элементу с разделителями
            filename := filepath.Join(path, f.Name())
            // печатаем имя элемента, если путь к нему содержит filter, который получим из внешнего контекста
            if strings.Contains(filename, filter) {
                fmt.Println(filename)
            }
            // если элемент — директория, то вызываем для него рекурсивно ту же функцию
            if f.IsDir() {
                walk(filename)
            }
        }
    }
    // теперь вызовем функцию walk
    walk(path)
}

*/
