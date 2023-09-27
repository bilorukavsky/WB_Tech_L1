/*
Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров,
которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/
package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func worker(id int, dataChannel <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range dataChannel {
		// Воркеры читают данные из канала и выводят их в stdout.
		fmt.Printf("Воркер %d: %s\n", id, data)
	}
	fmt.Printf("Воркер %d завершил работу.\n", id)
}

func main() {
	dataChannel := make(chan string)

	// Создаем канал для ожидания завершения работы всех воркеров.
	var wg sync.WaitGroup

	numWorkers := 5

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataChannel, &wg)
	}

	// Обработчик сигнала Ctrl+C для завершения программы.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		close(dataChannel) // Закрываем канал данных при получении сигнала завершения.
	}()

	for i := 1; i <= 10; i++ {
		dataChannel <- fmt.Sprintf("Сообщение %d", i)
	}

	// Ожидаем завершения всех воркеров.
	wg.Wait()
	fmt.Println("Все воркеры завершили работу. Программа завершена.")
}
