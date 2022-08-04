package main

import "fmt"

// Фабричный метод — это порождающий паттерн проектирования,
//который решает проблему создания различных продуктов, без указания конкретных классов продуктов.

// Плюсы:
// Избавляет от привязки к конкретному типу объекта при помощи конструктора;
// Упрощает добавление новых объектов от базового класса;
// Реализует принцип open-closed.

// Минусы:
// Может привести к созданию больших иерапрхий объектов (большое количество структур, которые сложны при сопровождении);
// Появляется один главный конструктор, к которому будет привязана вся логика программы.

type iGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getPower() int {
	return g.power
}

type Ak47 struct {
	Gun
}

func newAk47() iGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

type musket struct {
	Gun
}

func newMusket() iGun {
	return &musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}

func getGun(gunType string) (iGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}

func printDetails(g iGun) {
	fmt.Printf("Gun: %s\n", g.getName())
	fmt.Printf("Power: %d\n", g.getPower())

}
func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printDetails(ak47)
	printDetails(musket)
}
