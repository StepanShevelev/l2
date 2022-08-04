package main

import "fmt"

//Паттерн Комманда является поведенческим паттерном уровня объекта. Он позволяет
// представить запрос в виде объекта. Из этого следует, что команда - это объект.
// Такие запросы, например, можно ставить в очередь, отменять или возобновлять.
// Command - запрос в виде объекта на выполнение;
// Receiver - объект-получатель запроса, который будет обрабатывать нашу команду;
// Invoker* - объект-инициатор запроса.
// *Invoker умеет складывать команды в стопку и инициировать их выполнение по какому-то событию.
//  Обратившись к Invoker можно отменить команду, пока та не выполнена.
// Комманда отделяет объект, инициирующий операцию, от объекта, который знает, как ее выполнить.
// Единственное, что должен знать инициатор, это как отправить команду.

// Плюсы:

// Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
// Позволяет реализовать простую отмену и повтор операций.
// Позволяет реализовать отложенный запуск операций.
// Позволяет собирать сложные команды из простых.
// Реализует принцип открытости/закрытости.

// Минусы:
// Усложняет код программы из-за введения множества дополнительных классов.

type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

type Command interface {
	execute()
}

type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}

type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}

type Device interface {
	on()
	off()
}

type Tv struct {
	isRunning bool
}

func (t *Tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *Tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

func main() {
	tv := &Tv{}

	onCommand := &OnCommand{
		device: tv,
	}

	offCommand := &OffCommand{
		device: tv,
	}

	onButton := &Button{
		command: onCommand,
	}
	onButton.press()

	offButton := &Button{
		command: offCommand,
	}
	offButton.press()
}

//type Command interface {
//	Execute()
//	Undo()
//}
//
//type ConcreteCommand struct {
//	reciever Receiver
//	command  Command
//}
//
//func (cc ConcreteCommand) Execute() {
//	cc.reciever.Operation()
//}
//
//func (cc ConcreteCommand) Undo() {
//	cc.command.Undo()
//}
//
//type Receiver interface {
//	Operation()
//}
//
//type Invoker struct {
//	command Command
//}
//
//func (i Invoker) SetCommand(c Command) {
//	i.command = c
//}
//
//func (i Invoker) Run() {
//	i.command.Execute()
//}
//
//func (i Invoker) Cancel() {
//	i.command.Undo()
//}
