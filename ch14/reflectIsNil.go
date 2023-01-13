package main

import (
	"fmt"
	"reflect"
)

func hasNoValue(i interface{}) bool {
	iv := reflect.ValueOf(i)
	if !iv.IsValid() {
		return true
	}
	switch iv.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Func, reflect.Interface:
		return iv.IsNil()
	default:
		return false
	}
}

func main() {
	var f func() bool
	var s struct{}
	var p *int

	fmt.Println(hasNoValue(f))
	// true
	fmt.Println(hasNoValue(s))
	// false
	fmt.Println(hasNoValue(p))
	// true
}
