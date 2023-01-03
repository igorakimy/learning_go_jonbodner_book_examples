package main

import (
	"fmt"
	"strings"
	"sync"
)

type SlowComplicatedParser interface {
	Parse(string) string
}

var parser SlowComplicatedParser
var once sync.Once

func Parse(dataToParse string) string {
	once.Do(func() {
		fmt.Println("call in callback")
		parser = initParser()
	})
	return parser.Parse(dataToParse)
}

type SomeParser struct{}

func (sp SomeParser) Parse(data string) string {
	return strings.ToUpper(data)
}

func initParser() SlowComplicatedParser {
	// здесь выполняются различные операции настройки и загрузки
	return SomeParser{}
}

func main() {
	fmt.Println(Parse("hello"))
	// call in callback
	// HELLO
	fmt.Println(Parse("world"))
	// WORLD
}
