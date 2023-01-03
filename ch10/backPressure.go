package main

import (
	"errors"
	"log"
	"net/http"
	"time"
)

// Создаем структуру, которая содержит буферизованный канал с несколькими токенами.

type PressureGauge struct {
	ch chan struct{}
}

func (pg *PressureGauge) Process(f func()) error {
	// Оператор select пытается считать токен из канала.
	select {
	// Если это возможно
	case <-pg.ch:
		// Переданная функция выполняется
		f()
		pg.ch <- struct{}{}
		return nil
	// Если он не может считать токен, выполняется ветвь default
	default:
		// И вместо токена возвращается ошибка
		return errors.New("no more capacity")
	}
}

func New(limit int) *PressureGauge {
	ch := make(chan struct{}, limit)
	for i := 0; i < limit; i++ {
		ch <- struct{}{}
	}
	return &PressureGauge{
		ch: ch,
	}
}

func doThingThatShouldBeLimited() string {
	time.Sleep(2 * time.Second)
	return "done"
}

func main() {
	pg := New(10)
	http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		// Каждый раз, когда горутине требуется использовать функцию, она вызывает
		// функцию Process.
		err := pg.Process(func() {
			if _, err := w.Write([]byte(doThingThatShouldBeLimited())); err != nil {
				log.Fatal("failed to write to file")
			}
		})
		if err != nil {
			w.WriteHeader(http.StatusTooManyRequests)
			if _, err = w.Write([]byte("To many requests")); err != nil {
				log.Fatal("failed to write to file")
			}
		}
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("failed listen the server")
	}
}
