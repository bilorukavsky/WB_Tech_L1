/*
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
*/
package main

import (
	"fmt"
	"math/big"
)

func multiply(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func divide(a, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b)
}

func add(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func subtract(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func main() {
	a := new(big.Int)
	b := new(big.Int)

	a.SetString("2222222", 10) // Здесь устанавливаем значение a > 2^20
	b.SetString("3333333", 10) // Здесь устанавливаем значение b > 2^20

	multiplication := multiply(a, b)
	fmt.Printf("Умножение: %s\n", multiplication.String())

	division := divide(a, b)
	fmt.Printf("Деление: %s", division.String())

	addition := add(a, b)
	fmt.Printf("Сложение: %s\n", addition.String())

	subtraction := subtract(a, b)
	fmt.Printf("Вычитание: %s\n", subtraction.String())
}
