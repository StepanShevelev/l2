package main

import "fmt"

// Паттерн Состояние является поведенческим паттерном уровня объекта. Он позволяет
// позволяет объекту изменять свое поведение в зависимости от внутреннего состояния.
// Поведение объекта изменяется настолько, что создается впечатление, будто изменился тип объекта.
// Паттерн следует применять:
// Когда поведение объекта зависит от его состояния;
// Поведение объекта должно изменяться во время выполнения программы;
// Состояний достаточно много и использовать для этого условные операторы, разбросанные по коду, достаточно затруднительно.

type Water struct {
	state string
}

func (w *Water) Heat() {
	if w.state == "лед" {
		fmt.Println("Лед стал жидкостью")
		w.state = "жидкость"
	} else if w.state == "жидкость" {
		fmt.Println("Жидкость стала паром")
		w.state = "пар"
	}
}

func (w *Water) Frost() {
	if w.state == "жидкость" {
		fmt.Println("Жидкость стала льдом")
		w.state = "лед"
	} else if w.state == "пар" {
		fmt.Println("Пар стал жидкостью")
		w.state = "жидкость"
	}
}

func main() {
	water := Water{state: "жидкость"}
	water.Frost()
	water.Heat()
	water.Heat()
	water.Frost()
}