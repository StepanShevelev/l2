package main

import (
	"fmt"
)

func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}

func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}

func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}

//Ответ:
//2
//1
//В связи с тем, что в первом случае определена возвращаемая переменная,
//и она же является изменяемой в defer-е, ее значение также поменяется, т.к. defer вызывается после return.
//
//Во втором случае переменная для возврата опеределяется в return,
//а изменение этой переменной происходит после определения. Поэтому в функцию main вернется неизмененной значение,
//при этом в defer-е оно все равно изменится.
