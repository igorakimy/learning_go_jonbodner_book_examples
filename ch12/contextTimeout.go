package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	// Устанавливаем двухсекундный тайм-аут в родительском контексте
	parent, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	// Устанавливаем трехсекундный тайм-аут в дочернем контексте
	child, cancel2 := context.WithTimeout(parent, 3*time.Second)
	defer cancel2()
	start := time.Now()
	// Затем ждем отмены дочернего контекста, дожидаясь момента возвращения значения
	// каналом, возвращенным методом Done, вызванным в дочернем экземпляре типа
	// context.Context
	<-child.Done()
	end := time.Now()
	fmt.Println(end.Sub(start))
	// 2.0081765s
}
