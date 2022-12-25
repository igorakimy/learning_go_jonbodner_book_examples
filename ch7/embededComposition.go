package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee // Встроенное поле
	Reports  []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	// Выполнение бизнес-логики
	var employees []Employee
	return employees
}

func main() {
	m := Manager{
		Employee: Employee{
			Name: "Bob Bobson",
			ID:   "12345",
		},
		Reports: []Employee{},
	}
	fmt.Println(m.ID)
	// 12345
	fmt.Println(m.Description())
	// Bob Bobson (12345)
}
