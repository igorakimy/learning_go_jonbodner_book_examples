package main

import (
	"fmt"
	"reflect"
)

func Filter(slice interface{}, filter interface{}) interface{} {
	sv := reflect.ValueOf(slice)
	fmt.Println(sv)
	fv := reflect.ValueOf(filter)
	fmt.Println(fv)
	sliceLen := sv.Len()
	out := reflect.MakeSlice(sv.Type(), 0, sliceLen)
	for i := 0; i < sliceLen; i++ {
		curVal := sv.Index(i)
		values := fv.Call([]reflect.Value{curVal})
		if values[0].Bool() {
			out = reflect.Append(out, curVal)
		}
	}
	return out.Interface()
}

func main() {
	names := []string{"Andrew", "Bob", "Clara", "Hortense"}
	longNames := Filter(names, func(s string) bool {
		return len(s) > 3
	}).([]string)
	fmt.Println(longNames)
	// [Andrew Clara Hortense]

	ages := []int{20, 50, 13}
	adults := Filter(ages, func(age int) bool {
		return age >= 18
	}).([]int)
	fmt.Println(adults)
	// [20 50]
}
