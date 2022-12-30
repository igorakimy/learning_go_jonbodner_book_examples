package pet

import (
	"github.com/learning-go-book/circular_dependency/person"
)

type Pet struct {
	Name      string
	Kind      string
	OwnerName string
}

var owners = map[string]person.Person{
	"Bob":   {"Bob", 30, "Fluffy"},
	"Julia": {"Julia", 40, "Rex"},
}

func GetPetsOwners() map[string]person.Person {
	return owners
}
