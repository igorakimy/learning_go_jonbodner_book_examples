package person

import (
	"github.com/learning-go-book/circular_dependency/pet"
)

type Person struct {
	Name    string
	Age     int
	PetName string
}

var pets = map[string]pet.Pet{
	"Fluffy": {"Fluffy", "Cat", "Bob"},
	"Rex":    {"Rex", "Dog", "Julia"},
}

func GetPersonsPets() map[string]pet.Pet {
	return pets
}
