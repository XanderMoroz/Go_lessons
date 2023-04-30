package main

import (
	"fmt"
	"time"
)

func main() {
	time_now := time.Now()

	year, month, day := time_now.Year(), time_now.Month(), time_now.Day()
	hour, minute := time_now.Hour(), time_now.Minute()

	fmt.Println("Hello, 世界")
	fmt.Println(day, ".", month, ".", year, " ", hour, ":", minute)
}
