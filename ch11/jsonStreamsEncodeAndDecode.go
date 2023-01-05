package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type SomePerson struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func process(p SomePerson) SomePerson {
	return SomePerson{
		Name: fmt.Sprintf("%s Johnson", p.Name),
		Age:  p.Age,
	}
}

func main() {
	data := `
		{"name": "Fred", "age": 40}
		{"name": "Mary", "age": 21}
		{"name": "Pat", "age": 30}`

	dec := json.NewDecoder(strings.NewReader(data))
	var t struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	// Чтение нескольких значений
	for dec.More() {
		err := dec.Decode(&t)
		if err != nil {
			panic(err)
		}
		fmt.Println(t.Name)
	}
	// Fred
	// Mary
	// Pat

	// Запись нескольких значений
	var allInputs []SomePerson
	for i := 0; i < 10; i++ {
		allInputs = append(allInputs, SomePerson{
			Name: "Fred",
			Age:  40,
		})
	}
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	for _, input := range allInputs {
		t2 := process(input)
		err := enc.Encode(t2)
		if err != nil {
			panic(err)
		}
	}
	out := b.String()
	fmt.Println(out)
}
