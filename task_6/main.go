/*
Реализовать все возможные способы остановки выполнения горутины.
*/
package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Горутина начала работу")
	time.Sleep(2 * time.Second)
	fmt.Println("Горутина завершила работу")
	done <- true // Отправляем сигнал, что горутина завершила выполнение
}

func main() {
	done := make(chan bool)
	go worker(done)

	// Ждем, пока горутина завершится
	<-done
}
