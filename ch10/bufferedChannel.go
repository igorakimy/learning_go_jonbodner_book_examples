package main

import (
	"fmt"
)

func process(num int) int {
	return num * 2
}

func processChannel(ch chan int) []int {
	const conc = 10
	results := make(chan int, conc)
	for i := 0; i < conc; i++ {
		go func() {
			v := <-ch
			results <- process(v)
		}()
	}
	var out []int
	for i := 0; i < conc; i++ {
		out = append(out, <-results)
	}
	return out
}

func main() {
	channel := make(chan int)
	for i := 0; i < 10; i++ {
		go func(val int) {
			channel <- val
		}(i)
	}
	res := processChannel(channel)
	fmt.Println(res)
}
