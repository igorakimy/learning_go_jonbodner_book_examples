package main

import (
	"errors"
	"fmt"
	"net/http"
)

// Функция журналирования

func LogOutput(message string) {
	fmt.Println(message)
}

// Простое хранилище данных

type SimpleDataStorage struct {
	userData map[string]string
}

func (sds SimpleDataStorage) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

// Фабричная функция, для создания хранилища

func NewSimpleDataStorage() SimpleDataStorage {
	return SimpleDataStorage{
		userData: map[string]string{
			"1": "Fred",
			"2": "Mary",
			"3": "Pat",
		},
	}
}

// Интерфейс для хранилища данных

type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

// Интерфейс для журналирования

type Logger interface {
	Log(message string)
}

// Чтобы функция LogOutput соответствовала интерфейсу, нужно
// определить функциональный тип с требуемым методом

type LoggerAdapter func(message string)

func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

// Реализация бизнес-логики

type SimpleLogic struct {
	l  Logger
	ds DataStore
}

func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("in SayHello for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Hello, " + name, nil
}

func (sl SimpleLogic) SayGoodbye(userID string) (string, error) {
	sl.l.Log("is SayGoodbye for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Goodbye, " + name, nil
}

// Фабричная функция, которая принимает интерфейсы и возвращает структуру

func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

// Интерфейс бизнес-логики

type SomeLogic interface {
	SayHello(userID string) (string, error)
}

type Controller struct {
	l     Logger
	logic SomeLogic
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("In SayHello")
	userID := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

// Фабрична функция для контроллера

func NewController(l Logger, logic SomeLogic) Controller {
	return Controller{
		l:     l,
		logic: logic,
	}
}

func main() {
	l := LoggerAdapter(LogOutput)
	ds := NewSimpleDataStorage()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
	http.HandleFunc("/hello", c.SayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Start server error:", err)
	}
}
