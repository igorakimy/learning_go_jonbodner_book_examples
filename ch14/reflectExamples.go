package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x int
	type Foo struct {
		A int    `myTag:"value"`
		B string `myTag:"value2"`
	}
	// Получить название простого типа
	xt := reflect.TypeOf(x)
	fmt.Println(xt.Name())
	// int
	f := Foo{}
	// Получить название структурного типа
	ft := reflect.TypeOf(f)
	fmt.Println(ft.Name())
	// Foo
	xpt := reflect.TypeOf(&x)
	// При попытке получения названия типа у указателя возвращается пустая строка
	fmt.Println(xpt.Name())
	//
	// Получить вид типа
	fmt.Println(xpt.Kind())
	// ptr
	// Получить имя типа, на который ссылается указатель
	fmt.Println(xpt.Elem().Name())
	// int
	// Получить вид типа, на который ссылается указатель
	fmt.Println(xpt.Elem().Kind())
	// int

	// Получить количество полей в структуре
	for i := 0; i < ft.NumField(); i++ {
		// Получить конкретное поле структуры по его индексу
		curField := ft.Field(i)
		fmt.Println(
			// Получить название поля
			curField.Name,
			// Получить название типа поля
			curField.Type.Name(),
			// Получить значение тега поля
			curField.Tag.Get("myTag"),
		)
	}
	// A int value
	// B string value2

	s := []string{"a", "b", "c"}
	// Создать экземпляр reflect.Value, представляющий значение переменной
	sv := reflect.ValueOf(s)
	fmt.Println(sv)
	// [a b c]
	// Метод Interface возвращает значение переменной как пустой интерфейс
	s2 := sv.Interface().([]string)
	fmt.Println(s2)
	// [a b c]

	// Присвоение значений переменной, при помощи рефлексии
	i := 10
	// Получить значение указателя
	iv := reflect.ValueOf(&i)
	// Получить reflect.Value, на которое указывает указатель
	ivv := iv.Elem()
	// Задать новое значение reflect.Value, которое отразится на значении переменной
	ivv.SetInt(20)
	fmt.Println(i)
	// 20
}
