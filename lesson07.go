// ОСНОВЫ ЯЗЫКА. КОНСТРУКЦИИ. ЦЫКЛЫ В GO

// В Go есть только одна циклическая конструкция - цикл for.
//Базовый цикл for имеет три компонента разделенные точкой с запятой:

// 1) блок инициализации: выполняется перед первой итерацией
// 2) условный блок: выполняется перед каждой итерацией
// 4) завершающий блок: выполняется в конце каждой итерацииУсловные конструкции проверяют истинность некоторого условия.

// ПРИМЕР for счетчик; условие; изменение счетчика {}

package main

import "fmt"

func main() {
	baseCicleFor()
	baseCicleWhile()
}

func baseCicleFor() {
	fmt.Println("ТЕОРИЯ: ЦИКЛ С ОПРЕДЕЛЕННЫМ КОЛИЧЕСВОМ ПОВТОРЕНИЙ")
	// Объявляем переменную
	sum := 0
	// цикл будет повторен 10 раз (пока счетчик не достигнет значения 9)
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("В результате 10 повторений мы получили сумму =", sum) // 45

}

func baseCicleWhile() {
	fmt.Println("ТЕОРИЯ: ЦИКЛ С УСЛОВИЕМ")
	// Объявляем переменную
	sum := 1
	// цикл будет повторяться до тех пор, пока переменная sum будет меньше 1000.
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("В результате условного цикла мы получили сумму =", sum) // 1024
}

//func endlesCicle() {
//	fmt.Println("ТЕОРИЯ: БЕСКОНЕЧНЫЙ ЦИКЛ for")
//	for {
//		fmt.Println("БЕСКОНЕЧНЫЙ ЦИКЛ")
//	}
//
//}
