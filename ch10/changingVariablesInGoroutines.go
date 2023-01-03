package main

import (
	"fmt"
)

func main() {
	a := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(a))
	for _, v := range a {
		// Для решения этой трудноуловимой проблемы необходимо либо затенить
		// переменную: v := v
		// либо использовать параметр для горутины:
		// go func(val int) {
		// 		ch <- val * 2
		// }(v)
		go func() {
			// Трудноуловимая ошибка
			ch <- v * 2
		}()
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}
}
