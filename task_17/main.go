/*
Реализовать бинарный поиск встроенными методами языка.
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	// Создаем срез с отсортированными данными.
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	target := 5

	// Используем sort.Search для выполнения бинарного поиска.
	index := sort.Search(len(data), func(i int) bool {
		return data[i] >= target
	})

	// Проверяем, был ли элемент найден.
	if index < len(data) && data[index] == target {
		fmt.Printf("Элемент %d найден в позиции %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден\n", target)
	}
}
