/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/
package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) SayHello() {
	fmt.Printf("Привет, меня зовут %s\n", h.Name)
}

type Action struct {
	Human      Human
	NameAction string
}

func (a *Action) DoAction() {
	fmt.Printf("%s делает: %s\n", a.Human.Name, a.NameAction)

}

func main() {
	// Создаем объект Action
	a := Action{
		Human: Human{
			Name: "Иван",
			Age:  30,
		},
		NameAction: "бег",
	}

	// Вызываем методы для Action, которые также доступны из Human
	a.Human.SayHello()
	a.DoAction()
}
