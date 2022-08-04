package main

import "fmt"

// Паттерн Посетитель является поведенческим паттерном уровня объекта. Он позволяет обойти
// набор элементов (объектов) с разнородными интерфейсами, а также позволяет добавить
// новый метод в тип объекта, при этом, не изменяя сам тип этого объекта.
// Решает задачу определения новой операции, не изменяя типы объектов, над которыми выполняется одна или более операций.
// Паттерн следует применять если:
// 1. Имеются различные объекты разных типов с разными интерфейсами, но над ними нужно совершать операции, зависящие от конкретных типов;
// 2. Необходимо над структурой выполнить различные, усложняющие структуру операции;
// 3. Часто добавляются новые операции над структурой.

// Плюсы:
// 1. Упрощается добавление новых операций;
// 2. Объединение родственных операции Посетителе;
// 3. Посетитель может запоминать в себе какое-то состояние по мере обхода контейнера.

// Минусы:
// 1. Затруднено добавление новых типов, поскольку нужно обновлять иерархию Посетителя и его сыновей.

type Shape interface {
	getType() string
	accept(Visitor)
}

type Square struct {
	side int
}

func (s *Square) accept(v Visitor) {
	v.visitForSquare(s)
}

func (s *Square) getType() string {
	return "Square"
}

type Circle struct {
	radius int
}

func (c *Circle) accept(v Visitor) {
	v.visitForCircle(c)
}

func (c *Circle) getType() string {
	return "Circle"
}

type Rectangle struct {
	l int
	b int
}

func (t *Rectangle) accept(v Visitor) {
	v.visitForRectangle(t)
}

func (t *Rectangle) getType() string {
	return "rectangle"
}

type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForRectangle(*Rectangle)
}

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) visitForSquare(s *Square) {
	// Calculate area for square.
	// Then assign in to the area instance variable.
	fmt.Println("Calculating area for square")
}

func (a *AreaCalculator) visitForCircle(s *Circle) {
	fmt.Println("Calculating area for circle")
}
func (a *AreaCalculator) visitForRectangle(r *Rectangle) {
	fmt.Println("Calculating area for rectangle")
}

type MiddleCoordinates struct {
	x int
	y int
}

func (a *MiddleCoordinates) visitForSquare(s *Square) {
	// Calculate middle point coordinates for square.
	// Then assign in to the x and y instance variable.
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *MiddleCoordinates) visitForCircle(c *Circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}
func (a *MiddleCoordinates) visitForRectangle(r *Rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}

func main() {
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{l: 2, b: 3}

	areaCalculator := &AreaCalculator{}

	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &MiddleCoordinates{}
	square.accept(middleCoordinates)
	circle.accept(middleCoordinates)
	rectangle.accept(middleCoordinates)
}
