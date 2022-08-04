package main

// Паттерн Стратегия является поведенческим паттерном. Он определяет схожие алгоритмы
// и помещает каждый из них в отдельную структуру. После чего, алгоритмы могут
// взаимодействовать в исполняемой программе.
// Решает проблему часто расширширяющихся и изменяющихся алгоритмов, путем
// выноса их в собственный объект. Применяется в случае необходимости использования
// разных вариантов одного алгоритма внутри одного объекта.

// Плюсы:
// Замена алгоритмов налету;
// Изоляция кода и данных алгоритмов от остальных объектов бизнес-логики;
// Уход от наследования;
// Реализует принцип open-closed.

// Минусы:
// Усложнение программы за счет дополнительных объектов;
// Необходимость клиенту знать, в чем состоит разница между стратегиями, чтобы выбрать подходящую.

import "fmt"

type System interface {
	CheckSystem()
}

type Windows struct {
	System
}

func (w Windows) CheckSystem() {
	fmt.Println("Это система на windows")
}

type Linux struct {
	System
}

func (l Linux) CheckSystem() {
	fmt.Println("Это система на linux")
}

type Computer struct {
	system System
}

func (c Computer) CheckSystem() {
	c.system.CheckSystem()
}

func main() {
	comp := Computer{system: Windows{}}
	comp.CheckSystem()
	comp.system = Linux{}
	comp.CheckSystem()
}
