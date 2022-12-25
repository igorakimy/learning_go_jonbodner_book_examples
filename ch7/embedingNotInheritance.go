package main

import (
	"fmt"
)

type Inner2 struct {
	A int
}

func (i Inner2) IntPointer(val int) string {
	return fmt.Sprintf("Inner2: %d", val)
}

func (i Inner2) Double() string {
	return i.IntPointer(i.A * 2)
}

type Outer2 struct {
	Inner2
	S string
}

func (o Outer2) IntPointer(val int) string {
	return fmt.Sprintf("Outer2: %d", val)
}

func main() {
	o := Outer2{
		Inner2: Inner2{
			A: 10,
		},
		S: "Hello",
	}
	fmt.Println(o.Double())
	// Inner2: 20
}
