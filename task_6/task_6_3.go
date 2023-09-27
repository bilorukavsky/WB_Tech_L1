/*
Реализовать все возможные способы остановки выполнения горутины.
*/
package main

import (
	"fmt"
	"runtime"
	"time"
)

func worker3() {
	defer fmt.Println("Горутина завершила работу")
	fmt.Println("Горутина начала работу")
	runtime.Goexit() // Остановка выполнения горутины
}

func task_6_3() {
	go worker3()

	// Добавьте какую-то задержку, чтобы дать горутине время начать выполнение
	time.Sleep(100 * time.Millisecond)
}
