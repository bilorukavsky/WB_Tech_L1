/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/
package main

import (
	"fmt"
	"sync"
)

func Square(number int, wg *sync.WaitGroup) {
	square := number * number
	fmt.Printf("Квадрат %d: %d\n", number, square)
	defer wg.Done() // Уменьшаем счетчик, когда функция завершает выполнение
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup // Создаем объект WaitGroup для синхронизации горутин

	for _, number := range numbers {
		wg.Add(1)              // Увеличиваем счетчик для каждой горутины, которую мы создаем
		go Square(number, &wg) // Запускаем горутину для вычисления квадрата числа
	}

	wg.Wait() // Ожидаем завершения всех горутин

	testSquare2()
}
