package main

import (
	"encoding/json"
	"fmt"
)

type Foo struct {
	Field1 string
	Field2 int
}

// Не делайте так

func MakeFoo(f *Foo) error {
	f.Field1 = "val"
	f.Field2 = 20
	return nil
}

// Делайте так

func MakeFoo2() (Foo, error) {
	f := Foo{
		Field1: "val",
		Field2: 20,
	}
	return f, nil
}

func main() {
	f := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{}
	err := json.Unmarshal([]byte(`{"name": "Bob", "age": 20}`), &f)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f)
	}
}
