/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{5, 2, 9, 3, 6}

	// Используем функцию sort.Ints() для сортировки массива
	sort.Ints(arr) // Так же есть sort.Float64s для чисел с плавающей точкой, или sort.Strings для строк.

	fmt.Println("Отсортированный массив:", arr)
}
