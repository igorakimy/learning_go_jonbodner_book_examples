package main

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobbert", 23},
		{"Fred", "Fredson", 18},
	}
	fmt.Println(people)
	// [{Pat Patterson 37} {Tracy Bobbert 23} {Fred Fredson 18}]

	// сортировка по фамилии
	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)
	// [{Tracy Bobbert 23} {Fred Fredson 18} {Pat Patterson 37}]

	// сортировка по возрасту
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)
	// [{Fred Fredson 18} {Tracy Bobbert 23} {Pat Patterson 37}]
}
