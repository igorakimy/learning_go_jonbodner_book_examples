package main

import (
	"fmt"
	"os"
)

func doPanic(msg string) {
	panic(msg)
}

func div60(i int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()
	fmt.Println(60 / i)
}

func main() {
	for _, val := range []int{1, 2, 0, 6} {
		div60(val)
	}
	// 60
	// 30
	// runtime error: integer divide by zero
	// 10

	doPanic(os.Args[0])
	// panic: ~\go-build3573408948\b001\exe\panicAndrecoverFunctions.exe
	//
	// goroutine 1 [running]:
	// main.doPanic(...)
	//        D:/Projects/Go/go_idioms_and_design_patterns/ch8/panicAndrecoverFunctions.go:8
	// main.main()
	//        D:/Projects/Go/go_idioms_and_design_patterns/ch8/panicAndrecoverFunctions.go:12 +0x45
	// exit status 2
}
