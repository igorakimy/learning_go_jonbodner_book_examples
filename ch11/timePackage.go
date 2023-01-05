package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	d := 2*time.Hour + 30*time.Minute
	fmt.Println(d)
	// 2h30m0s

	fmt.Println(reflect.TypeOf(d))
	// time.Duration

	// Получить ссылку на текущий момент времени
	fmt.Println(time.Now())

	// 2023-01-03 21:18:13.0480469 +0300 MSK m=+0.001562101
	fmt.Println(reflect.TypeOf(time.Now()))
	// time.Time

	// Преобразование из string в time.Time
	t, _ := time.Parse("2006-02-01 15:04:05 -0700", "2023-04-01 04:30:00 +0300")
	fmt.Println(t)
	// 2023-01-04 04:30:00 +0300 MSK

	// Преобразование из time.Time в string
	t2 := t.Format("January 2, 2006 at 3:04:05PM MST")
	fmt.Println(t2)
	// January 4, 2023 at 4:30:00AM MSK

	fmt.Println(t.Year())
	// 2023
	fmt.Println(t.Date())
	// 2023 January 4
	h, m, s := t.Clock()
	fmt.Println(h, m, s)
	// 4 30 0

	// Сравнить время time.Time с time.Time
	fmt.Println(t.Equal(time.Now()))
	// false
	fmt.Println(t.Before(time.Now()))
	// true

	// Получить разницу во времени
	fmt.Println(t.Sub(time.Now()))
	// -14m9.0115946s
}
