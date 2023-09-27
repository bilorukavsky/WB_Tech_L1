/*
Поменять местами два числа без создания временной переменной.
*/
package main

import "fmt"

func swapWithoutTemp(a, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b
	return a, b
}

func swapWithoutTemp2(a, b int) (int, int) {
	a = a * b
	b = a / b
	a = a / b
	return a, b
}

func main() {
	x := 5
	y := 10
	x, y = swapWithoutTemp(x, y)
	fmt.Printf("После обмена: x = %d, y = %d\n", x, y)
	x, y = swapWithoutTemp2(x, y)
	fmt.Printf("После обмена: x = %d, y = %d\n", x, y)
}
