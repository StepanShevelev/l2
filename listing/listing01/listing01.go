package main

import (
	"fmt"
)

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4]
	fmt.Println(b)
}

//Ответ:[77 78 79]
//В результате получится срез от второго до пятого числа.
