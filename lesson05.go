// ОСНОВЫ ЯЗЫКА. МАПЫ В GO
//
// Map'a представляет собой неупорядоченный набор элементов одного типа,
// которые попарно связаны с другими элементами по принципу [ключ : значение]. Являются аналогом хэш-таблиц и словарей.
// Позволяют быстро найти найти, получить, изменить и удалить значение по ключу.
// В качестве ключей могут быть значения типов данных, которые можно сравнивать (string, float, int, bool)
// Значения в мапах могут иметь любой тип данных включая слайсы, мапы и структуры

package main

import (
	"fmt"
	"sort"
)

func main() {
	mapDefinition()
	takeValueFromMap()
	addToMap()
	deleteFromMap()
	mapIteration()

	homeWork09()
}

func mapDefinition() {
	fmt.Println("ТЕОРИЯ: СОЗДАЕМ Map'ы")

	// Инициируем пустую map'у по схеме map[тип данных ключей]тип данных значений
	var book map[string]int // не делайте так В нее нельзя ничего добавить

	// Инициируем map'у и присваиваем ей значения с ключом
	var author map[string]string = map[string]string{
		"name":     "Aleksandr",
		"lastname": "Pushkin"}

	// Есть краткий способ определить map'у с присваиванием ей значения с ключом
	reader := map[string]string{
		"name":     "Ivan",
		"lastname": "Turgenev",
		"city":     "Moscow",
	}
	fmt.Println("Выводим созданные мапы, на экран:")
	fmt.Println("Пустая мапа book =", book)
	fmt.Println("Мапа author =", author)
	fmt.Println("Мапа reader =", reader)

	// Инициируем map'у и присваиваем ей мапу в качестве значения
	sales := map[int]map[string]int{
		2022: {
			"automobile": 430,
			"bike":       340,
		},
		2023: {
			"automobile": 450,
			"bike":       320,
		}}
	fmt.Println("Выводим мапу на экран:", sales)
}

func takeValueFromMap() {
	fmt.Println("ТЕОРИЯ: ИЗВЛЕКАЕМ ЗНАЧЕНИЕ ИЗ МАПЫ ПО КЛЮЧУ")
	// Инициируем map'у и присваиваем ей ключи и значения
	m := map[string]int{
		"раз": 1,
		"два": 2,
		"три": 3}
	fmt.Println("Выводим исходную мапу:", m)
	// Извлечь значение из мапы можно по ключу
	first := m["раз"]
	fmt.Println("Выводим извлеченное по ключу значение, на экран:")
	fmt.Println("Значение переченныой под ключом 'раз' =", first)
}

func addToMap() {
	fmt.Println("ТЕОРИЯ: ДОБАВЛЯЕМ В Map'у новые значения")
	// Инициируем map'у и присваиваем ей ключи и значения
	answers := map[string]int{
		"ответ01": 3,
		"ответ02": 3,
		"ответ03": 3}
	fmt.Println("Выводим исходную мапу:", answers)
	// Добавляем в map'у новые ключ и значения
	answers["ответ04"] = 5
	fmt.Println("Выводим мапу с добавлением:", answers)
}

func deleteFromMap() {
	fmt.Println("ТЕОРИЯ: УДАЛЯЕМ ЗНАЧЕНИЕ ИЗ МАПЫ ПО КЛЮЧУ")
	// Инициируем map'у и присваиваем ей ключи и значения
	users := map[string]bool{
		"user01": true,
		"user02": true,
		"user03": false}
	fmt.Println("Выводим исходную мапу:", users)
	// Удаление из мапы реалуизуется при помощи функции delete(мапа, ключ из мапы)
	delete(users, "user03")
	fmt.Println("Выводим мапу после удаления user03:", users)
}

func mapIteration() {
	fmt.Println("ТЕОРИЯ: ОБХОД ЭЛЕМЕТНОВ МАПЫ В ЦИКЛЕ ")
	// Инициируем map'у и присваиваем ей ключи и значения
	users := map[string]bool{
		"user96": true,
		"user42": true,
		"user73": false}
	fmt.Println("Выводим исходную мапу:", users)
	// Инициируем слайс со строками длиной равной количеству элементов в мапе
	keys := make([]string, 0, len(users))
	// В цикле наполняем слайс ключами из мапа
	for k := range users {
		keys = append(keys, k)
	}
	// Сортируем слайс ключами из мапа
	sort.Strings(keys) // [user42 user73 user96]
	// В цикле обращаемся к каждому значению мапа по ключу
	for _, v := range keys {
		fmt.Println(users[v])
	}

	fmt.Println(len(keys), cap(keys))

}

func homeWork09() {
	// Задание
	// 1) Создайте мап в которой необходимо хранить информацию
	// о выданных читателю печатных изданиях: книгах, переодических изданиях
	// 2) Тип ключей мапа - строка, тип значений - мапа с
	// [ ключем типа - строка, значением слайс строк ]
	// 3) Добавьте произвольные значения, моделирующие нахождение у читателя книг и периодики
	// Выведите на экран количество читателей с изданиями на руках.
	// Выведите на экран общее количество изданий на руках одного читателя.

	readersBalance := map[string]map[string]int{
		"reader01": {
			"book":     12,
			"periodic": 6,
		},
		"reader02": {
			"book":     8,
			"periodic": 2,
		},
		"reader03": {
			"book":     3,
			"periodic": 7,
		},
	}
	// Инициируем слайс со строками длиной равной количеству элементов в мапе
	readers := make([]string, 0, len(readersBalance))

	for reader := range readersBalance {
		readers = append(readers, reader)
	}
	fmt.Println(readers, len(readers))

	// Сортируем слайс ключами из мапа
	sort.Strings(readers)
	// В цикле обращаемся к каждому значению мапа по ключу
	var bookCounter, periodicCounter int
	for _, v := range readers {
		thisReader := readersBalance[v]
		bookCounter += thisReader["book"]
		periodicCounter += thisReader["periodic"]
	}
	fmt.Println("Читателей всего:", len(readers), readers)
	fmt.Println("На руках читателей всего:", bookCounter+periodicCounter)
	fmt.Println("Из них книг:", bookCounter, "переодических изданий:", periodicCounter)
}
