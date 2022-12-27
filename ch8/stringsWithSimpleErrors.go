package main

import (
	"errors"
	"fmt"
)

// Первый способ создания ошибки из строки

func doubleEvenErr(i int) (int, error) {
	if i%2 != 0 {
		return 0, errors.New("only even numbers are processed")
	}
	return i * 2, nil
}

// Второй способ создания ошибки из строки

func doubleEvenFmt(i int) (int, error) {
	if i%2 != 0 {
		return 0, fmt.Errorf("%d isn't an even number", i)
	}
	return i * 2, nil
}

func main() {
	if _, err := doubleEvenErr(3); err != nil {
		fmt.Println(err)
		// only even numbers are processed
	}

	if _, err := doubleEvenFmt(5); err != nil {
		fmt.Println(err)
		// 5 isn't an even number
	}
}
