package main

import (
	"fmt"
)

func main() {
	// инициируем переменные
	const mile = 1.609 // километры
	// переводим 120.4 метров/сек в км/час
	var EurooeanVelocity = 120.4 * 3.6
	// переводим 130 метров/сек в мили/час
	var AmericanVelocity = 130 * 2.236936

	fmt.Println(120.4, "метров/сек переводим в Евростандарт:", EurooeanVelocity, "км/ч")
	fmt.Println(130, "метров/сек переводим в Американский стандарт:", AmericanVelocity, "миль/ч")
}
