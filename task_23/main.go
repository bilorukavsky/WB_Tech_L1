/*
Удалить i-ый элемент из слайса.
*/
package main

import "fmt"

func removeElement(slice []int, i int) []int {
	if i >= 0 && i < len(slice) {
		copy(slice[i:], slice[i+1:])
		return slice[:len(slice)-1]
	}
	return slice
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	i := 2
	slice = removeElement(slice, i)
	fmt.Printf("Слайс полсе удалинеия: %d\n", slice)

}
