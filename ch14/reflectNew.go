package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Создать переменную reflect.Type, если значение отсутствует
	var stringType = reflect.TypeOf((*string)(nil)).Elem()
	var stringSliceType = reflect.TypeOf([]string(nil))
	fmt.Println(stringType, stringSliceType)
	// string []string

	ssv := reflect.MakeSlice(stringSliceType, 0, 10)
	sv := reflect.New(stringType).Elem()
	sv.SetString("hello")
	ssv = reflect.Append(ssv, sv)
	ss := ssv.Interface().([]string)
	fmt.Println(ss)
	// [hello]
}
