/*
Реализовать собственную функцию sleep.
*/
package main

import (
	"fmt"
	"time"
)

func Sleep(seconds int) {
	select {
	case <-time.After(time.Duration(seconds) * time.Second): // Преобразует количество секунд в интервал времени, выраженный в секундах, создаем канал time
		return
	}
}

func main() {
	fmt.Println("Начало программы")
	Sleep(3)
	fmt.Println("Прошло 3 секунды")
}
