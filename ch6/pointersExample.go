package main

import (
	"fmt"
)

func main() {
	x := 10
	var PointerToX *int
	PointerToX = &x
	fmt.Println(*PointerToX)
	// 10

	// Создать переменную указательного типа. Функция new() возвращает указатель
	// на экземпляр нулевого значения указанного типа.
	var s = new(int)
	fmt.Println(s == nil)
	// false
	fmt.Println(*s)
	// 0
}
