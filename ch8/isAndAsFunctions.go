package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
)

func fileChecker2(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker2: %w", err)
	}
	_ = f.Close()
	return nil
}

// Реализация метода Is для ошибки

type MyErr struct {
	Codes []int
}

func (m MyErr) Error() string {
	return fmt.Sprintf("codes: %v", m.Codes)
}

func (m MyErr) Is(target error) bool {
	if m2, ok := target.(MyErr); ok {
		// Функция DeepEqual сравнивает что угодно
		return reflect.DeepEqual(m, m2)
	}
	return false
}

// Выявить частичное соответствие с экземпляром ошибки, сравнив ошибки по маске

type ResourceErr struct {
	Resource string
	Code     int
}

func (re ResourceErr) Error() string {
	return fmt.Sprintf("%s: %d", re.Resource, re.Code)
}

func (re ResourceErr) Is(target error) bool {
	if other, ok := target.(ResourceErr); ok {
		ignoreResource := other.Resource == ""
		ignoreCode := other.Code == 0
		matchResource := other.Resource == re.Resource
		matchCode := other.Code == re.Code
		return matchResource && matchCode ||
			matchResource && ignoreCode ||
			ignoreResource && matchCode
	}
	return false
}

func AFunctionThatReturnsAnError() error {
	return MyErr2{
		Code: 456,
	}
}

type MyErr2 struct {
	Code int
}

func (mer MyErr2) Error() string {
	return fmt.Sprintf("That is MyErr2 with code: %d", mer.Code)
}

func main() {
	err := fileChecker2("not_here.txt")
	if err != nil {
		// Отлов системной ошибки
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("That file doesn't exist")
			// That file doesn't exist
		}
	}

	// Теперь можно найти все ошибки, которые имеют отношение к базе данных,
	// безотносительно к их коду
	err = ResourceErr{
		Resource: "Database",
		Code:     123,
	}

	if errors.Is(err, ResourceErr{Resource: "Database"}) {
		fmt.Println("The database is broken:", err)
		// The database is broken: Database: 123
	}

	// Проверить, не совпадает ли ошибка с определенным типом

	err = AFunctionThatReturnsAnError()
	var myErr MyErr2
	if errors.As(err, &myErr) {
		fmt.Println(myErr.Code)
		// 456
	}

	// Необязательно передавать указатель на переменную, можно передать
	// указатель на интерфейс
	err = AFunctionThatReturnsAnError()
	var coder interface {
		Code() int
	}
	if errors.As(err, &coder) {
		fmt.Println(coder.Code())
	}
}
