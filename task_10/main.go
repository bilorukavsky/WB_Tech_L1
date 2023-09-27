/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов.
Последовательность в подмножноствах не важна.
Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/
package main

import (
	"fmt"
	"math"
)

func groupTemperatures(temperatures []float64, step int) map[int][]float64 {
	temperatureGroups := make(map[int][]float64)

	for _, temp := range temperatures {
		var group int
		if temp >= 0 {
			group = int(math.Floor(temp/float64(step))) * step // Округление вниз
		} else {
			group = int(math.Ceil(temp/float64(step))) * step // Округление вверх
		}

		temperatureGroups[group] = append(temperatureGroups[group], temp)
	}

	return temperatureGroups
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groupStep := 10

	temperatureGroups := groupTemperatures(temperatures, groupStep)

	for group, temps := range temperatureGroups {
		fmt.Printf("%d: %v\n", group, temps)
	}
}
