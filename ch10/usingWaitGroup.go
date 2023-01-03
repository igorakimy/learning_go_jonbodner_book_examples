package main

import (
	"fmt"
	"sync"
)

func processorFunc(num int) int {
	return num * 2
}

// Конкурентная обработка значений канала, заносит результаты
// в один срез и возвращает его

func processAndGather(in <-chan int, processor func(int) int, num int) []int {
	out := make(chan int, num)
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			for v := range in {
				out <- processor(v)
			}
		}()
	}
	// Запуск следящей горутины, которая дожидается завершения
	// всех обрабатывающих горутин
	go func() {
		wg.Wait()
		// После их завершения следящая горутина вызывает функцию
		// close для выходного канала
		close(out)
	}()
	var result []int
	for v := range out {
		result = append(result, v)
	}
	return result
}

func doThing1() {
	fmt.Println("do thing 1")
}

func doThing2() {
	fmt.Println("do thing 2")
}

func doThing3() {
	fmt.Println("do thing 3")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		doThing1()
	}()
	go func() {
		defer wg.Done()
		doThing2()
	}()
	go func() {
		defer wg.Done()
		doThing3()
	}()
	wg.Wait()

	// Конкурентная обработка значений канала
	n := 5
	wg.Add(n)
	in := make(chan int, n)
	for i := 0; i < n; i++ {
		go func(val int) {
			defer wg.Done()
			in <- val
		}(i)
	}
	go func() {
		wg.Wait()
		close(in)
	}()
	res := processAndGather(in, processorFunc, n)
	fmt.Println(res)
}
