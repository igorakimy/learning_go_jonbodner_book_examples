package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type HelloHandler struct{}

func (hh HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!\n"))
}

// Промежуточный генератор для рассчитывания времени выполнения запросов

func RequestTimer(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		end := time.Now()
		log.Printf("request time for %s: %v", r.URL.Path, end.Sub(start))
	})
}

// Неудачный промежуточный генератор для контроля доступа
var securityMsg = []byte("You didn't give the secret password\n")

func TerribleSecurityProvider(password string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("X-Secret-Password") != password {
				w.WriteHeader(http.StatusUnauthorized)
				_, _ = w.Write(securityMsg)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

func main() {
	// Клиентская сторона
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		"https://jsonplaceholder.typicode.com/todos/1",
		nil,
	)
	if err != nil {
		panic(err)
	}
	req.Header.Add("X-My-Client", "Learning Go")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("unexprected status: got %v", res.Status))
	}
	fmt.Println(res.Header.Get("Content-Type"))
	var data struct {
		UserID    int    `json:"userId"`
		ID        int    `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", data)

	// Серверная сторона
	person := http.NewServeMux()
	person.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("greetings!\n"))
	})
	person.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("goodbye!"))
	})

	dog := http.NewServeMux()
	dog.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("good puppy!\n"))
	})

	mux := http.NewServeMux()
	// Запрос адреса /person/greet обрабатывается обработчиками, привязанными к
	// переменной person...
	mux.Handle("/person/", http.StripPrefix("/person", person))
	// ...а запросы адреса /dog/greet - обработчиками, привязанными к переменной dog
	mux.Handle("/dog/", http.StripPrefix("/dog", dog))

	// Добавление промежуточных слоев к своим обработчикам
	terribleSecurity := TerribleSecurityProvider("GORHER")
	mux.Handle("/hello", terribleSecurity(RequestTimer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("Hello!\n"))
		}))))

	server := http.Server{
		// Трафик какого хоста и на каком порту будет прослушан
		Addr: ":8080",
		// Тайм-аут на чтение
		ReadTimeout: 30 * time.Second,
		// Тайм-аут на запись
		WriteTimeout: 30 * time.Second,
		// Там-аут режима ожидания
		IdleTimeout: 120 * time.Second,
		// Обработчик запросов
		Handler: mux,
	}
	err = server.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}
