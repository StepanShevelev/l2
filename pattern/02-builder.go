package main

import "fmt"

// Паттерн Строитель является пораждающим паттерном проектирования. Он позволяет создавать сложные объекты
// используя шаги. На каждом шаге производится какая-то часть общего объекта. Тем самым, выполняя все шаги по очереди,
// формируется некий объект, представляющий сложную структуру. Строитель позволяет использовать один и тот же код
// строительства объекта для получения разных представлений этого объекта.

// Плюсы:
// Позволяет пошагово создавать общий объект, который зависит от составляющих частей;
// Позволяет использовать один и тот же код для создания различных объектов;
// Изолирует сложный код при сборке объекта и его бизнес-логики.

// Минусы:
// Усложняет код программы из-за введения дополнительных классов (структур, интерфейсов).
// Привязка клиента к конкретному объекту строителя, т.к. в интерфейсе может не быть какого-то метода,
// поэтому будет необходимо его добавить.

type House struct {
	windowType string
	doorType   string
	floor      int
}

type Builder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

func getBuilder(builderType string) Builder {
	if builderType == "normal" {
		return newNormalBuilder()
	}

	if builderType == "igloo" {
		return newIglooBuilder()
	}
	return nil
}

type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}

func (b *NormalBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *NormalBuilder) setNumFloor() {
	b.floor = 2
}

func (b *NormalBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (i *IglooBuilder) setWindowType() {
	i.windowType = "Snow Window"
}

func (i *IglooBuilder) setDoorType() {
	i.doorType = "Snow Door"
}

func (i *IglooBuilder) setNumFloor() {
	i.floor = 1
}

func (i *IglooBuilder) getHouse() House {
	return House{
		doorType:   i.doorType,
		windowType: i.windowType,
		floor:      i.floor,
	}
}

type Director struct {
	builder Builder
}

func newDirector(b Builder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b Builder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

func main() {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)

}
