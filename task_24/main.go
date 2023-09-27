/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/
package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

// Конструктор для создания экземпляра Point
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (p1 *Point) DistanseTo(p2 *Point) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	return math.Sqrt(dx*dx + dy*dy) //Вычесляем расстояние
}

func main() {
	point1 := NewPoint(1.0, 2.0)
	point2 := NewPoint(4.0, 6.0)

	distanse := point1.DistanseTo(point2)
	fmt.Printf("Расстояние между двумя точками: %f\n", distanse)
}
