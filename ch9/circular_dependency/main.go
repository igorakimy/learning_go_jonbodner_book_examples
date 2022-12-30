package main

import (
	"fmt"
	"github.com/learning-go-book/circular_dependency/person"
	"github.com/learning-go-book/circular_dependency/pet"
)

func main() {
	fmt.Println(person.GetPersonsPets())
	fmt.Println(pet.GetPetsOwners())
	// package github.com/learning-go-book/circular_dependency
	//        imports github.com/learning-go-book/circular_dependency/person
	//        imports github.com/learning-go-book/circular_dependency/pet
	//        imports github.com/learning-go-book/circular_dependency/person:
	//        		import cycle not allowed
}
