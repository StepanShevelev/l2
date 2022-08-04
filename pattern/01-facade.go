package main

import "fmt"

// Паттерн Фасад является структурным паттерном проектирования. Представляет собой простой интерфейс к сложной системе.
// Если имеется много разных подсистем, которые используют свои интерфейсы и реализуют свой функционал поведения,
// следуюет применить паттерн Фасад, чтобы создать простой интерфейс для максимально простого взаимодействия с подсистемами.

// Плюсы:
// Изолирует клиентов от поведения сложной системы
// Минусы:
// Интерфейс Фасада может стать супер-объектом (супер-классом).
// Другими словами, все последующие функции будут проходить через этот объект.
type (
	SubSystemA struct {
	}

	SubSystemB struct {
	}

	Facade struct {
		subSystemA *SubSystemA
		subSystemB *SubSystemB
	}
)

func (sA *SubSystemA) A1() {
	fmt.Println("Operation A")
}

func (sB *SubSystemB) B1() {
	fmt.Println("Operation B")
}

func (f *Facade) Operation() {
	f.subSystemA.A1()
	f.subSystemB.B1()
}

func main() {
	f := new(Facade)
	f.Operation()
}
