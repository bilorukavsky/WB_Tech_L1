/*
Реализовать паттерн «адаптер» на любом примере.
*/
package main

import "fmt"

type Square struct {
	Side float64
}

func (s *Square) GetArea() float64 {
	return s.Side * s.Side
}

type Shape interface {
	GetArea() float64
}

// Адаптер для преобразования Square в Shape
type ShapeAdapter struct {
	Square *Square
}

func (sa ShapeAdapter) GetArea() float64 {
	return sa.Square.GetArea()
}

func main() {
	square := &Square{Side: 5}
	fmt.Printf("Площадь квадрата: %.2f\n", square.GetArea())

	shapeAdapter := ShapeAdapter{Square: square}
	fmt.Printf("Площадь: %.2f\n", shapeAdapter.GetArea())
}
