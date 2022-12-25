package main

import (
	"fmt"
	"io"
)

func doThings(i interface{}) {
	switch j := i.(type) {
	case nil:
		// переменная i равна nil, переменная j обладает типом interface{}
	case int:
		// переменная j обладает типом int
	case MyInt:
		// переменная j обладает типом MyInt
	case io.Reader:
		// переменная j обладает типом io.Reader
	case string:
		// переменная j обладает типом string
	case bool, rune:
		// переменная i содержит булево значение или руну,
		// поэтому переменная j обладает типом interface{}
	default:
		// неизвестно, что содержит переменная i, поэтому переменная j
		// обладает типом interface{}
		fmt.Println(j)
	}
}

type MyInt int

func main() {
	var i interface{}
	var mine MyInt = 20
	i = mine
	i2 := i.(MyInt)
	fmt.Println(i2 + 1)
	// 21
}
