package main

import "fmt"

func countTo2(max int) (<-chan int, func()) {
	// Создать канал для возвращения данных
	ch := make(chan int)
	// Создать сигнальный канал
	done := make(chan struct{})
	// Вместо того чтобы возвращать канал done напрямую, создаем замыкание,
	// которое закрывает канал done и возвращаем это замыкание
	cancel := func() {
		close(done)
	}
	go func() {
		for i := 0; i < max; i++ {
			select {
			case <-done:
				return
			default:
				ch <- i
			}
		}
		close(ch)
	}()
	return ch, cancel
}

func main() {
	ch, cancel := countTo2(10)
	for i := range ch {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
	cancel()

}
