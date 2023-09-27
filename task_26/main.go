/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.
Например:
abcd — true
abCdefAaf — false
aabcd — false
*/
package main

import (
	"fmt"
	"strings"
)

func IsUnique(str string) bool {
	// Приводим строку к нижнему регистру
	str = strings.ToLower(str)

	// Создаем карту для отслеживания уникальных символов
	charMap := make(map[rune]bool)

	for _, char := range str {
		// Если символ уже встречался, возвращаем false
		if charMap[char] {
			return false
		}
		// Иначе отмечаем символ как встреченный
		charMap[char] = true
	}

	return true
}

func main() {
	fmt.Println(IsUnique("abcd"))         // true
	fmt.Println(IsUnique("abCdefAaf"))    // false
	fmt.Println(IsUnique("aabcd"))        // false
	fmt.Println(IsUnique("Hello, World")) // false
}
