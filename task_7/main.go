/*
Реализовать конкурентную запись данных в map.
*/
package main

import (
	"fmt"
	"sync"
)

// writeData записывает данные в map с использованием мьютекса для конкурентной записи.
func writeData(data map[string]int, key string, value int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик ожидания после завершения горутины.

	mutex.Lock() // Захватываем мьютекс перед записью в map.
	data[key] = value
	mutex.Unlock() // Освобождаем мьютекс после записи.
}

// readData читает данные из map с использованием мьютекса для конкурентного доступа.
func readData(data map[string]int, mutex *sync.Mutex) {
	mutex.Lock() // Захватываем мьютекс перед чтением из map.
	for key, value := range data {
		fmt.Printf("%s: %d\n", key, value)
	}
	mutex.Unlock() // Освобождаем мьютекс после чтения.
}

func main() {
	data := make(map[string]int)
	var mutex sync.Mutex

	var wg sync.WaitGroup
	numRoutines := 5

	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go writeData(data, fmt.Sprintf("Ключ%d", i), i*10, &mutex, &wg)
	}

	wg.Wait() // Ожидаем завершения всех горутин.

	readData(data, &mutex)
}
