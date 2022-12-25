package main

import (
	"fmt"
)

type Inner struct {
	X int
}

type Outer struct {
	Inner
	X int
}

func main() {
	o := Outer{
		Inner: Inner{
			X: 10,
		},
		X: 20,
	}
	fmt.Println(o.X)
	// 20
	fmt.Println(o.Inner.X)
}
