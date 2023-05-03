// ОСНОВЫ ЯЗЫКА. МАCCИВЫ В GO
//
// Массив – это коллекция фиксированного размера.
// Акцент здесь ставится именно на фиксированный размер,
// поскольку, как только вы зададите длину массива, позже вы уже не сможете ее изменить.
// Массив(array) в go является составным типом данных, который может содержать элементы одного типа.
// Доступ к элементам осуществляется по индексу. Попытка обратиться к несущестующему интексу - приведет к ошибке.

package main

import "fmt"

func main() {
	arrayDefinition()
	arrayElements()
}

func arrayDefinition() {
	fmt.Println("ТЕОРИЯ: СОЗДАЕМ МАССИВЫ(ARRAY)")
	// При инициации пустого массива обязательно указываем длину
	var randomArray01 [5]int // [0 0 0 0 0]
	fmt.Println(randomArray01)

	// Инициацию массива можно делать вместе с присвоением значения
	var randomArray02 = [3]string{"first", "second", "third"}
	fmt.Println(randomArray02)

	// Есть краткий способ создания массива при помощи := (моржа)
	days := [7]string{"ПН", "ВТ", "СР", "ЧТ", "ПТ", "СБ", "ВС"}
	fmt.Println(days)
}

func arrayElements() {
	fmt.Println("ТЕОРИЯ: ДОСТУП К ЭЛЕМЕНТАМ МАССИВА(ARRAY)")
	digits := [5]int{1, 1, 1, 1, 1}
	// Чтобы установить значение элемента в массиве, нужно обратиться к индексу элемента, например: массив[индекс]
	digits[2] = 5 // [1 1 5 1 1]
	fmt.Println(digits)
}