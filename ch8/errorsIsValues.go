package main

import (
	"errors"
	"fmt"
)

func login(userID string, password string) error {
	return errors.New("invalid login")
}

func getData(filename string) ([]byte, error) {
	return []byte("Hello World"), nil
}

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
}

func (se StatusErr) Error() string {
	return se.Message
}

// Использование структуры StatusErr для предоставления более подробных сведений
// о том, что пошло не так

func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
	err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status:  InvalidLogin,
			Message: fmt.Sprintf("invalid credentials for user %s", uid),
		}
	}
	data, err := getData(file)
	if err != nil {
		return nil, StatusErr{
			Status:  NotFound,
			Message: fmt.Sprintf("file %s not found", file),
		}
	}
	return data, nil
}

// Функция всегда должна возвращать инициализированный экземпляр кастомной ошибки

func GenerateError(flag bool) error {
	var genErr StatusErr
	if flag {
		genErr = StatusErr{
			Status: NotFound,
		}
	}
	return genErr
}

func main() {
	_, err := LoginAndGetData("40", "12345", "f")
	if err != nil {
		fmt.Println(err)
	}

	err = GenerateError(true)
	fmt.Println(err != nil)
	// true
	err = GenerateError(false)
	fmt.Println(err != nil)
	// true
}
