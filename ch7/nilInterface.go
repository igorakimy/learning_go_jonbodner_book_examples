package main

import (
	"fmt"
)

func main() {
	var s *string
	fmt.Println(s == nil)
	// true
	var i interface{}
	fmt.Println(i == nil)
	// true
	fmt.Println(i == s)
	// false
}
