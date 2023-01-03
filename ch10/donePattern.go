package main

import (
	"fmt"
	"strings"
)

// Передаем одни и те же данные нескольким функциям и получаем
// результат только от самой быстрой функции

func searchData(s string, searchers []func(string) []string) []string {
	done := make(chan []struct{})
	result := make(chan []string)
	for _, searcher := range searchers {
		go func(searcher func(string) []string) {
			select {
			case result <- searcher(s):
			case <-done:
			}
		}(searcher)
	}
	r := <-result
	close(done)
	return r
}

func searcher1(s string) []string {
	var matches []string
	for _, word := range strings.Split(s, " ") {
		if word == "hello" {
			matches = append(matches, word)
		}
	}
	return matches
}

func searcher2(s string) []string {
	return []string{
		"my",
		"name",
		"age",
	}
}

func main() {
	str := "hello my best friend and hello again"
	searchers := []func(string) []string{
		searcher2,
		searcher1,
	}
	res := searchData(str, searchers)
	fmt.Println(res)
}
