package main

import (
	"fmt"
)

func division(numerator, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

func main() {
	result := division(5, 2)
	fmt.Println(result)
	// 2
}
