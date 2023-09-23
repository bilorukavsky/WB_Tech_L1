/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Функция sumSquare2 вычисляет квадрат числа и добавляет его к общей сумме, используя атомарные операции.
func sumSquare2(number int, sum *int64, wg *sync.WaitGroup) {
	defer wg.Done()
	square := int64(number * number)
	atomic.AddInt64(sum, square) // Атомарно добавляем квадрат числа к общей сумме
}

func testSumSquare2() {
	numbers := []int{2, 4, 6, 8, 10}
	var sum int64 // Переменная для хранения суммы квадратов

	var wg sync.WaitGroup // Создаем объект WaitGroup для синхронизации горутин

	for _, number := range numbers {
		wg.Add(1) // Увеличиваем счетчик ожидаемых горутин
		go sumSquare2(number, &sum, &wg)
	}

	wg.Wait() // Ожидаем завершения всех горутин
	fmt.Printf("Сумма квадратов: %d\n", sum)
}
