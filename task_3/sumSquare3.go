/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/
package main

import (
	"fmt"
	"sync"
)

// calculateSquare вычисляет квадрат числа и отправляет его в resultChan.
func sumSquare3(number int, wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done()
	square := number * number
	resultChan <- square
}

func testSumSquare3() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	resultChan := make(chan int, len(numbers)) // Создаем буферизированный канал

	// Запускаем горутины для вычисления квадратов чисел и отправки их в канал.
	for _, number := range numbers {
		wg.Add(1) // Увеличиваем счетчик ожидаемых горутин
		go sumSquare3(number, &wg, resultChan)
	}

	// Завершаем работу горутины после завершения всех горутин-вычислителей.
	go func() {
		wg.Wait()         // Ожидаем завершения всех горутин-вычислителей
		close(resultChan) // Закрываем канал, чтобы сигнализировать, что больше данных не будет
	}()

	var sum int
	// Считываем результаты из канала и суммируем их.
	for square := range resultChan {
		sum += square
	}

	fmt.Printf("Сумма квадратов: %d\n", sum)
}
