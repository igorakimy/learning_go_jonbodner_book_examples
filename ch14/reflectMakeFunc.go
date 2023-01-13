package main

import (
	"fmt"
	"reflect"
	"time"
)

// Функция принимает на вход любую функцию

func MakeTimedFunction(f interface{}) interface{} {
	// Получить экземпляр типа reflect.Type, представляющий эту функцию
	ft := reflect.TypeOf(f)
	fv := reflect.ValueOf(f)
	// Передаем этот экземпляр в функцию reflect.MakeFunc, вместе с замыканием
	wrapperF := reflect.MakeFunc(ft, func(in []reflect.Value) []reflect.Value {
		// Фиксируем начальный момент времени
		start := time.Now()
		// Вызываем исходную функцию с помощью рефлексии
		out := fv.Call(in)
		// Фиксируем конечный момент времени
		end := time.Now()
		// Выводим разницу между началом и концом
		fmt.Println(end.Sub(start))
		// Возвращаем значение, вычисленное исходной функцией
		return out
	})
	return wrapperF.Interface()
}

func timeMe(a int) int {
	time.Sleep(time.Duration(a) * time.Second)
	result := a * 2
	return result
}

func main() {
	timed := MakeTimedFunction(timeMe).(func(int) int)
	fmt.Println(timed(2))
	// 2.0083342s
	// 4
}
