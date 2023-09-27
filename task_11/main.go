/*
Реализовать пересечение двух неупорядоченных множеств.
*/
package main

import "fmt"

func intersection(set1, set2 []int) []int {
	// Создаем карты для хранения элементов множеств.
	// Ключи будут элементами множества, значения - флагами присутствия.
	set1Map := make(map[int]bool)
	set2Map := make(map[int]bool)

	for _, item := range set1 {
		set1Map[item] = true
	}

	for _, item := range set2 {
		set2Map[item] = true
	}

	// Создаем результат как пустой срез.
	result := []int{}

	// Итерируемся по элементам первого множества и проверяем их наличие во втором множестве.
	for item := range set1Map {
		if set2Map[item] {
			result = append(result, item)
		}
	}

	return result
}

func printSet(set []int) {
	fmt.Print("{")
	for i, item := range set {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(item)
	}
	fmt.Println("}")
}

func main() {
	set1 := []int{3, 2, 1, 5, 4}
	set2 := []int{7, 4, 6, 5, 3}
	intersection := intersection(set1, set2)

	fmt.Print("Пересечение: ")
	printSet(intersection)

	task_11_2()

}
