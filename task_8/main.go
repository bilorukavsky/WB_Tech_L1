/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/
package main

import (
	"fmt"
)

func setBit(num int64, pos uint, value uint) int64 {
	// Создаем битовую маску, которая имеет 1 в позиции pos, остальные биты - 0.
	mask := int64(1) << pos

	// Если value равно 1, устанавливаем бит, иначе сбрасываем бит.
	if value == 1 {
		num |= mask
	} else {
		num &= ^mask
	}

	return num
}

func main() {
	var num int64 = 42
	pos := uint(2)
	value := uint(1)

	result := setBit(num, pos, value)
	fmt.Printf("Число после установки %d-го бита в %d: %d\n", pos, value, result)
}
