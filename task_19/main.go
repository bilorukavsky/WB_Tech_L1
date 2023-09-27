/*
Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
Символы могут быть unicode.
*/
package main

import (
	"fmt"
)

func reverseString(input string) string {
	// Преобразование строки в массив рун (Unicode символов)
	runes := []rune(input)

	length := len(runes)

	// Переворачиваем строку, меняя местами символы
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	input := "главрыба —"

	reversed := reverseString(input)

	fmt.Println(reversed)
}
