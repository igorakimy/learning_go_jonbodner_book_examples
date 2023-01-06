package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func longRunningThing(ctx context.Context, data string) (string, error) {
	time.Sleep(2 * time.Second)
	return "some string", errors.New("timeout error")
}

func longRunningThingManager(ctx context.Context, data string) (string, error) {
	type wrapper struct {
		result string
		err    error
	}
	// Создаем буферизованный канал типа wrapper, с размером буфера, равным 1.
	// Используя буферизованный канал, мы делаем возможным выход из горутины даже
	// в том случае, если буферизованное значение никогда не будет считано из-за отмены
	ch := make(chan wrapper, 1)
	go func() {
		// Принимаем результат, возвращаемый долго выполняющейся функцией
		result, err := longRunningThing(ctx, data)
		// и заносим его в буферизованный канал
		ch <- wrapper{result, err}
	}()
	select {
	// В первой ветви считываем и возвращаем данные, выданные долго выполняющейся функцией.
	// Эта ветвь срабатывает в том случае, когда контекст не отменяется срабатыванием
	// таймера или вызовом функции отмены
	case data := <-ch:
		return data.result, data.err
	// Вторая ветвь срабатывает в том случае, когда контекст отменяется.
	// При этом возвращаем нулевое значение в качестве данный и извлекаемую из контекста
	// ошибку, сообщая тем самым о причине выполнения отмены
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func main() {
	ctx := context.Background()
	res, err := longRunningThingManager(ctx, "Hello world")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
