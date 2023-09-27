/*
Реализовать пересечение двух неупорядоченных множеств.
*/
package main

import "fmt"

func intersection2(set1, set2 []int) []int {
	var result []int

	for _, elem1 := range set1 {
		for _, elem2 := range set2 {
			if elem1 == elem2 {
				result = append(result, elem1)
				break
			}
		}
	}

	return result
}

func task_11_2() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}

	intersectionResult := intersection2(set1, set2)
	fmt.Println("Пересечение:", intersectionResult)
}
