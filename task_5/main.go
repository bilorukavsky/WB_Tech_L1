/*
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/
package main

import (
	"fmt"
	"time"
)

func sendData(ch chan int) {
	// Отправляем значения в канал каждую секунду.
	for i := 1; ; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
}

func main() {
	dataChannel := make(chan int)

	N := 5

	// Запускаем горутину для отправки данных в канал.
	go sendData(dataChannel)

	// Запускаем таймер для завершения программы через N секунд.
	timer := time.NewTimer(time.Duration(N) * time.Second)

	for {
		select {
		case data := <-dataChannel:
			// Чтение данных из канала и их обработка.
			fmt.Printf("Принято значение из канала: %d\n", data)
		case <-timer.C:
			// Таймер истек, завершаем программу.
			fmt.Println("Программа завершена.")
			return
		}
	}
}
