/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/
package main

import "fmt"

func Square2(number int, resultChan chan int) {
	square := number * number
	resultChan <- square // Отправляем результат в канал
}

func testSquare2() {
	numbers := []int{2, 4, 6, 8, 10}
	resultChan := make(chan int, len(numbers)) // Создаем буферизированный канал

	for _, number := range numbers {
		go Square2(number, resultChan) // Запускаем горутины для вычисления квадратов
	}

	// Считываем результаты из канала и выводим их
	for range numbers {
		square := <-resultChan // Считываем значение из канала
		fmt.Printf("Квадрат : %d\n", square)
	}
	close(resultChan) // Закрываем канал
}
