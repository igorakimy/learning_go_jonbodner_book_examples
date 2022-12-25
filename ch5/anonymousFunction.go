package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println("printing", j, "from inside an anonymous function")
		}(i)
	}
	// printing 0 from inside an anonymous function
	// printing 1 from inside an anonymous function
	// printing 2 from inside an anonymous function
	// printing 3 from inside an anonymous function
	// printing 4 from inside an anonymous function
}
