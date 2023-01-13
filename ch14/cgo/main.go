package main

import (
	"fmt"
)

/*
#cgo LDFLAGS: -lm
#include <stdio.h>
#include <math.h>
#include "mylib.h"

int add(int a, int b) {
	int sum = a + b;
	printf("a: %d, b: %d, sum: %d\n", a, b, sum);
	return sum;
}
*/
import "C"

func main() {
	sum := C.add(3, 2)
	// a: 3, b: 2, sum: 5
	fmt.Println(sum)
	// 5
	fmt.Println(C.sqrt(100))
	// 10
	fmt.Println(C.multiply(10, 20))
	// 200
}
