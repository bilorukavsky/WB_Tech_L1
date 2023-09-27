/*
Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.
*/
package main

import (
	"fmt"
	"sync"
)

// Функция для чтения чисел из input, удвоения их и записи в output
func processNumbers(input <-chan int, output chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for x := range input {
		result := x * 2
		output <- result
	}
	close(output)
}

// Функция для вывода результатов из output в stdout
func printResults(output <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for result := range output {
		fmt.Println(result)
	}
}

func main() {
	input := make(chan int)
	output := make(chan int)

	// Создаем WaitGroup для ожидания завершения горутин
	var wg sync.WaitGroup

	// Запускаем горутину для чтения из первого канала и удвоения числа
	wg.Add(1)
	go processNumbers(input, output, &wg)

	// Запускаем горутину для вывода результатов в stdout
	wg.Add(1)
	go printResults(output, &wg)

	numbers := []int{1, 2, 3, 4, 5}

	// Пишем числа в первый канал
	for _, num := range numbers {
		input <- num
	}
	close(input) // Закрываем первый канал, чтобы завершить горутину чтения

	// Ожидаем завершения горутин
	wg.Wait()
}
