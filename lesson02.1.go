// ОСНОВЫ ЯЗЫКА. ПЕРЕМЕННЫЕ И КОНСТАНТЫ В GO
//
// Переменные и константы служат для хранения информации. Значения переменных можно поменять, а значения констант - нет.
// Переменные и константы должны быть объявлены перед их использованием.
// Golang выведет ошибку при попытке присвоить значение переменной, которая не была объявлена.
// Переменные объявляются при помощи var, константы - при помощи const.

package main

import (
	"fmt" // библиотека для вывода в консоль
)

func main() {
	fullVarDefinition()    // Способы определения переменных при помощи var
	shortVarDefinition()   // Способ краткого определения переменных при помощи :=
	constDefinition()      // Способы определения констант при помощи const
	groupConstDefinition() // Способы определения группы констант
	constGenerator_iota()  // Способы определения группы констант генератором iota
}

func fullVarDefinition() {
	fmt.Println("ТЕОРИЯ: ОПРЕДЕЛЯЕМ ПЕРЕМЕННЫЕ ПРИ ПОМОЩИ VAR и УКАЗАНИЯ ТИПА ДАННЫХ")
	// Переменную можно объявить не устанавливая значение
	// Но тогда нужно обязательно указать тип данных
	var var1 int
	var var2 float32
	var var3 string
	var var4 bool

	// Можно записать это более кратко
	var (
		var5 int
		var6 float32
		var7 string
		var8 bool
	)

	fmt.Println("Тогда ей присвоется значение по умолчанию:")
	fmt.Println("Для int:", var1, "для float32:", var2, "для string:", var3, "(пустая строка)", "для bool:", var4)

	// После определения переменной можно присвоить значения
	var1 = 123
	var2 = 12.3
	var3 = "Строка"
	var4 = true

	// Можно записать это в одну строку
	var5, var6, var7, var8 = 123, 12.3, "Строка", true

	fmt.Println("Переменным присвоены новые значения:")
	fmt.Println("Для int:", var5, "для float32:", var6, "для string:", var7, "(пустая строка)", "для bool:", var8)
}

func shortVarDefinition() {
	fmt.Println("ТЕОРИЯ: ОПРЕДЕЛЯЕМ ПЕРЕМЕННЫЕ КРАТКО")
	// Переменную можно объявить более кратко если сразу присвоить значение. Это делается при помощи оператоа := (Морж)
	var1 := 123
	var2 := 12.3
	var3 := "Строка"
	var4 := true
	// Можно записать это в одну строку
	var5, var6, var7, var8 := 123, 12.3, "Строка", true
	fmt.Println("var1 = var5 = 123", var1 == var5)
	fmt.Println("var2 = var6 = 12.3", var2 == var6)
	fmt.Println("var3 = var7 = Строка", var3 == var7)
	fmt.Println("var4 = var8 = true", var4 == var8)
}

func constDefinition() {
	fmt.Println("ТЕОРИЯ: ОПРЕДЕЛЕНИЕ КОНСТАНТ")
	// При объявлении константам должны быть присвоены значения.
	// Тип данных при этом указывать НЕ НУЖНО.
	const const01 = 123
	const const02 = 12.3
	const const03 = "Строка"
	const const04 = true

	// Можно записать это более кратко
	const (
		const05 = 123
		const06 = 12.3
		const07 = "Строка"
		const08 = true
	)
	fmt.Println("const01 = const05 = 123", const01 == const05)
	fmt.Println("const02 = const06 = 12.3", const02 == const06)
	fmt.Println("const03 = const07 = Строка", const03 == const07)
	fmt.Println("const04 = const08 = true", const04 == const08)
}

func groupConstDefinition() {
	fmt.Println("ТЕОРИЯ: ОПРЕДЕЛЕНИЕ ПОСЛЕДОВАТЕЛЬНОСТЕЙ КОНСТАНТ")
	// При объявлении последовательности констант можно пропускать некоторые значения.
	// В таком случае константе присвоится значение предыдущей константы.
	const (
		random01 = 123
		random02
		random03 = "Строка"
		random04
	)
	fmt.Println("random01 = random02 = 123 ?", random01 == random02)
	fmt.Println("random03 = random04 = Строка ?", random03 == random04)
}

func constGenerator_iota() {
	fmt.Println("ТЕОРИЯ: ГЕНЕРАТОР ПОСЛЕДОВАТЕЛЬНОСТЕЙ КОНСТАНТ iota")
	// При групповом объявлении связанных констант можно использовать генератор констант iota (автоинкремент)
	// Значение iota увеличивается на елиницу начиная с 0.
	// Для пропуска ненужных значений можно использовать пустой идентификатор _.
	fmt.Println("Создаем список констант iota, пропуская значения пустым идентификатором _")
	const (
		_       = 10 * iota
		speed10 = 10 * iota
		_       = 10 * iota
		speed30 = 10 * iota
		_       = 10 * iota
		speed50 = 10 * iota
	)
	fmt.Println("Получаются следующие константы:")
	fmt.Println(speed10, speed30, speed50)
}
