/*
Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
*/
package main

import (
	"fmt"
)

func checkType(val interface{}) {
	switch val.(type) { // Определяем тип переменной, можно использовать только в сочетании с switch
	case int:
		fmt.Println("Это int")
	case string:
		fmt.Println("Это string")
	case bool:
		fmt.Println("Это bool")
	case chan int:
		fmt.Println("Это channel")
	default:
		fmt.Println("Неизвестный тип")
	}
}

func main() {
	var a int = 42
	var b string = "Hello, world!"
	var c bool = true
	var ch chan int = make(chan int)

	checkType(a)
	checkType(b)
	checkType(c)
	checkType(ch)
}
