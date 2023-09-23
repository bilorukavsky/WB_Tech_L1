/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/
package main

import (
	"fmt"
	"sync"
)

func sumSquare(number int, mu *sync.Mutex, wg *sync.WaitGroup, sum *int) {
	square := number * number
	mu.Lock() // Заблокировать мьютекс перед изменением общей суммы
	*sum += square
	mu.Unlock()     // Разблокировать мьютекс после завершения операции
	defer wg.Done() // Уменьшаем счетчик, когда функция завершает выполнение
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var mu sync.Mutex // Мьютекс для синхронизации доступа к общей сумме
	var sum int
	var wg sync.WaitGroup // Создаем объект WaitGroup для синхронизации горутин

	for _, number := range numbers {
		wg.Add(1) // Увеличиваем счетчик для каждой горутины, которую мы создаем
		go sumSquare(number, &mu, &wg, &sum)
	}

	wg.Wait() // Ожидаем завершения всех горутин
	fmt.Printf("Сумма квадратов: %d\n", sum)

	testSumSquare2()
	testSumSquare3()
}
