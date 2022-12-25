package main

import "fmt"

type MailCategory int

const (
	Uncategorized MailCategory = iota
	Personal
	Spam
	Social
	Advertisements
)

type BitField int

const (
	Field1 BitField = 1 << iota
	Filed2
	Field3
	Field4
)

func main() {
	fmt.Println(Spam)
	// 2
	fmt.Println(Field4)
	// 8
}
