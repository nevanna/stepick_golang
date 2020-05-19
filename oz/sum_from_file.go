package main

// Импортирум модули
import (
	"fmt"
	"io/ioutil"
)

func main() {
	// Зачитываем содержимое файла
	data, err := ioutil.ReadFile("input.txt")
	// Если во время считывания файла произошла ошибка
	// выводим ее
	if err != nil {
		fmt.Println(err)
	}

	// Если чтение данных прошло успено
	// выводим их в консоль
	var splited [][]string = Split(string(data), ' ')
	fmt.Print(string(data))
}
