package main

import (
	"fmt"     // библиотека для вывода в консоль
	"reflect" // библиотека для определения типа
	"strconv" // библиотека для конвертации типов
)

func main() {
	// инициируем переменные со строкой и числом
	var strdig, intdig = "104", 35
	// приводим строку к числу
	res1, err1 := strconv.Atoi(strdig)
	if err1 != nil {
		// обработка ошибки
		panic(err1)
	}
	// приводим число к строке
	res2 := strconv.Itoa(intdig)
	fmt.Println("Даны строка:", strdig, "и число:", intdig)
	fmt.Println(strdig, reflect.TypeOf(strdig), "становится числом", res1, reflect.TypeOf(res1))
	fmt.Println(intdig, reflect.TypeOf(intdig), "становится строкой", res2, reflect.TypeOf(res2))
}
