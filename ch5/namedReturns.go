package main

import (
	"errors"
	"fmt"
)

func divAndReminder(numerator int, denominator int) (result int, remainder int,
	err error) {
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return result, remainder, err
	}
	result, remainder = numerator/denominator, numerator%denominator
	return result, remainder, err
}

func main() {
	x, y, z := divAndReminder(5, 2)
	fmt.Println(x, y, z)
	// 2 1 <nil>
}
