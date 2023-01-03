package main

import (
	"fmt"
)

func countTo(max int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < max; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func main() {
	for i := range countTo(10) {
		fmt.Println(i)
	}

	// Здесь горутина заблокируется и будет бесконечно ждать,
	// когда из канала будет считано еще одно значение
	for i := range countTo(10) {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
}
