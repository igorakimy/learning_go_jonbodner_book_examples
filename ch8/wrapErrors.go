package main

import (
	"errors"
	"fmt"
	"os"
)

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	_ = f.Close()
	return nil
}

// Обновленная версия кастомной ошибки, которая реализует метод Unwrap()

type Status2 int

const (
	InvalidLogin2 Status2 = iota + 1
	NotFound2
)

type StatusErr2 struct {
	Status  Status2
	Message string
	Err     error
}

func (se StatusErr2) Error() string {
	return se.Message
}

func (se StatusErr2) Unwrap() error {
	return se.Err
}

func login2(uid, pwd string) error {
	return fmt.Errorf("ivalid 1 credentials for user: %v", uid)
}

func getData2(filename string) ([]byte, error) {
	return nil, fmt.Errorf("file %v not found", filename)
}

func loginAndGetData2(uid, pwd, file string) ([]byte, error) {
	err := login2(uid, pwd)
	if err != nil {
		return nil, StatusErr2{
			Status:  InvalidLogin2,
			Message: fmt.Sprintf("invalid 2 credentials for user %s", uid),
			Err:     err,
		}
	}
	data, err := getData2(file)
	if err != nil {
		return nil, StatusErr2{
			Status:  NotFound2,
			Message: fmt.Sprintf("file %s not found", file),
			Err:     err,
		}
	}
	return data, nil
}

func main() {
	err := fileChecker("not_here.txt")
	if err != nil {
		fmt.Println(err)
		// in fileChecker: open not_here.txt: The system cannot find the file specified.
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println(wrappedErr)
			// open not_here.txt: The system cannot find the file specified.
		}
	}

	_, err = loginAndGetData2("4034", "password", "f")
	if err != nil {
		fmt.Println(errors.Unwrap(err))
	}
}
