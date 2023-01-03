package main

import (
	"fmt"
)

func main() {
	in := make(chan int)
	in2 := make(chan int)
	done := make(chan struct{})

	for i := 0; i < 10; i++ {
		in <- i
	}

	for i := 0; i < 10; i++ {
		in2 <- i
	}

	done <- struct{}{}

	for {
		select {
		case v, ok := <-in:
			if ok {
				in = nil // Эта ветвь больше не будет успешно выполняться!
				continue
			}
			fmt.Println(v)
			// Обработка значения переменной v, считанного из канала in
		case v, ok := <-in2:
			if !ok {
				in2 = nil // Эта ветвь больше не будет успешно выполнятся!
				continue
			}
			// Обработка значения переменной v, считанного из канала in2
			fmt.Println(v)
		case <-done:
			return
		}
	}
}
